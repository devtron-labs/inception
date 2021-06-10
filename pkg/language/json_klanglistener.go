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
)

func (l *KlangListener) handleJson_select_Fn(ctx parser.IJson_select_fnContext) valHolder {
	jctx := ctx.(*parser.Json_select_fnContext)
	id := jctx.ID().GetText()
	patternLabel := jctx.String_or_id().(*parser.String_or_idContext)
	pattern := l.GetTextFromStringOrId(patternLabel)
	if data, ok := l.values[id]; ok && data.dataType == STRING && len(pattern) != 0 {
		res := JsonSelect(data.value.(string), pattern)
		return res
	}
	return newEmptyHolder()
}

// ExitJson_delete_fn is called when production json_delete_fn is exited.
func (l *KlangListener) ExitJson_delete_fn(ctx *parser.Json_delete_fnContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleJson_delete_fn(ctx)
}

func (l *KlangListener) handleJson_delete_fn(ctx *parser.Json_delete_fnContext) valHolder {
	json := l.values[ctx.ID().GetText()]
	json = l.getValIfID(json)
	if json.dataType != STRING || len(json.value.(string)) == 0 {
		return newErrHolder(fmt.Errorf("json should be string of non zero length %+v\n", json))
	}
	patternLabel := ctx.String_or_id().(*parser.String_or_idContext)
	pattern := l.GetTextFromStringOrId(patternLabel)
	res := JsonDelete(json.value.(string), pattern)
	json.value = res.value
	l.values[json.name] = json
	return newEmptyHolder()
}

// ExitJsonselector_assignment is called when production jsonselector_assignment is exited.
func (l *KlangListener) ExitJson_edit_fn(ctx *parser.Json_edit_fnContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleJson_edit_fn(ctx)
}

func (l *KlangListener) handleJson_edit_fn(ctx *parser.Json_edit_fnContext) valHolder {
	json := l.values[ctx.ID().GetText()]
	json = l.getValIfID(json)
	if json.dataType != STRING || len(json.value.(string)) == 0 {
		return newErrHolder(fmt.Errorf("json should be string of non zero length %+v\n", json))
	}
	valVh := l.handleExpr(ctx.Expr())
	valVh = l.getValIfID(valVh)
	if valVh.value == nil {
		return newErrHolder(fmt.Errorf("value cannot be nil"))
	}
	val := valVh.value
	patternLabel := ctx.String_or_id().(*parser.String_or_idContext)
	pattern := l.GetTextFromStringOrId(patternLabel)
	res := JsonEdit(json.value.(string), pattern, val)
	json.value = res.value
	l.values[json.name] = json
	return newEmptyHolder()
}
