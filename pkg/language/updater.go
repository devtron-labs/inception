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
	"encoding/json"
	"fmt"
	"github.com/devtron-labs/inception/pkg/language/flatten"
	"github.com/ghodss/yaml"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func updateKubernetesObjectsJson(from, to string, patterns []string) string {
	var originalTo string
	var originalToMap map[string]interface{}
	err := json.Unmarshal([]byte(to), &originalToMap)
	if err != nil {
		fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
		return to
	}
	originalToByte, err := json.Marshal(originalToMap)
	if err != nil {
		fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
		return to
	}
	originalTo = string(originalToByte)
	for _, pattern := range patterns {
		v1 := gjson.Get(from, pattern)
		v2 := gjson.Get(to, pattern)
		if !equalResult(v1, v2) {
			return originalTo
		}
	}
	data := map[string]interface{}{}
	err = json.Unmarshal([]byte(from), &data)
	if err != nil {
		fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
		return originalTo
	}
	if m, ok := data["update"]; ok {
		update := m.(map[string]interface{})
		updatableJson, err := flatten.Flatten(update, "", flatten.DotStyle)
		if err != nil {
			fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
			return originalTo
		}
		for k, v := range updatableJson {
			to, err = sjson.Set(to, k, v)
			if err != nil {
				fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
				return originalTo
			}
		}
	}
	return to
}

func updateKubernetesObjectsYaml(from, to string, patterns []string) string {
	var originalTo string
	var originalToMap map[string]interface{}
	err := yaml.Unmarshal([]byte(to), &originalToMap)
	if err != nil {
		fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
		return to
	}
	originalToByte, err := yaml.Marshal(originalToMap)
	if err != nil {
		fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
		return to
	}
	originalTo = string(originalToByte)
	toJsonB, err := yaml.YAMLToJSON([]byte(to))
	if err != nil {
		fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
		return to
	}
	toJson := string(toJsonB)
	fromJsonB, err := yaml.YAMLToJSON([]byte(from))
	if err != nil {
		fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
		return to
	}
	fromJson := string(fromJsonB)
	for _, pattern := range patterns {
		v1 := gjson.Get(fromJson, pattern)
		v2 := gjson.Get(toJson, pattern)
		if !equalResult(v1, v2) {
			return originalTo
		}
	}
	data := map[string]interface{}{}
	err = yaml.Unmarshal([]byte(from), &data)
	if err != nil {
		fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
		return originalTo
	}
	if m, ok := data["update"]; ok {
		update := m.(map[string]interface{})
		updatableJson, err := flatten.Flatten(update, "", flatten.DotStyle)
		if err != nil {
			fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
			return originalTo
		}
		for k, v := range updatableJson {
			toJson, err = sjson.Set(toJson, k, v)
			if err != nil {
				fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
				return originalTo
			}
		}
	}
	toYaml, err := yaml.JSONToYAML([]byte(toJson))
	if err != nil {
		fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
		return originalTo
	}
	return string(toYaml)
}
