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
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/restmapper"
	"reflect"
	"testing"
)

func TestNewFactory(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		args args
		want *restmapper.DeferredDiscoveryRESTMapper
	}{
		{
			name: "factory test",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := schema.GroupVersionResource{
				Resource: "custom.metrics.k8s.io",
			}
			m := NewMapperFactory()
			r := NewFactory(m)
			o, err := r.mapper.mapper.ResourcesFor(res)
			if err != nil {
				panic(err)
			}
			fmt.Printf("out %+v\n", o)
			if got := NewFactory(m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgsProcessor_ResourceTypeOrNameArgs(t *testing.T) {
	type fields struct {
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "ResourceType test",
			fields: fields{},
			args: args{
				args: []string{"po", "abc"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMapperFactory()
			a := NewFactory(m)
			a.ResourceTypeOrNameArgs(tt.args.args...)
			fmt.Printf("resources %+v, names %v\n", a.resources, a.names)
			if len(a.resources) > 0 {
				restMapping, err := a.mappingFor(a.resources[0])
				if err != nil {
					panic(err)
				}
				fmt.Printf("resources %+v, names %v, out %+v\n", a.resources, a.names, restMapping)
			} else if len(a.resourceTuples) > 0 {
				for _, rt := range a.resourceTuples {
					restMapping, err := a.mappingFor(rt.Resource)
					if err != nil {
						panic(err)
					}
					fmt.Printf("resourceTouples %+v, out %+v\n", rt, restMapping)
				}
			}

		})
	}
}
