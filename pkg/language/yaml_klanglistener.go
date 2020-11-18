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
	"strconv"
	"strings"
)

func (l *KlangListener) handleYaml_select_fn(ctx parser.IYaml_select_fnContext) valHolder {
	yctx := ctx.(*parser.Yaml_select_fnContext)
	id := yctx.ID().GetText()
	index := -1
	if yctx.NUMBER() != nil && len(yctx.NUMBER().GetText()) > 0 {
		i, err := strconv.Atoi(yctx.NUMBER().GetText())
		if err == nil {
			index = i
		}
	}

	patternLabel := yctx.String_or_id().(*parser.String_or_idContext)
	pattern := l.GetTextFromStringOrId(patternLabel)
	if data, ok := l.values[id]; ok && data.dataType == STRING && len(pattern) != 0 {
		yamls := strings.Split(data.value.(string), yamlSeperator)
		if index == -1 && len(yamls) != 1 {
			return newErrHolder(fmt.Errorf("in case of multiyaml (len is %d) its important to define index of yaml for selection", len(yamls)))
		}
		for i, yaml := range yamls {
			if index == -1 || i == index {
				res := YamlSelect(yaml, pattern)
				return res
			}
		}
	}
	return newNilValHolder()
}

// ExitYaml_delete_fn is called when production yaml_delete_fn is exited.
func (l *KlangListener) ExitYaml_delete_fn(ctx *parser.Yaml_delete_fnContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleYaml_delete_fn(ctx)
}

func (l *KlangListener) handleYaml_delete_fn(ctx *parser.Yaml_delete_fnContext) valHolder {
	yaml := l.values[ctx.ID().GetText()]
	yaml = l.getValIfID(yaml)
	if yaml.dataType != STRING || len(yaml.value.(string)) == 0 {
		return newErrHolder(fmt.Errorf("yaml should be string of non zero length found %+v\n", yaml))
	}
	patternLabel := ctx.String_or_id().(*parser.String_or_idContext)
	pattern := l.GetTextFromStringOrId(patternLabel)
	index := -1
	if ctx.NUMBER() != nil && len(ctx.NUMBER().GetText()) > 0 {
		i, err := strconv.Atoi(ctx.NUMBER().GetText())
		if err == nil {
			index = i
		}
	}
	yamls := strings.Split(yaml.value.(string), yamlSeperator)
	//yamls, err := kube.SplitYAML([]byte(yaml.value.(string)))
	var outYamls []string
	for i, yaml := range yamls {
		if index == -1 || i == index {
			res := YamlDelete(yaml, pattern)
			outYamls = append(outYamls, res.value.(string))
		} else {
			outYamls = append(outYamls, yaml)
		}
	}
	yaml.value = strings.Join(outYamls, yamlSeperator)
	l.values[yaml.name] = yaml
	//fmt.Printf("delete pattern %s, out %+v\n", pattern, yaml)
	return newEmptyHolder()
}

// ExitYamledit_fn is called when production yamledit_fn is exited.
func (l *KlangListener) ExitYaml_edit_fn(ctx *parser.Yaml_edit_fnContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleYaml_edit_fn(ctx)
}

func (l *KlangListener) handleYaml_edit_fn(ctx *parser.Yaml_edit_fnContext) valHolder {
	yaml := l.values[ctx.ID().GetText()]
	yaml = l.getValIfID(yaml)
	if yaml.dataType != STRING || len(yaml.value.(string)) == 0 {
		return newErrHolder(fmt.Errorf("yaml should be string of non zero length found %+v\n", yaml))
	}
	patternLabel := ctx.String_or_id().(*parser.String_or_idContext)
	pattern := l.GetTextFromStringOrId(patternLabel)
	valVh := l.handleExpr(ctx.Expr())
	valVh = l.getValIfID(valVh)
	if valVh.value == nil {
		return newErrHolder(fmt.Errorf("value cannot be nil"))
	}
	val := valVh.value
	index := -1
	if ctx.NUMBER() != nil && len(ctx.NUMBER().GetText()) > 0 {
		i, err := strconv.Atoi(ctx.NUMBER().GetText())
		if err == nil {
			index = i
		}
	}
	yamls := strings.Split(yaml.value.(string), yamlSeperator)
	var outYamls []string
	for i, yaml := range yamls {
		if index == -1 || i == index {
			res := YamlEdit(yaml, pattern, val)
			outYamls = append(outYamls, res.value.(string))
		} else {
			outYamls = append(outYamls, yaml)
		}
	}
	yaml.value = strings.Join(outYamls, yamlSeperator)
	l.values[yaml.name] = yaml
	//fmt.Printf("val %+v, pattern %s, out %+v\n", val, pattern, yaml)
	return newEmptyHolder()
}