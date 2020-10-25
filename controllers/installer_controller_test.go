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
	installerv1alpha1 "github.com/devtron-labs/inception/api/v1alpha1"
	"github.com/devtron-labs/inception/pkg"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
)

func TestInstallerReconciler_apply(t *testing.T) {
	scheme := runtime.NewScheme()
	mapper := pkg.NewMapperFactory()
	config := ctrl.GetConfigOrDie()
	mgr, err := ctrl.NewManager(config, ctrl.Options{Scheme: scheme})
	installer := &installerv1alpha1.Installer{
		Spec: installerv1alpha1.InstallerSpec{URL: "https://raw.githubusercontent.com/pghildiyal/script-test/main/script"},
		Status: installerv1alpha1.InstallerStatus{
			Sync:            installerv1alpha1.SyncStatus{
				Status:     installerv1alpha1.SyncStatusCodeOutOfSync,
				URL:        "https://raw.githubusercontent.com/pghildiyal/script-test/main/script",
				Data:       "",
			},
		},
	}
	if err != nil {
		panic(err)
	}
	type fields struct {
		Client client.Client
		Log    logr.Logger
		Scheme *runtime.Scheme
		Mapper *pkg.Mapper
	}
	type args struct {
		installer *installerv1alpha1.Installer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "reconciler parsing",
			fields: fields{
				Client: mgr.GetClient(),
				Log:    mgr.GetLogger(),
				Scheme: scheme,
				Mapper: mapper,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InstallerReconciler{
				Client: tt.fields.Client,
				Log:    tt.fields.Log,
				Scheme: tt.fields.Scheme,
				Mapper: tt.fields.Mapper,
			}
			r.sync(installer)
			r.apply(installer)
		})
	}
}

func Test_downloadDSL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "download file",
			args: args{url: "https://raw.githubusercontent.com/pghildiyal/script-test/main/script"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := downloadDSL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("downloadDSL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("downloadDSL() got = %v, want %v", got, tt.want)
			}
		})
	}
}