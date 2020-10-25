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

import "testing"

func Test_updateKubernetesObjectsJson(t *testing.T) {
	type args struct {
		from     string
		to       string
		patterns []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateKubernetesObjectsJson(tt.args.from, tt.args.to, tt.args.patterns); got != tt.want {
				t.Errorf("updateKubernetesObjectsJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateKubernetesObjectsYaml(t *testing.T) {
	type args struct {
		from     string
		to       string
		patterns []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateKubernetesObjectsYaml(tt.args.from, tt.args.to, tt.args.patterns); got != tt.want {
				t.Errorf("updateKubernetesObjectsYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}