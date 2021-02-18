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
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser2 "github.com/devtron-labs/inception/pkg/language/parser"
	"io/ioutil"
	"testing"
)

func TestKlangListener_handleNestedIf(t *testing.T) {
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser2.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Nested if",
			fields: fields{
				input: `
name = "name";
if name == name && name == name && name == name && name == name || name == name { a=1+2; }
if 2==2 {
if 1==1 {
a = 2; 
}
else {
a=4;
}
} 
else { a = 6;}
b = 2 - 3;
c = 6+8;
d = 2 / 3.3;
e = 3.3 * 2.2;
f = "abc" + name;`,
				values: map[string]valHolder{
					"a": {
						dataType: INT,
						name:     "a",
						value:    int64(2),
					},
					"b": {
						dataType: INT,
						name:     "b",
						value:    int64(-1),
					},
					"c": {
						dataType: INT,
						name:     "c",
						value:    int64(14),
					},
					"d": {
						dataType: FLOAT,
						name:     "d",
						value:    0.6060606060606061,
					},
					"e": {
						dataType: FLOAT,
						name:     "e",
						value:    7.26,
					},
					"f": {
						dataType: STRING,
						name:     "f",
						value:    "abcname",
					},
					"name": {
						dataType: STRING,
						name:     "name",
						value:    "name",
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			if diff := compare(r.Values(), tt.fields.values); !diff {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_handleWhile(t *testing.T) {
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser2.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "While",
			fields: fields{
				input: `
x = 0;
while x < 2 {
x = x+1;
}`,
				values: map[string]valHolder{
					"x": {
						dataType: "INT",
						name:     "x",
						value:    int64(2),
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			if diff := compare(tt.fields.values, r.values); !diff {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_handlekubectlget(t *testing.T) {
	d := `x = kubectl get -n dev cm/test-cm;
z = "metadata.name";
y = jsonSelect(x, z);
if !x {
  log("missing");
}
`
	a := `items.#(metadata.labels.app\.kubernetes\.io/name=="argocd-server").metadata.name`
	e := `x = kubectl get -n devtroncd po;
z = jsonSelect(x, `
	e = e + "`" + a + "`);"
	fmt.Printf("%s", d)
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser2.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "kubectl get",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"y": {
						dataType: "STRING",
						name:     "y",
						value:    "test-cm",
					},
				},
			},
			args: args{},
		},
		{
			name: "kubectl get multi",
			fields: fields{
				input: e,
				values: map[string]valHolder{
					"y": {
						dataType: "STRING",
						name:     "y",
						value:    "test-cm",
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
				if diff := compare(tt.fields.values, m); diff {
					t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
				}
			} else {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_handlekubectlapply(t *testing.T) {
	a := `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm
  namespace: dev
  labels:
    app.kubernetes.io/instance: my-app
data:
  name: abc
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm2
  namespace: dev
  labels:
    app.kubernetes.io/instance: my-app
data:
  name: def
`
	d := "a = `" + a + "`" + `;
x = kubectl apply a;
`
	u := `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm2
update:
  data:
    name: xyz
`
	e := "a = `" + a + "`;\n" + "u = `" + u + "`;" + `
x = kubectl apply a -u u
`
	fmt.Printf("%s", d)
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser2.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "kubectl apply",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"x": {
						dataType: BOOLEAN,
						name:     "x",
						value:    true,
					},
				},
			},
			args: args{},
		},
		{
			name: "kubectl apply with edit",
			fields: fields{
				input: e,
				values: map[string]valHolder{
					"x": {
						dataType: BOOLEAN,
						name:     "x",
						value:    true,
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			x := r.values["x"]
			m := map[string]valHolder{
				"x": x,
			}
			if diff := compare(tt.fields.values, m); !diff {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_handlekubectlpatch(t *testing.T) {
	d := `a = kubectl patch -n dev cm test-cm --type "application/merge-patch+json" -p '{"data":{"age":"36"}}';
b = kubectl get -n dev cm test-cm;
c = jsonSelect(b, "data.age");`
	fmt.Printf("%s", d)
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser2.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "kubectl patch",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"c": {
						dataType: STRING,
						name:     "c",
						value:    "36",
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			c := r.values["c"]
			m := map[string]valHolder{
				"c": c,
			}
			if diff := compare(tt.fields.values, m); !diff {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_handlekubectldelete(t *testing.T) {
	d := `a = kubectl delete -n dev cm test-cm test-cm2;`
	fmt.Printf("%s", d)
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser2.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "kubectl delete",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"a": {
						dataType: BOOLEAN,
						name:     "a",
						value:    true,
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			a := r.values["a"]
			m := map[string]valHolder{
				"a": a,
			}
			if diff := compare(tt.fields.values, m); !diff {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_integration(t *testing.T) {
	a := `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm
  namespace: dev
  labels:
    app.kubernetes.io/instance: my-app
data:
  name: abc
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm2
  namespace: dev
  labels:
    app.kubernetes.io/instance: my-app
data:
  name: def
`
	d := "a = `" + a + "`" + `;
x = kubectl apply -n dev a;
k = "kind";
z = "metadata.name";
o1 = yamlSelect(a, k, 0) + "/" + yamlSelect(a, z, 0);
k2 = yamlSelect(a, k, 1);
n2 = yamlSelect(a, z, 1);
o2 = k2 + "/" + n2;
age = "36";
pla = '{"data":{"age":"' + age + '"}}';
pa = kubectl patch -n dev k2 n2 --type "application/merge-patch+json" -p pla;
fo = kubectl get -n dev o1 o2;
selector = 'items.#(metadata.name="'+n2+'").data.age';
age = jsonSelect(fo, selector);
`
	//programmers.#(lastName="Hunter").firstName
	fmt.Printf("%s", d)
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser2.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "kubectl integration",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"age": {
						dataType: "STRING",
						name:     "age",
						value:    "36",
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			age := r.values["age"]
			m := map[string]valHolder{
				"age": age,
			}
			if diff := compare(tt.fields.values, m); !diff {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_secret_wo_base64(t *testing.T) {
	a := `
apiVersion: v1
kind: Secret
metadata:
  name: test-cm
  namespace: dev
  labels:
    app.kubernetes.io/instance: my-app
data:
  name: abc
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm2
  namespace: dev
  labels:
    app.kubernetes.io/instance: my-app
data:
  name: def
`
	d := "a = `" + a + "`" + `;
x = kubectl apply -n pras a;
migDelete = kubectl delete -n devtroncd job  postgresql-migrate-devtron postgresql-migrate-casbin postgresql-migrate-gitsensor  postgresql-migrate-lens;
;
`
	//programmers.#(lastName="Hunter").firstName
	fmt.Printf("%s", d)
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser2.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "kubectl integration",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"age": {
						dataType: "STRING",
						name:     "age",
						value:    "36",
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			age := r.values["age"]
			m := map[string]valHolder{
				"age": age,
			}
			if diff := compare(tt.fields.values, m); !diff {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_temp(t *testing.T) {
	//	base64DecoderPrefix := `#!/bin/bash
	//echo -n "`
	//	base64DecoderSuffix := `" | base64 -d`
	data, err := ioutil.ReadFile("/Users/pghildiy/Documents/devtronCode/installation-yamls/devtron-installation-script/installation-script-part")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", data)
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser2.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "kubectl temp",
			fields: fields{
				input: string(data),
				values: map[string]valHolder{
					"age": {
						dataType: "STRING",
						name:     "age",
						value:    "36",
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			age := r.values["age"]
			m := map[string]valHolder{
				"age": age,
			}
			if diff := compare(tt.fields.values, m); !diff {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_handleShellScript(t *testing.T) {
	d := `
#!/bin/bash
echo 'hello' | base64`
	d = "a = shellScript `" + d + "`;"
	fmt.Printf("%s", d)
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser2.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "shell script",
			fields: fields{
				input: d,
				values: map[string]valHolder{
					"a": {
						dataType: STRING,
						name:     "a",
						value:    "hello\n",
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			a := r.values["a"]
			m := map[string]valHolder{
				"a": a,
			}
			if diff := compare(tt.fields.values, m); !diff {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func setup(input string) *KlangListener {
	is := antlr.NewInputStream(input)
	lexer := parser2.NewKlangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser2.NewKlangParser(stream)
	p.BuildParseTrees = true
	mapper := NewMapperFactory()
	r := NewKlangListener(mapper)
	antlr.ParseTreeWalkerDefault.Walk(r, p.Parse())
	return r
}

func compare(first, second map[string]valHolder) bool {
	return checkFirstInSecond(first, second) && checkFirstInSecond(second, first)
}

func checkFirstInSecond(first, second map[string]valHolder) bool {
	for k, fv := range first {
		if sv, ok := second[k]; ok {
			ft := fmt.Sprintf("%T", fv.value)
			st := fmt.Sprintf("%T", sv.value)
			r := ft != st || fv.dataType != sv.dataType || fv.value != sv.value
			if r {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

//func compareInterface(first, second interface{}) bool {
//	ft := fmt.Sprintf("%T", first)
//	st := fmt.Sprintf("%T", second)
//	if ft != st {
//		return false
//	}
//	switch first.(type) {
//	case int64:
//		return first.(int64) == second.(int64)
//	case float64:
//		return first.(float64) == second.(float64)
//	case bool:
//		return first.(bool) == second.(bool)
//	case string:
//		return first.(string) == second.(string)
//	default:
//		return false
//	}
//}

