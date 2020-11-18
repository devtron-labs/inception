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
	"github.com/devtron-labs/inception/pkg/language/parser"
	"testing"
)

func TestKlangListener_handleJsonSelect(t *testing.T) {
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Json Select",
			fields: fields{
				input: `
x = {"name":{"first":"abc","last":"def"}};
y = jsonSelect(x, "name.last");
`,
				values: map[string]valHolder{
					"y": {
						dataType: "STRING",
						name:     "y",
						value:    "def",
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			if d, ok := r.values["y"]; ok {
				m := map[string]valHolder{
					"y": d,
				}
				if diff := compare(tt.fields.values, m); !diff {
					t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
				}
			} else {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_handleJsonEdit(t *testing.T) {
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Json Edit",
			fields: fields{
				input: `
x = {"name":{"first":"abc","last":"def"}};
jsonEdit(x, "name.first", "xyz");
`,
				values: map[string]valHolder{
					"x": {
						dataType: "STRING",
						name:     "x",
						value:    "{\"name\":{\"first\":\"xyz\",\"last\":\"def\"}}",
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			if d, ok := r.values["x"]; ok {
				m := map[string]valHolder{
					"x": d,
				}
				if diff := compare(tt.fields.values, m); !diff {
					t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
				}
			} else {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

