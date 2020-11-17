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
	"github.com/tidwall/gjson"
	"sigs.k8s.io/yaml"
)

func YamlSelect(yml, pattern string) valHolder {
	json, err := yaml.YAMLToJSON([]byte(yml))
	if err != nil {
		return newErrHolder(err)
	}
	val := gjson.Get(string(json), pattern)
	vh := toValHolder(val)
	if val.Type == gjson.JSON {
		ry, err := yaml.JSONToYAML([]byte(val.Raw))
		if err != nil {
			return newErrHolder(err)
		}
		vh.value = string(ry)
	}
	return vh
}

func YamlEdit(yml, pattern string, value interface{}) valHolder {
	json, err := yaml.YAMLToJSON([]byte(yml))
	if err != nil {
		return newErrHolder(err)
	}
	r := JsonEdit(string(json), pattern, value)
	ry, err := yaml.JSONToYAML([]byte(r.value.(string)))
	if err != nil {
		return newErrHolder(err)
	}
	return newStringValHolder(string(ry))
}

func YamlDelete(yml, pattern string) valHolder {
	json, err := yaml.YAMLToJSON([]byte(yml))
	if err != nil {
		return newErrHolder(err)
	}
	r := JsonDelete(string(json), pattern)
	ry, err := yaml.JSONToYAML([]byte(r.value.(string)))
	if err != nil {
		return newErrHolder(err)
	}
	return newStringValHolder(string(ry))
}
