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

package pkg

import (
	"context"
	"github.com/argoproj/gitops-engine/pkg/utils/kube"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"reflect"
	"testing"
)

//TODO: test using mock methods
func Test_kubectl_GetResource(t *testing.T) {
	type fields struct {
	}
	type args struct {
		gvk       schema.GroupVersionKind
		name      string
		namespace string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "test get resource",
			fields: fields{},
			args: args{
				gvk:       schema.GroupVersionKind{Version: "v1", Kind: "Service"},
				name:      "swagger-swaggerui",
				namespace: "default",
			},
			want: "/Service/default/swagger-swaggerui",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := NewKubectl()
			r := GetRequest{
				Name:             tt.args.name,
				Namespace:        tt.args.namespace,
				GroupVersionKind: tt.args.gvk,
			}
			got, err := k.GetResource(context.Background(), &r)
			if err != nil {
				t.Errorf("GetResource() error = %v", err)
			}
			out, err := kube.SplitYAML([]byte(got.Manifest))
			if err != nil {
				t.Errorf("GetResource() error = %v", err)
			}
			for _, o := range out {
				gr := kube.GetResourceKey(o)
				if gr.String() != tt.want {
					t.Errorf("GetResource() got = %s, want %s", gr.String(), tt.want)
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_kubectl_PatchResource(t *testing.T) {
	pr := PatchRequest{
		Name:      "test-cm2",
		Namespace: "dev",
		GroupVersionKind: schema.GroupVersionKind{
			Version: "v1",
			Kind:    "ConfigMap",
		},
		Patch:     `{"data":{"age":"35"}}`,
		PatchType: "application/merge-patch+json",
	}
	type fields struct{}
	type args struct {
		r *PatchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ManifestResponse
		wantErr bool
	}{
		{
			name: "patch configmap",
			args: args{
				r: &pr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := NewKubectl()
			got, err := k.PatchResource(context.TODO(), tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("PatchResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PatchResource() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kubectl_DeleteResource(t *testing.T) {
	ad := DeleteRequest{
		Name:      "test-cm",
		Namespace: "dev",
		GroupVersionKind: schema.GroupVersionKind{
			Version: "v1",
			Kind:    "ConfigMap",
		},
		Force: nil,
	}
	type fields struct {
	}
	type args struct {
		r *DeleteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ManifestResponse
		wantErr bool
	}{
		{
			name: "delete resource",
			args: args{
				r: &ad,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := NewKubectl()
			got, err := k.DeleteResource(context.TODO(), tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteResource() got = %v, want %v", got, tt.want)
			}
		})
	}
}

const apply = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm
  namespace: dev
  labels:
    app.kubernetes.io/instance: my-app
data:
  name: Prashant
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm2
  namespace: dev
  labels:
    app.kubernetes.io/instance: my-app
data:
  name: Ghildiyal
`

func Test_kubectl_ApplyResource(t *testing.T) {
	ar := ApplyRequest{
		Manifest:  apply,
		Namespace: "dev",
		Force:     nil,
		Validate:  nil,
	}
	type fields struct {
	}
	type args struct {
		r *ApplyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []ApplyResponse
		wantErr bool
	}{
		{
			name: "apply resource",
			args: args{r: &ar},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := NewKubectl()
			got, err := k.ApplyResource(context.TODO(), tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyResource() got = %v, want %v", got, tt.want)
			}
		})
	}
}
