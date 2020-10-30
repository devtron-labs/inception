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
	"reflect"
	"testing"
)

func TestYamlSelect(t *testing.T) {
	yml := `
items:
- metadata:
    labels:
      first: abc
    name: efg
- metadata:
    labels:
      first: def
    name: hij
`
	type args struct {
		yml     string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want valHolder
	}{
		{
			name: "yaml select",
			args: args{
				yml:     yml,
				pattern: `items.#(metadata.labels.first="def").metadata.name`,
			},
			want: valHolder{
				dataType: STRING,
				value:    "hij",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YamlSelect(tt.args.yml, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YamlSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
