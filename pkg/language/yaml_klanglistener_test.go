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
	"github.com/devtron-labs/inception/pkg/language/parser"
	"testing"
)

func TestKlangListener_handleYamlSelect(t *testing.T) {
	y := `
name:
  first: abc
  last: def
`
	d := "x = `" + y + "`" + `;
y = yamlSelect(x, "name.last");
`
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
			name: "Yaml Select",
			fields: fields{
				input: d,
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

func TestKlangListener_handleYamlMultiSelect(t *testing.T) {
	y := `
name:
  first: ghi
  last: jkl
---
name:
  first: abc
  last: def
`
	d := "x = `" + y + "`" + `;
y = yamlSelect(x, "name.last", 1);
`
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
			name: "Yaml Multi Select",
			fields: fields{
				input: d,
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

func TestKlangListener_handleYamlEdit(t *testing.T) {
	y := `
name:
  first: abc
  last: def
`
	d := "x = `" + y + "`" + `;
yamlEdit(x, "name.first", "xyz");
`
	o := `name:
  first: xyz
  last: def
`
	fmt.Printf("%s", d)
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
			name: "Yaml Edit",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"x": {
						dataType: "STRING",
						name:     "x",
						value:    o,
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

func TestKlangListener_handleYamlMultiEdit(t *testing.T) {
	y := `
name:
  first: efg
  last: hij
---
name:
  first: abc
  last: def
`
	d := "x = `" + y + "`" + `;
yamlEdit(x, "name.first", "xyz", 1);
`
	o := `
name:
  first: efg
  last: hij
---
name:
  first: xyz
  last: def
`
	fmt.Printf("%s", d)
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
			name: "Yaml Edit",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"x": {
						dataType: "STRING",
						name:     "x",
						value:    o,
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

func TestKlangListener_handleYamlEdit_var(t *testing.T) {
	y := `
name:
  first: abc
  last: def
`
	d := "x = `" + y + "`" + `;
y = "name.first";
z = "xyz";
yamlEdit(x, y, z);
`
	o := `name:
  first: xyz
  last: def
`
	fmt.Printf("%s", d)
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
			name: "Yaml Edit with var argument",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"x": {
						dataType: "STRING",
						name:     "x",
						value:    o,
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

func TestKlangListener_handleYamlEdit_id_assignment(t *testing.T) {
	y := `name:
  first: abc
  last: def
`
	d := "x = `" + y + "`" + `;
y = x;
`
	o := `name:
  first: abc
  last: def
`
	fmt.Printf("%s", d)
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
			name: "id to id assignment",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"x": {
						dataType: "STRING",
						name:     "x",
						value:    o,
					},
					"y": {
						dataType: "STRING",
						name:     "y",
						value:    o,
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			if diff := compare(tt.fields.values, r.Values()); !diff {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

