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
	"github.com/argoproj/gitops-engine/pkg/utils/kube"
	"github.com/devtron-labs/inception/pkg/language/parser"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
	"strings"
)

// ExitJson_delete_fn is called when production json_delete_fn is exited.
func (l *KlangListener) ExitKube_json_delete_fn(ctx *parser.Kube_json_delete_fnContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleKube_json_delete_fn(ctx)
}

func (l *KlangListener) handleKube_json_delete_fn(ctx *parser.Kube_json_delete_fnContext) valHolder {
	json := l.values[ctx.ID().GetText()]
	json = l.getValIfID(json)
	if json.dataType != STRING || len(json.value.(string)) == 0 {
		return newErrHolder(fmt.Errorf("json should be string of non zero length %+v\n", json))
	}
	data := json.value.(string)
	pattern := ""
	if ctx.Pattern() != nil {
		pattern = l.GetTextFromStringOrId(ctx.Pattern().(*parser.PatternContext).String_or_id().(*parser.String_or_idContext))
	}
	filter := ""
	if ctx.Filter() != nil {
		filter = l.GetTextFromStringOrId(ctx.Filter().(*parser.FilterContext).String_or_id().(*parser.String_or_idContext))
		filter = StripQuotes(filter)
	}
	res := handleKubeJsonDelete(data, filter, pattern)
	json.value = res.value
	l.values[json.name] = json
	return res
}

func handleKubeJsonDelete(data string, filter string, pattern string) valHolder {
	obj := &unstructured.Unstructured{}
	err := obj.UnmarshalJSON([]byte(data))
	if err != nil {
		return newStringValHolder(data)
	}
	if obj.IsList() {
		objs, err := obj.ToList()
		if err != nil {
			log.Printf("error while converting to list %+v\n", err)
		}
		var items []interface{}
		for _, o := range objs.Items {
			resourceKey := kube.GetResourceKey(&o)
			if len(filter) == 0 || (&resourceKey).String() == filter {
				if len(pattern) == 0 {
					//If pattern is missing
					//this object will not be appended in final result
					continue
				} else {
					json, err := o.MarshalJSON()
					if err != nil {
						return newStringValHolder(data)
					}
					out := JsonDelete(string(json), pattern)
					if out.dataType != STRING || out.value == nil {
						return newStringValHolder(data)
					}
					oUnStruct := &unstructured.Unstructured{}
					err = (oUnStruct).UnmarshalJSON([]byte(out.value.(string)))
					if err != nil {
						return newStringValHolder(data)
					}
					items = append(items, oUnStruct.Object)
				}
			} else {
				//if filter is not matching then append
				items = append(items, o.Object)
			}
		}
		obj.Object["items"] = items
		out, err := obj.MarshalJSON()
		if err != nil {
			return newStringValHolder(data)
		}
		return newStringValHolder(string(out))
	} else {
		resourceKey := kube.GetResourceKey(obj)
		rKey := (&resourceKey).String()
		if len(filter) == 0 || rKey == filter {
			if len(pattern) == 0 {
				return newEmptyHolder()
			} else {
				out := JsonDelete(data, pattern)
				return out
			}
		} else {
			return newStringValHolder(data)
		}
	}
	return newStringValHolder(data)
}

// ExitKube_json_edit_fn is called when production kube_yaml_edit_fn is exited.
func (l *KlangListener) ExitKube_json_edit_fn(ctx *parser.Kube_json_edit_fnContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleKube_json_edit_fn(ctx)
}

func (l *KlangListener) handleKube_json_edit_fn(ctx *parser.Kube_json_edit_fnContext) valHolder {
	json := l.values[ctx.ID().GetText()]
	json = l.getValIfID(json)
	if json.dataType != STRING || len(json.value.(string)) == 0 {
		return newErrHolder(fmt.Errorf("json should be string of non zero length %+v\n", json))
	}
	data := json.value.(string)
	valVh := l.handleExpr(ctx.Expr())
	valVh = l.getValIfID(valVh)
	if valVh.value == nil {
		return newErrHolder(fmt.Errorf("value cannot be nil"))
	}
	val := valVh.value
	patternLabel := ctx.String_or_id(0).(*parser.String_or_idContext)
	pattern := l.GetTextFromStringOrId(patternLabel)
	filter := ""
	if len(ctx.AllString_or_id()) > 1 {
		filter = l.GetTextFromStringOrId(ctx.String_or_id(1).(*parser.String_or_idContext))
	}
	if valVh.dataType == STRING {
		val = convertToInterface(val.(string))
	}
	res := handleKubeJsonEdit(data, filter, pattern, val)
	json.value = res.value
	l.values[json.name] = json
	return res
}

func handleKubeJsonEdit(data string, filter string, pattern string, value interface{}) valHolder {
	obj := &unstructured.Unstructured{}
	err := obj.UnmarshalJSON([]byte(data))
	if err != nil {
		return newStringValHolder(data)
	}
	//in := convertToInterface(data)
	if obj.IsList() {
		objs, err := obj.ToList()
		if err != nil {
			log.Printf("error while converting to list %+v\n", err)
		}
		var items []interface{}
		for _, o := range objs.Items {
			resourceKey := kube.GetResourceKey(&o)
			if len(filter) == 0 || (&resourceKey).String() == filter {
				json, err := o.MarshalJSON()
				if err != nil {
					return newStringValHolder(data)
				}
				out := JsonEdit(string(json), pattern, value)
				if out.dataType != STRING || out.value == nil {
					return newStringValHolder(data)
				}
				oUnStruct := &unstructured.Unstructured{}
				err = (oUnStruct).UnmarshalJSON([]byte(out.value.(string)))
				if err != nil {
					return newStringValHolder(data)
				}
				items = append(items, oUnStruct.Object)
			} else {
				//if filter is not matching then append
				items = append(items, o.Object)
			}
		}
		obj.Object["items"] = items
		out, err := obj.MarshalJSON()
		if err != nil {
			return newStringValHolder(data)
		}
		return newStringValHolder(string(out))
	} else {
		resourceKey := kube.GetResourceKey(obj)
		rKey := (&resourceKey).String()
		if len(filter) == 0 || rKey == filter {
			out := JsonEdit(data, pattern, value)
			return out
		} else {
			return newStringValHolder(data)
		}
	}
	return newStringValHolder(data)
}

func (l *KlangListener) ExitKube_yaml_delete_fn(ctx *parser.Kube_yaml_delete_fnContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleKube_yaml_delete_fn(ctx)
}

func (l *KlangListener) handleKube_yaml_delete_fn(ctx *parser.Kube_yaml_delete_fnContext) valHolder {
	yml := l.values[ctx.ID().GetText()]
	yml = l.getValIfID(yml)
	if yml.dataType != STRING || len(yml.value.(string)) == 0 {
		return newErrHolder(fmt.Errorf("yml should be string of non zero length %+v\n", yml))
	}
	data := yml.value.(string)
	pattern := ""
	if ctx.Pattern() != nil {
		pattern = l.GetTextFromStringOrId(ctx.Pattern().(*parser.PatternContext).String_or_id().(*parser.String_or_idContext))
	}
	filter := ""
	if ctx.Filter() != nil {
		filter = l.GetTextFromStringOrId(ctx.Filter().(*parser.FilterContext).String_or_id().(*parser.String_or_idContext))
	}

	res := handleKubeYamlDelete(data, filter, pattern)
	yml.value = res.value
	l.values[yml.name] = yml
	return yml
}

func handleKubeYamlDelete(data string, filter string, pattern string) valHolder {
	ymls := strings.Split(data, yamlSeperator)
	var outYmls []string
	for _, y := range ymls {
		json, err := yaml.YAMLToJSON([]byte(y))
		if err != nil {
			return newStringValHolder(data)
		}
		out := handleKubeJsonDelete(string(json), filter, pattern)
		if out.value != nil && len(out.value.(string)) > 0 {
			outYml, err := yaml.JSONToYAML([]byte(out.value.(string)))
			if err != nil {
				return newStringValHolder(data)
			}
			outYmls = append(outYmls, string(outYml))
		}
		//If return is empty string because of filter only then skip it
	}
	return newStringValHolder(strings.Join(outYmls, yamlSeperator))
}

// ExitKube_yaml_edit_fn is called when production kube_yaml_edit_fn is exited.
func (l *KlangListener) ExitKube_yaml_edit_fn(ctx *parser.Kube_yaml_edit_fnContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleKube_yaml_edit_fn(ctx)
}

func (l *KlangListener) handleKube_yaml_edit_fn(ctx *parser.Kube_yaml_edit_fnContext) valHolder {
	yml := l.values[ctx.ID().GetText()]
	yml = l.getValIfID(yml)
	if yml.dataType != STRING || len(yml.value.(string)) == 0 {
		return newErrHolder(fmt.Errorf("yml should be string of non zero length %+v\n", yml))
	}
	data := yml.value.(string)
	valVh := l.handleExpr(ctx.Expr())
	valVh = l.getValIfID(valVh)
	if valVh.value == nil {
		return newErrHolder(fmt.Errorf("value cannot be nil"))
	}
	val := valVh.value
	patternLabel := ctx.String_or_id(0).(*parser.String_or_idContext)
	pattern := l.GetTextFromStringOrId(patternLabel)
	filter := ""
	if len(ctx.AllString_or_id()) > 1 {
		filter = l.GetTextFromStringOrId(ctx.String_or_id(1).(*parser.String_or_idContext))
	}
	if valVh.dataType == STRING {
		val = convertToInterface(val.(string))
	}
	res := handleKubeYamlEdit(data, filter, pattern, val)
	yml.value = res.value
	l.values[yml.name] = yml
	return res
}

func handleKubeYamlEdit(data string, filter string, pattern string, value interface{}) valHolder {
	ymls := strings.Split(data, yamlSeperator)
	var outYmls []string
	for _, y := range ymls {
		json, err := yaml.YAMLToJSON([]byte(y))
		if err != nil {
			return newStringValHolder(data)
		}
		out := handleKubeJsonEdit(string(json), filter, pattern, value)
		if out.value != nil && len(out.value.(string)) > 0 {
			outYml, err := yaml.JSONToYAML([]byte(out.value.(string)))
			if err != nil {
				return newStringValHolder(data)
			}
			outYmls = append(outYmls, string(outYml))
		}
		//If return is empty string because of filter only then skip it
	}
	return newStringValHolder(strings.Join(outYmls, yamlSeperator))
}

func convertToInterface(data string) interface{} {
	var out map[string]interface{}
	err := json.Unmarshal([]byte(data), &out)
	if err == nil {
		return out
	}
	err = yaml.Unmarshal([]byte(data), &out)
	if err == nil {
		return out
	}
	return data
}
