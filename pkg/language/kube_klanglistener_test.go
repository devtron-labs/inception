package language

import (
	json2 "encoding/json"
	"fmt"
	"github.com/argoproj/gitops-engine/pkg/utils/kube"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"reflect"
	"sigs.k8s.io/yaml"
	"strings"
	"testing"
)

func Test_handleJsonDelete(t *testing.T) {
	jsonList := `{"apiVersion": "v1", "kind": "List", "items":[{"apiVersion": "v1", "kind": "service", "metadata": {"name": "abc", "namespace": "abc"}, "data": {"school": "abc"}}, {"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "def"}}]}`
	json := `{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "def"}}`
	pattern := "data.school"
	o := unstructured.Unstructured{}
	o.UnmarshalJSON([]byte(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "def"}}`))
	resourceKey := kube.GetResourceKey(&o)
	filter := (&resourceKey).String()
	type args struct {
		data    string
		filter  string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want valHolder
	}{
		{
			name: "delete in list with filter and pattern",
			args: args{
				data:    jsonList,
				filter:  filter,
				pattern: pattern,
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "List", "items":[{"apiVersion": "v1", "kind": "service", "metadata": {"name": "abc", "namespace": "abc"}, "data": {"school": "abc"}}, {"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {}}]}`),
		},
		{
			name: "delete in list with filter only",
			args: args{
				data:   jsonList,
				filter: filter,
			},
			want: newStringValHolder(`{"apiVersion":"v1","items":[{"apiVersion":"v1","data":{"school":"abc"},"kind":"service","metadata":{"name":"abc","namespace":"abc"}}],"kind":"List"}`),
		},
		{
			name: "delete in list with pattern only",
			args: args{
				data:    jsonList,
				pattern: pattern,
			},
			want: newStringValHolder(`{"apiVersion":"v1","items":[{"apiVersion":"v1","data":{},"kind":"service","metadata":{"name":"abc","namespace":"abc"}},{"apiVersion":"v1","data":{},"kind":"service","metadata":{"name":"def","namespace":"abc"}}],"kind":"List"}`),
		},
		{
			name: "delete in object with filter and pattern",
			args: args{
				data:    json,
				filter:  filter,
				pattern: pattern,
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {}}`),
		},
		{
			name: "delete in object with pattern only",
			args: args{
				data:    json,
				pattern: pattern,
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {}}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handleKubeJsonDelete(tt.args.data, tt.args.filter, tt.args.pattern)
			if got.dataType != tt.want.dataType {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			var gotstruct map[string]interface{}
			var wantStruct map[string]interface{}
			err := json2.Unmarshal([]byte(got.value.(string)), &gotstruct)
			if err != nil {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			err = json2.Unmarshal([]byte(tt.want.value.(string)), &wantStruct)
			if err != nil {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotstruct, wantStruct) {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleKubeYamlDelete(t *testing.T) {
	ymlList := `
apiVersion: v1
kind: List
items:
- apiVersion: v1
  kind: service
  metadata:
    name: abc
    namespace: abc
  data:
    school: abc
- apiVersion: v1
  kind: service
  metadata:
    name: def
    namespace: abc
  data:
    school: def`
	//outYmlist := `{"apiVersion":"v1","items":[{"apiVersion":"v1","data":{"school":"abc"},"kind":"service","metadata":{"name":"abc","namespace":"abc"}},{"apiVersion":"v1","data":{},"kind":"service","metadata":{"name":"def","namespace":"abc"}}],"kind":"List"}`
	ymls := `
apiVersion: v1
kind: service
metadata:
  name: abc
  namespace: abc
data:
  school: abc
---
apiVersion: v1
kind: service
metadata:
  name: def
  namespace: abc
data:
  school: def
`
	yml := `
apiVersion: v1
kind: service
metadata:
  name: def
  namespace: abc
data:
  school: def`
	pattern := "data.school"
	o := unstructured.Unstructured{}
	o.UnmarshalJSON([]byte(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "def"}}`))
	resourceKey := kube.GetResourceKey(&o)
	filter := (&resourceKey).String()
	type args struct {
		data    string
		filter  string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want valHolder
	}{
		{
			name: "delete in list with filter and pattern",
			args: args{
				data:    ymlList,
				filter:  filter,
				pattern: pattern,
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "List", "items":[{"apiVersion": "v1", "kind": "service", "metadata": {"name": "abc", "namespace": "abc"}, "data": {"school": "abc"}}, {"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {}}]}`),
		},
		{
			name: "delete in list with filter only",
			args: args{
				data:   ymlList,
				filter: filter,
			},
			want: newStringValHolder(`{"apiVersion":"v1","items":[{"apiVersion":"v1","data":{"school":"abc"},"kind":"service","metadata":{"name":"abc","namespace":"abc"}}],"kind":"List"}`),
		},
		{
			name: "delete in list with pattern only",
			args: args{
				data:    ymlList,
				pattern: pattern,
			},
			want: newStringValHolder(`{"apiVersion":"v1","items":[{"apiVersion":"v1","data":{},"kind":"service","metadata":{"name":"abc","namespace":"abc"}},{"apiVersion":"v1","data":{},"kind":"service","metadata":{"name":"def","namespace":"abc"}}],"kind":"List"}`),
		},
		{
			name: "delete in yamls with filter and pattern",
			args: args{
				data:    ymls,
				filter:  filter,
				pattern: pattern,
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "List", "items":[{"apiVersion": "v1", "kind": "service", "metadata": {"name": "abc", "namespace": "abc"}, "data": {"school": "abc"}}, {"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {}}]}`),
		},
		{
			name: "delete in yamls with filter only",
			args: args{
				data:   ymls,
				filter: filter,
			},
			want: newStringValHolder(`{"apiVersion":"v1","data":{"school":"abc"},"kind":"service","metadata":{"name":"abc","namespace":"abc"}}`),
		},
		{
			name: "delete in yamls with pattern only",
			args: args{
				data:    ymls,
				pattern: pattern,
			},
			want: newStringValHolder(`{"apiVersion":"v1","items":[{"apiVersion":"v1","data":{},"kind":"service","metadata":{"name":"abc","namespace":"abc"}},{"apiVersion":"v1","data":{},"kind":"service","metadata":{"name":"def","namespace":"abc"}}],"kind":"List"}`),
		},
		{
			name: "delete in object with filter and pattern",
			args: args{
				data:    yml,
				filter:  filter,
				pattern: pattern,
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {}}`),
		},
		{
			name: "delete in object with pattern only",
			args: args{
				data:    yml,
				pattern: pattern,
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {}}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handleKubeYamlDelete(tt.args.data, tt.args.filter, tt.args.pattern)
			if got.dataType != tt.want.dataType {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			oy, err := yaml.YAMLToJSON([]byte(got.value.(string)))
			if err != nil {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			yamls := strings.Split(got.value.(string), yamlSeperator)
			if len(yamls) > 1 {
				var jsons []string
				for _, y := range yamls {
					j, err := yaml.YAMLToJSON([]byte(y))
					if err != nil {
						t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
					}
					jsons = append(jsons, string(j))
				}
				oj := strings.Join(jsons, ",")
				oy = []byte(`{"apiVersion": "v1",    "items": [` + oj + `], "kind": "List"}`)
			}

			var gotstruct map[string]interface{}
			var wantStruct map[string]interface{}
			err = json2.Unmarshal(oy, &gotstruct)
			if err != nil {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			err = json2.Unmarshal([]byte(tt.want.value.(string)), &wantStruct)
			if err != nil {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			diff := cmp.Diff(gotstruct, wantStruct)
			fmt.Printf("%+v\n", diff)
			if !cmp.Equal(gotstruct, wantStruct) {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", gotstruct, wantStruct)
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleKubeJsonEdit(t *testing.T) {
	jsonList := `{"apiVersion": "v1", "kind": "List", "items":[{"apiVersion": "v1", "kind": "service", "metadata": {"name": "abc", "namespace": "abc"}, "data": {"school": "abc"}}, {"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "def"}}]}`
	json := `{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "def"}}`
	pattern := "data.school"
	o := unstructured.Unstructured{}
	o.UnmarshalJSON([]byte(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "def"}}`))
	resourceKey := kube.GetResourceKey(&o)
	filter := (&resourceKey).String()
	type args struct {
		data    string
		filter  string
		pattern string
		value   interface{}
	}
	tests := []struct {
		name string
		args args
		want valHolder
	}{
		{
			name: "edit in list with filter and pattern",
			args: args{
				data:    jsonList,
				filter:  filter,
				pattern: pattern,
				value:   "ghi",
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "List", "items":[{"apiVersion": "v1", "kind": "service", "metadata": {"name": "abc", "namespace": "abc"}, "data": {"school": "abc"}}, {"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "ghi"}}]}`),
		},
		{
			name: "edit in list with pattern only",
			args: args{
				data:    jsonList,
				pattern: pattern,
				value:   "ghi",
			},
			want: newStringValHolder(`{"apiVersion":"v1","items":[{"apiVersion":"v1","data":{"school": "ghi"},"kind":"service","metadata":{"name":"abc","namespace":"abc"}}, {"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "ghi"}}],"kind":"List"}`),
		},
		{
			name: "edit in object with filter and pattern",
			args: args{
				data:    json,
				filter:  filter,
				pattern: pattern,
				value:   "ghi",
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "ghi"}}`),
		},
		{
			name: "edit in object with pattern only",
			args: args{
				data:    json,
				pattern: pattern,
				value:   "ghi",
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "ghi"}}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handleKubeJsonEdit(tt.args.data, tt.args.filter, tt.args.pattern, tt.args.value)
			if got.dataType != tt.want.dataType {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			var gotstruct map[string]interface{}
			var wantStruct map[string]interface{}
			err := json2.Unmarshal([]byte(got.value.(string)), &gotstruct)
			if err != nil {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			err = json2.Unmarshal([]byte(tt.want.value.(string)), &wantStruct)
			if err != nil {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotstruct, wantStruct) {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleKubeYamlEdit(t *testing.T) {
	ymlList := `
apiVersion: v1
kind: List
items:
- apiVersion: v1
  kind: service
  metadata:
    name: abc
    namespace: abc
  data:
    school: abc
- apiVersion: v1
  kind: service
  metadata:
    name: def
    namespace: abc
  data:
    school: def`
	//outYmlist := `{"apiVersion":"v1","items":[{"apiVersion":"v1","data":{"school":"abc"},"kind":"service","metadata":{"name":"abc","namespace":"abc"}},{"apiVersion":"v1","data":{},"kind":"service","metadata":{"name":"def","namespace":"abc"}}],"kind":"List"}`
	ymls := `
apiVersion: v1
kind: service
metadata:
  name: abc
  namespace: abc
data:
  school: abc
---
apiVersion: v1
kind: service
metadata:
  name: def
  namespace: abc
data:
  school: def
`
	yml := `
apiVersion: v1
kind: service
metadata:
  name: def
  namespace: abc
data:
  school: def`
	pattern := "data.school"
	o := unstructured.Unstructured{}
	o.UnmarshalJSON([]byte(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "def"}}`))
	resourceKey := kube.GetResourceKey(&o)
	filter := (&resourceKey).String()
	type args struct {
		data    string
		filter  string
		pattern string
		value   interface{}
	}
	tests := []struct {
		name string
		args args
		want valHolder
	}{
		{
			name: "edit in list with filter and pattern",
			args: args{
				data:    ymlList,
				filter:  filter,
				pattern: pattern,
				value:   "ghi",
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "List", "items":[{"apiVersion": "v1", "kind": "service", "metadata": {"name": "abc", "namespace": "abc"}, "data": {"school": "abc"}}, {"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "ghi"}}]}`),
		},
		{
			name: "edit in list with pattern only",
			args: args{
				data:    ymlList,
				pattern: pattern,
				value:   "ghi",
			},
			want: newStringValHolder(`{"apiVersion":"v1","items":[{"apiVersion":"v1","data":{"school": "ghi"},"kind":"service","metadata":{"name":"abc","namespace":"abc"}},{"apiVersion":"v1","data":{"school": "ghi"},"kind":"service","metadata":{"name":"def","namespace":"abc"}}],"kind":"List"}`),
		},
		{
			name: "edit in yamls with filter and pattern",
			args: args{
				data:    ymls,
				filter:  filter,
				pattern: pattern,
				value:   "ghi",
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "List", "items":[{"apiVersion": "v1", "kind": "service", "metadata": {"name": "abc", "namespace": "abc"}, "data": {"school": "abc"}}, {"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "ghi"}}]}`),
		},
		{
			name: "edit in yamls with pattern only",
			args: args{
				data:    ymls,
				pattern: pattern,
				value:   "ghi",
			},
			want: newStringValHolder(`{"apiVersion":"v1","items":[{"apiVersion":"v1","data":{"school": "ghi"},"kind":"service","metadata":{"name":"abc","namespace":"abc"}},{"apiVersion":"v1","data":{"school": "ghi"},"kind":"service","metadata":{"name":"def","namespace":"abc"}}],"kind":"List"}`),
		},
		{
			name: "edit in object with filter and pattern",
			args: args{
				data:    yml,
				filter:  filter,
				pattern: pattern,
				value:   "ghi",
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "ghi"}}`),
		},
		{
			name: "edit in object with pattern only",
			args: args{
				data:    yml,
				pattern: pattern,
				value:   "ghi",
			},
			want: newStringValHolder(`{"apiVersion": "v1", "kind": "service", "metadata": {"name": "def", "namespace": "abc"}, "data": {"school": "ghi"}}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handleKubeYamlEdit(tt.args.data, tt.args.filter, tt.args.pattern, tt.args.value)
			if got.dataType != tt.want.dataType {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			oy, err := yaml.YAMLToJSON([]byte(got.value.(string)))
			if err != nil {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			yamls := strings.Split(got.value.(string), yamlSeperator)
			if len(yamls) > 1 {
				var jsons []string
				for _, y := range yamls {
					j, err := yaml.YAMLToJSON([]byte(y))
					if err != nil {
						t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
					}
					jsons = append(jsons, string(j))
				}
				oj := strings.Join(jsons, ",")
				oy = []byte(`{"apiVersion": "v1",    "items": [` + oj + `], "kind": "List"}`)
			}

			var gotstruct map[string]interface{}
			var wantStruct map[string]interface{}
			err = json2.Unmarshal(oy, &gotstruct)
			if err != nil {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			err = json2.Unmarshal([]byte(tt.want.value.(string)), &wantStruct)
			if err != nil {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
			diff := cmp.Diff(gotstruct, wantStruct)
			fmt.Printf("%+v\n", diff)
			if !cmp.Equal(gotstruct, wantStruct) {
				t.Errorf("handleKubeJsonDelete() = %v, want %v", gotstruct, wantStruct)
				t.Errorf("handleKubeJsonDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
