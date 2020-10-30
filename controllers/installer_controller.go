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

package controllers

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/devtron-labs/inception/pkg"
	"github.com/devtron-labs/inception/pkg/parser"
	"io"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"net/http"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	installerv1alpha1 "github.com/devtron-labs/inception/api/v1alpha1"
)

// InstallerReconciler reconciles a Installer object
type InstallerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	//Instead of KLangListener
	Mapper *pkg.Mapper
}

// +kubebuilder:rbac:groups=installer.devtron.ai,resources=installers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=installer.devtron.ai,resources=installers/status,verbs=get;update;patch

func (r *InstallerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("installer", req.NamespacedName)

	installer := &installerv1alpha1.Installer{}
	err := r.Client.Get(context.Background(), req.NamespacedName, installer)
	if err != nil {
		//Dont requeue if it is not found
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		//This will requeue the request
		return reconcile.Result{}, err
	}

	if len(installer.Spec.URL) == 0 {
		return reconcile.Result{}, fmt.Errorf("url is not specified")
	}
	updated := false
	if hasURLChanged(installer) {
		fmt.Println("url changed")
		installer.Status.Sync.Status = installerv1alpha1.SyncStatusCodeOutOfSync
		installer.Status.Sync.URL = installer.Spec.URL
		updated = true
	} else if shouldDownload(installer) {
		fmt.Println("should download")
		err := r.sync(installer)
		if err != nil {
			return reconcile.Result{}, err
		}
		updated = true
	} else if shouldApply(installer) {
		fmt.Println("applying")
		r.apply(installer)
		updated = true
	}

	if updated {
		fmt.Println("updating")
		err = r.Client.Update(context.Background(), installer)
		if err != nil {
			return reconcile.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}

func (r *InstallerReconciler) sync(installer *installerv1alpha1.Installer) error {
	//Download from url
	url := installer.Status.Sync.URL
	data, err := downloadDSL(url)
	if err != nil {
		return err
	}
	//Update url on which action was taken to handle race condition
	installer.Status.Sync.URL = url
	installer.Status.Sync.Data = data
	installer.Status.Sync.Status = installerv1alpha1.SyncStatusCodeDownloaded
	return nil
}

func (r *InstallerReconciler) apply(installer *installerv1alpha1.Installer) *pkg.KlangListener {
	url := installer.Status.Sync.URL
	data := installer.Status.Sync.Data
	//parse and process data
	is := antlr.NewInputStream(data)
	lexer := parser.NewKlangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewKlangParser(stream)
	p.BuildParseTrees = true
	listener := pkg.NewKlangListener(r.Mapper)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Parse())
	//TODO: Use r.Parser.Values() to check data and get resources
	//Update the status of resources
	var resourceStatuses []installerv1alpha1.ResourceStatus
	for _, resources := range listener.KubernetesResources() {
		for _, resource := range resources {
			rs := installerv1alpha1.ResourceStatus{
				Group:     resource.Group,
				Version:   resource.Version,
				Kind:      resource.Kind,
				Namespace: resource.Namespace,
				Name:      resource.Name,
				Status:    toSyncStatus(resource.Status),
				Health:    nil,
				Operation: resource.Operation,
			}
			resourceStatuses = append(resourceStatuses, rs)
		}
	}
	//Update URL and data on which action was taken to handle race condition
	installer.Status.Sync.URL = url
	installer.Status.Sync.Data = data
	installer.Status.Sync.Status = installerv1alpha1.SyncStatusCodeApplied
	installer.Status.Sync.Resources = resourceStatuses
	return listener
}

func toSyncStatus(code pkg.ResourceSyncStatusCode) installerv1alpha1.SyncStatusCode {
	switch code {
	case pkg.ResourceSyncStatusCodeOutOfSync:
		return installerv1alpha1.SyncStatusCodeOutOfSync
	case pkg.ResourceSyncStatusCodeSynced:
		return installerv1alpha1.SyncStatusCodeSynced
	case pkg.ResourceSyncStatusCodeUnknown:
		return installerv1alpha1.SyncStatusCodeUnknown
	default:
		return installerv1alpha1.SyncStatusCodeUnknown
	}
}

func downloadDSL(url string) (string, error) {
	hasher := sha1.New()
	hasher.Write([]byte(url))
	hash := hex.EncodeToString(hasher.Sum(nil))
	fileName := fmt.Sprintf("/tmp/%s", hash)
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			//Log
		}
	}()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func hasURLChanged(installer *installerv1alpha1.Installer) bool {
	if installer.Spec.URL != installer.Status.Sync.URL {
		return true
	}
	return false
}

func shouldDownload(installer *installerv1alpha1.Installer) bool {
	if installer.Status.Sync.Status == installerv1alpha1.SyncStatusCodeOutOfSync {
		return true
	}
	return false
}

func shouldApply(installer *installerv1alpha1.Installer) bool {
	if installer.Status.Sync.Status == installerv1alpha1.SyncStatusCodeDownloaded {
		return true
	}
	return false
}

func (r *InstallerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&installerv1alpha1.Installer{}).
		Complete(r)
}
