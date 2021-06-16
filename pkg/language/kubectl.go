/*
Copyright 2020 Devtron Labs Pvt Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package language

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/argoproj/gitops-engine/pkg/diff"
	"github.com/argoproj/gitops-engine/pkg/utils/kube"
	"github.com/argoproj/gitops-engine/pkg/utils/tracing"
	log "github.com/sirupsen/logrus"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/klog/klogr"
	"k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"time"
)

const (
	crdReadinessTimeout = time.Duration(3) * time.Second
)

type kubectl struct {
	restConfig          *rest.Config
	kubectl             kube.Kubectl
	extensionsClientset *clientset.Clientset
}

func NewKubectl() *kubectl {
	restConfig := ctrl.GetConfigOrDie()
	extensionsClientset, err := clientset.NewForConfig(restConfig)
	if err != nil {
		return nil
	}
	klog := klogr.New()
	return &kubectl{
		restConfig,
		&kube.KubectlCmd{
			Log:    klog,
			Tracer: tracing.NopTracer{},
		},
		extensionsClientset,
	}
}

func replaceSecretValues(obj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	if obj.GetKind() == kube.SecretKind && obj.GroupVersionKind().Group == "" {
		_, obj, err := diff.HideSecretData(nil, obj)
		if err != nil {
			return nil, err
		}
		return obj, err
	}
	return obj, nil
}

func (k *kubectl) GetResource(ctx context.Context, r *GetRequest) (*ManifestResponse, error) {
	obj, err := k.kubectl.GetResource(ctx, k.restConfig, r.GroupVersionKind, r.Name, r.Namespace)
	if err != nil {
		return nil, err
	}
	//obj, err = replaceSecretValues(obj)
	//if err != nil {
	//	return nil, err
	//}
	data, err := json.Marshal(obj.Object)
	if err != nil {
		return nil, err
	}
	logCtx := log.WithField("gkv", r.GroupVersionKind).WithField("name", r.Name).WithField("namespace", r.Namespace)
	logCtx.Infof("final data %s", string(data))
	return &ManifestResponse{string(data)}, nil
}

// PatchResource patches a resource
func (k *kubectl) PatchResource(ctx context.Context, r *PatchRequest) (*ManifestResponse, error) {
	manifest, err := k.kubectl.PatchResource(ctx, k.restConfig, r.GroupVersionKind, r.Name, r.Namespace, types.PatchType(r.PatchType), []byte(r.Patch))
	if err != nil {
		return nil, err
	}
	manifest, err = replaceSecretValues(manifest)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(manifest.Object)
	if err != nil {
		return nil, err
	}
	return &ManifestResponse{
		Manifest: string(data),
	}, nil
}

// DeleteResource deletes a specified resource
func (k *kubectl) DeleteResource(ctx context.Context, r *DeleteRequest) (*ManifestResponse, error) {
	deleteOptions := metav1.DeleteOptions{}
	if r.Force != nil {
		deleteOptions.GracePeriodSeconds = pointer.Int64Ptr(0)
	}
	err := k.kubectl.DeleteResource(ctx, k.restConfig, r.GroupVersionKind, r.Name, r.Namespace, deleteOptions)
	if err != nil {
		return nil, err
	}
	return &ManifestResponse{}, nil
}

func (k *kubectl) ApplyResource(ctx context.Context, r *ApplyRequest) ([]ApplyResponse, error) {
	objs, err := kube.SplitYAML([]byte(r.Manifest))
	var manifests []unstructured.Unstructured
	if err != nil {
		return make([]ApplyResponse, 0), err
	}
	for _, obj := range objs {
		if obj.IsList() {
			err = obj.EachListItem(func(object runtime.Object) error {
				unstructuredObj, ok := object.(*unstructured.Unstructured)
				if ok {
					manifests = append(manifests, *unstructuredObj)
					return nil
				}
				return fmt.Errorf("resource list item has unexpected type")
			})
			if err != nil {
				return make([]ApplyResponse, 0), err
			}
		} else if isNullList(obj) {
			// noop
		} else {
			manifests = append(manifests, *obj)
		}
	}
	var force, validate bool
	if r.Force != nil {
		force = *r.Force
	}
	if r.Validate != nil {
		validate = *r.Validate
	}
	manifests = SortManifestsByKind(manifests, InstallOrder)
	//https://github.com/argoproj/gitops-engine/blob/master/pkg/sync/sync_context.go
	responses := make([]ApplyResponse, 0)
	for _, manifest := range manifests {
		message, err := k.kubectl.ApplyResource(context.Background(), k.restConfig, &manifest, r.Namespace, util.DryRunNone, force, validate)
		res := ApplyResponse{
			Message:          message,
			GroupVersionKind: manifest.GroupVersionKind(),
			Name:             manifest.GetName(),
			Namespace:        manifest.GetNamespace(),
		}
		if err != nil {
			res.Err = err.Error()
		}
		responses = append(responses, res)
		if kube.IsCRD(&manifest) {
			k.ensureCRDReady(manifest.GetName())
		}
	}
	return responses, nil
}

func (k *kubectl) ListResources(ctx context.Context, r *ListRequest) (*ListResponse, error) {
	dynamicIf, err := dynamic.NewForConfig(k.restConfig)
	if err != nil {
		return nil, err
	}
	resourceList, err := dynamicIf.Resource(r.GroupVersionResource).Namespace(r.Namespace).List(ctx, r.ListOptions)
	if err != nil {
		return nil, err
	}
	var manifests []string
	for i, _ := range resourceList.Items {
		item := &resourceList.Items[i]
		//item, err := replaceSecretValues(item)
		//if err != nil {
		//	return nil, err
		//}
		data, err := json.Marshal(item.Object)
		if err != nil {
			return nil, err
		}
		manifests = append(manifests, string(data))
	}
	return &ListResponse{Manifests: manifests}, nil
}

// https://github.com/argoproj/gitops-engine/blob/master/pkg/sync/sync_context.go
// ensureCRDReady waits until specified CRD is ready (established condition is true). Method is best effort - it does not fail even if CRD is not ready without timeout.
func (k *kubectl) ensureCRDReady(name string) {
	_ = wait.PollImmediate(time.Duration(100)*time.Millisecond, crdReadinessTimeout, func() (bool, error) {
		crd, err := k.extensionsClientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, condition := range crd.Status.Conditions {
			if condition.Type == v1beta1.Established {
				return condition.Status == v1beta1.ConditionTrue, nil
			}
		}
		return false, nil
	})
}

// https://github.com/argoproj/gitops-engine/blob/master/pkg/sync/sync_context.go
// isNullList checks if the object is a "List" type where items is null instead of an empty list.
// Handles a corner case where obj.IsList() returns false when a manifest is like:
// ---
// apiVersion: v1
// items: null
// kind: ConfigMapList
func isNullList(obj *unstructured.Unstructured) bool {
	if _, ok := obj.Object["spec"]; ok {
		return false
	}
	if _, ok := obj.Object["status"]; ok {
		return false
	}
	field, ok := obj.Object["items"]
	if !ok {
		return false
	}
	return field == nil
}
