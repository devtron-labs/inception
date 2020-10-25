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
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/argoproj/gitops-engine/pkg/health"
	"github.com/argoproj/gitops-engine/pkg/utils/kube"
	"github.com/devtron-labs/inception/pkg/parser"
	"io"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/meta"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type valHolder struct {
	dataType DataType
	name     string
	value    interface{}
}

type DataType string

const (
	STRING  DataType = "STRING"
	BOOLEAN DataType = "BOOLEAN"
	INT     DataType = "INT"
	FLOAT   DataType = "FLOAT"
	NIL     DataType = "NIL"
	ID      DataType = "ID"
	ERR     DataType = "ERROR"
)

const (
	APPLY  string = "APPLY"
	PATCH  string = "PATCH"
	DELETE string = "DELETE"
)

const yamlSeperator = "\n---\n"

type Resource struct {
	Operation string
	Group     string
	Version   string
	Kind      string
	Namespace string
	Name      string
	Message   string
	Status    ResourceSyncStatusCode
	Health    *HealthStatus
}

type HealthStatus struct {
	Status  health.HealthStatusCode `json:"status,omitempty" protobuf:"bytes,1,opt,name=status"`
	Message string                  `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
}

type ResourceSyncStatusCode string

const (
	ResourceSyncStatusCodeUnknown   ResourceSyncStatusCode = "Unknown"
	ResourceSyncStatusCodeSynced    ResourceSyncStatusCode = "Synced"
	ResourceSyncStatusCodeOutOfSync ResourceSyncStatusCode = "OutOfSync"
)

type KlangListener struct {
	*parser.BaseKlangListener
	ifWhileCount        int
	values              map[string]valHolder
	kubernetesResources map[string][]Resource
	mapper              *Mapper
}

func NewKlangListener(mapper *Mapper) *KlangListener {
	return &KlangListener{
		ifWhileCount: 0,
		values:       make(map[string]valHolder, 0),
		mapper:       mapper,
		kubernetesResources: make(map[string][]Resource, 0),
	}
}

func (l *KlangListener) Values() map[string]valHolder {
	t := make(map[string]valHolder, len(l.values))
	for k, v := range l.values {
		t[k] = v
	}
	return t
}

func (l *KlangListener) KubernetesResources() map[string][]Resource {
	t := make(map[string][]Resource, len(l.kubernetesResources))
	for k, vs := range l.kubernetesResources {
		var holders []Resource
		for _, v := range vs {
			holders = append(holders, v)
		}
		t[k] = holders
	}
	return t
}

func newStringValHolder(val string) valHolder {
	return valHolder{dataType: STRING, value: val}
}

func newBooleanValHolder(val bool) valHolder {
	return valHolder{dataType: BOOLEAN, value: val}
}

func newIntValHolder(val int64) valHolder {
	return valHolder{dataType: INT, value: val}
}

func newFloatValHolder(val float64) valHolder {
	return valHolder{dataType: FLOAT, value: val}
}

func newNilValHolder() valHolder {
	return valHolder{dataType: NIL, value: "NIL"}
}

func newErrHolder(err error) valHolder {
	fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
	return valHolder{dataType: ERR, value: err}
}

func newEmptyHolder() valHolder {
	return valHolder{}
}

//ExitBlock is not implemented
func (l *KlangListener) handleBlock(ctx *parser.BlockContext) {
	for _, stat := range ctx.AllStat() {
		l.handleStat(stat)
	}
}

//ExitStat not implemented, called from ExitIf_stat and ExitWhile_stat
func (l *KlangListener) handleStat(ctx parser.IStatContext) {
	ts := ctx.(*parser.StatContext)
	if ts.Assignment() != nil {
		//ExitAssignment is implemented but is skipped in case of If_stat and While_stat
		// as it should be processed here depending on whether condition is true or not
		ta := ts.Assignment().(*parser.AssignmentContext)
		l.handleAssignment(ta)
	} else if ts.Json_edit_fn() != nil {
		//ExitJsonedit_assignment is implemented but is skipped in case of If_stat and While_stat
		// as it should be processed here depending on whether condition is true or not
		tja := ts.Json_edit_fn().(*parser.Json_edit_fnContext)
		l.handleJson_edit_fn(tja)
	} else if ts.Yaml_edit_fn() != nil {
		tya := ts.Yaml_edit_fn().(*parser.Yaml_edit_fnContext)
		l.handleYaml_edit_fn(tya)
	} else if ts.Json_delete_fn() != nil {
		//ExitJsondelete'
		//]_assignment is implemented but is skipped in case of If_stat and While_stat
		// as it should be processed here depending on whether condition is true or not
		tja := ts.Json_delete_fn().(*parser.Json_delete_fnContext)
		l.handleJson_delete_fn(tja)
	} else if ts.Yaml_delete_fn() != nil {
		tya := ts.Yaml_delete_fn().(*parser.Yaml_delete_fnContext)
		l.handleYaml_delete_fn(tya)
	} else if ts.Log() != nil {
		log := ts.Log().(*parser.LogContext)
		l.handleLog(log)
	} else if ts.If_stat() != nil {
		//skip as is handled by ExitIf_stat
	} else if ts.While_stat() != nil {
		//skip as is handled by ExitWhile_stat
	}
}

// ExitLog is called when production log is exited.
func (l *KlangListener) ExitLog(ctx *parser.LogContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleLog(ctx)
}

func (l *KlangListener) handleLog(ctx *parser.LogContext) {
	out := l.handleExpr(ctx.Expr())
	out = l.getValIfID(out)
	fmt.Println(out.value)
}

//ExitStat_block not implemented, called from ExitIf_stat and ExitWhile_stat
func (l *KlangListener) handleStat_block(ctx parser.IStat_blockContext) {
	switch v := ctx.(type) {
	case *parser.Stat_blockContext:
		if v.Block() != nil {
			tb := v.Block().(*parser.BlockContext)
			l.handleBlock(tb)
		} else if v.Stat() != nil {
			l.handleStat(v.Stat())
		}
	}
	//fmt.Printf("print %s\n", ctx.GetText())
}

// EnterIf_stat is called when production if_stat is entered.
func (l *KlangListener) EnterIf_stat(ctx *parser.If_statContext) {
	l.ifWhileCount++
}

// EnterWhile_stat is called when production while_stat is entered.
func (l *KlangListener) EnterWhile_stat(ctx *parser.While_statContext) {
	l.ifWhileCount++
}

// ExitWhile_stat is called when production while_stat is exited.
func (l *KlangListener) ExitWhile_stat(ctx *parser.While_statContext) {
	defer func() { l.ifWhileCount-- }()
	res := l.handleExpr(ctx.Expr())
	for res.value.(bool) {
		l.handleStat_block(ctx.Stat_block())
		res = l.handleExpr(ctx.Expr())
	}
}

// ExitIf_stat is called when production if_stat is exited.
func (l *KlangListener) ExitIf_stat(ctx *parser.If_statContext) {
	defer func() { l.ifWhileCount-- }()
	for _, cb := range ctx.AllCondition_block() {
		tcb := cb.(*parser.Condition_blockContext)
		res := l.handleCondition_block(tcb)
		if res.value.(bool) == true {
			if tcb.Stat_block() != nil {
				l.handleStat_block(tcb.Stat_block())
			}
			return
		}
	}
	if ctx.Stat_block() != nil {
		l.handleStat_block(ctx.Stat_block())
	}
}

// ExitAssignment is called when production assignment is exited.
func (l *KlangListener) ExitAssignment(ctx *parser.AssignmentContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleAssignment(ctx)
}

func (l *KlangListener) handleAssignment(ctx *parser.AssignmentContext) {
	if ctx.Expr() != nil {
		r := l.handleExpr(ctx.Expr())
		r = l.getValIfID(r)
		o := valHolder{
			dataType: r.dataType,
			name:     ctx.ID().GetText(),
			value:    r.value,
		}
		l.values[o.name] = o
	} else if ctx.Load_fn() != nil {
		r := newStringValHolder("")
		r.name = ctx.ID().GetText()
		loadContext := ctx.Load_fn().(*parser.Load_fnContext)
		fName := l.GetTextFromStringOrId(loadContext.String_or_id().(*parser.String_or_idContext))
		if len(fName) == 0 {
			l.values[r.name] = r
			return
		}
		data, err := ioutil.ReadFile(fName)
		if err != nil {
			l.values[r.name] = r
			return
		}
		r.value = string(data)
		l.values[r.name] = r
	}
}

// ExitShell_script is called when production shell_script is exited.
func (l *KlangListener) ExitShell_script(ctx *parser.Shell_scriptContext) {
	if l.ifWhileCount != 0 {
		return
	}
	l.handleShell_script(ctx)
}

func (l *KlangListener) handleShell_script(ctx *parser.Shell_scriptContext) valHolder {
	script := l.GetTextFromStringOrId(ctx.String_or_id().(*parser.String_or_idContext))
	if len(script) == 0 {
		return newEmptyHolder()
	}
	//TODO: go lang code to execute shell script
	//TODO: write script to file and use file location
	hasher := sha1.New()
	hasher.Write([]byte(script))
	hash := hex.EncodeToString(hasher.Sum(nil))
	file := fmt.Sprintf("/tmp/%s", hash)
	err := ioutil.WriteFile(file, []byte(script), 500)
	if err != nil {
		return newErrHolder(fmt.Errorf("unable to write content to file %s\n", file))
	}
	out, err := exec.Command("/bin/sh", file).Output()
	if err != nil {
		return newErrHolder(err)
	}
	return newStringValHolder(string(out))
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

func (l *KlangListener) GetTextFromStringOrId(stringOrId *parser.String_or_idContext) string {
	pattern := ""
	if stringOrId.ID() != nil {
		if pval, ok := l.values[stringOrId.ID().GetText()]; ok {
			if pval.dataType == ERR {
				return ""
			}
			pattern = pval.value.(string)
			return pattern
		}
	} else if stringOrId.STRING() != nil {
		pattern = stringOrId.STRING().GetText()
		return pattern[1 : len(pattern)-1]
	} else if stringOrId.RAW_STRING_LIT() != nil {
		pattern = stringOrId.RAW_STRING_LIT().GetText()
		return pattern[1 : len(pattern)-1]
	}
	return pattern
}

// EnterYaml_delete_fn is called when production yaml_delete_fn is entered.
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
	fmt.Printf("delete pattern %s, out %+v\n", pattern, yaml)
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
	fmt.Printf("val %+v, pattern %s, out %+v\n", val, pattern, yaml)
	return newEmptyHolder()
}

func (l *KlangListener) handleDownload_fn(ctx *parser.Download_fnContext) valHolder {
	url := l.GetTextFromStringOrId(ctx.String_or_id(0).(*parser.String_or_idContext))
	fileName := ""
	if ctx.String_or_id(1) != nil {
		fileName = l.GetTextFromStringOrId(ctx.String_or_id(1).(*parser.String_or_idContext))
	} else {
		hasher := sha1.New()
		hasher.Write([]byte(url))
		hash := hex.EncodeToString(hasher.Sum(nil))
		fileName = fmt.Sprintf("/tmp/%s", hash)
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return newErrHolder(err)
	}

	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		return newErrHolder(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return newErrHolder(err)
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return newErrHolder(err)
	}

	return newStringValHolder(string(data))
}

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

func (l *KlangListener) handleCondition_block(ctx *parser.Condition_blockContext) valHolder {
	res := l.handleExpr(ctx.Expr())
	//fmt.Printf("final result %v\n", res)
	return res
}

func (l *KlangListener) handleExpr(ctx parser.IExprContext) valHolder {
	switch v := ctx.(type) {
	case *parser.ExprContext:
		//fmt.Printf("%d\n", v.GetChildCount())
		break
	case *parser.AtomExprContext:
		return l.handleAtom(v.Atom())
	case *parser.EqualityExprContext:
		op := EQ
		if v.NEQ() != nil {
			op = NEQ
		}
		lhs := l.handleExpr(v.Expr(0))
		rhs := l.handleExpr(v.Expr(1))
		lhs = l.getValIfID(lhs)
		rhs = l.getValIfID(rhs)
		c1 := lhs.dataType == INT && rhs.dataType == INT && logicalIntIntOperation(lhs, rhs, op)
		c2 := lhs.dataType == FLOAT && rhs.dataType == FLOAT && logicalFloatFloatOperation(lhs, rhs, op)
		c3 := lhs.dataType == INT && rhs.dataType == FLOAT && logicalIntFloatOperation(lhs, rhs, op)
		c4 := lhs.dataType == FLOAT && rhs.dataType == INT && logicalFloatIntOperation(lhs, rhs, op)
		c5 := lhs.dataType == STRING && rhs.dataType == STRING && logicalStringStringOperation(lhs, rhs, op)
		c := c1 || c2 || c3 || c4 || c5
		return newBooleanValHolder(c)
	case *parser.KubectlExprContext:
		return l.handleKubectl_command(v.Kubectl_command())
	case *parser.JsonSelectFnContext:
		return l.handleJson_select_Fn(v.Json_select_fn())
	case *parser.YamlSelectFnContext:
		return l.handleYaml_select_fn(v.Yaml_select_fn())
	case *parser.RelationalExprContext:
		op := LT
		if v.LTEQ() != nil {
			op = LTEQ
		} else if v.GT() != nil {
			op = GT
		} else if v.GTEQ() != nil {
			op = GTEQ
		}
		lhs := l.handleExpr(v.Expr(0))
		rhs := l.handleExpr(v.Expr(1))
		lhs = l.getValIfID(lhs)
		rhs = l.getValIfID(rhs)
		c1 := lhs.dataType == INT && rhs.dataType == INT && logicalIntIntOperation(lhs, rhs, op)
		c2 := lhs.dataType == FLOAT && rhs.dataType == FLOAT && logicalFloatFloatOperation(lhs, rhs, op)
		c3 := lhs.dataType == INT && rhs.dataType == FLOAT && logicalIntFloatOperation(lhs, rhs, op)
		c4 := lhs.dataType == FLOAT && rhs.dataType == INT && logicalFloatIntOperation(lhs, rhs, op)
		c5 := lhs.dataType == STRING && rhs.dataType == STRING && logicalStringStringOperation(lhs, rhs, op)
		c := c1 || c2 || c3 || c4 || c5
		return newBooleanValHolder(c)
	case *parser.AndExprContext:
		lhs := l.handleExpr(v.Expr(0))
		rhs := l.handleExpr(v.Expr(1))
		lhs = l.getValIfID(lhs)
		rhs = l.getValIfID(rhs)
		if lhs.dataType != BOOLEAN || rhs.dataType != BOOLEAN {
			return valHolder{}
		}
		isTrue := lhs.value.(bool) == true && rhs.value.(bool) == true
		return newBooleanValHolder(isTrue)
	case *parser.OrExprContext:
		lhs := l.handleExpr(v.Expr(0))
		rhs := l.handleExpr(v.Expr(1))
		lhs = l.getValIfID(lhs)
		rhs = l.getValIfID(rhs)
		if lhs.dataType != BOOLEAN || rhs.dataType != BOOLEAN {
			return valHolder{}
		}
		isTrue := lhs.value.(bool) == true || rhs.value.(bool) == true
		return newBooleanValHolder(isTrue)
	case *parser.AdditiveExprContext:
		op := PLUS
		if v.MINUS() != nil {
			op = MINUS
		}
		lhs := l.handleExpr(v.Expr(0))
		rhs := l.handleExpr(v.Expr(1))
		lhs = l.getValIfID(lhs)
		rhs = l.getValIfID(rhs)
		if lhs.dataType == INT && rhs.dataType == INT {
			return mathematicalIntIntOperation(lhs, rhs, op)
		}
		if lhs.dataType == FLOAT && rhs.dataType == FLOAT {
			return mathematicalFloatFloatOperation(lhs, rhs, op)
		}
		if lhs.dataType == INT && rhs.dataType == FLOAT {
			return mathematicalIntFloatOperation(lhs, rhs, op)
		}
		if lhs.dataType == FLOAT && rhs.dataType == INT {
			return mathematicalFloatIntOperation(lhs, rhs, op)
		}
		if lhs.dataType == STRING && rhs.dataType == STRING {
			return mathematicalStringStringOperation(lhs, rhs, op)
		}
		return valHolder{}
	case *parser.MultiplicationExprContext:
		op := MULT
		if v.DIV() != nil {
			op = DIV
		} else if v.MOD() != nil {
			op = MOD
		}
		lhs := l.handleExpr(v.Expr(0))
		rhs := l.handleExpr(v.Expr(1))
		lhs = l.getValIfID(lhs)
		rhs = l.getValIfID(rhs)
		if lhs.dataType == INT && rhs.dataType == INT {
			return mathematicalIntIntOperation(lhs, rhs, op)
		}
		if lhs.dataType == FLOAT && rhs.dataType == FLOAT {
			return mathematicalFloatFloatOperation(lhs, rhs, op)
		}
		if lhs.dataType == INT && rhs.dataType == FLOAT {
			return mathematicalIntFloatOperation(lhs, rhs, op)
		}
		if lhs.dataType == FLOAT && rhs.dataType == INT {
			return mathematicalFloatIntOperation(lhs, rhs, op)
		}
		return valHolder{}
	case *parser.ShellScriptContext:
		return l.handleShell_script(v.Shell_script().(*parser.Shell_scriptContext))
	case *parser.DownloadFnContext:
		return l.handleDownload_fn(v.Download_fn().(*parser.Download_fnContext))
		//TODO: handledownload
	case *parser.NotExprContext:
		r := l.handleExpr(v.Expr())
		return l.isFalse(r)
	default:
		break
	}
	return valHolder{}
}

func (l *KlangListener) isFalse(r valHolder) valHolder {
	switch r.dataType {
	case INT:
		return newBooleanValHolder(r.value.(int64) == 0)
	case STRING:
		return newBooleanValHolder(len(r.value.(string)) == 0)
	case BOOLEAN:
		return newBooleanValHolder(!r.value.(bool))
	case FLOAT:
		return newBooleanValHolder(r.value.(float64) == 0.0)
	case NIL:
		return newBooleanValHolder(true)
	case ID:
		if val, ok := l.values[r.name]; ok {
			return l.isFalse(val)
		}
		return newBooleanValHolder(true)
	case ERR:
		return newBooleanValHolder(true)
	default:
		return newBooleanValHolder(true)
	}
}

func (l *KlangListener) getValIfID(val valHolder) valHolder {
	if val.dataType == ID {
		if v, ok := l.values[val.value.(string)]; ok {
			return v
		}
		return newEmptyHolder()
	}
	return val
}

func (l *KlangListener) handleAtom(ctx parser.IAtomContext) valHolder {
	switch v := ctx.(type) {
	case *parser.NumberAtomContext:
		if strings.Index(v.GetText(), ".") > 0 {
			val, err := strconv.ParseFloat(v.GetText(), 64)
			if err != nil {
				return valHolder{}
			}
			return newFloatValHolder(val)
		} else {
			val, err := strconv.ParseInt(v.GetText(), 10, 64)
			if err != nil {
				return valHolder{}
			}
			return newIntValHolder(val)
		}
	case *parser.StringAtomContext:
		val := ""
		l := len(v.GetText())
		if l > 2 {
			val = v.GetText()[1 : l-1]
		}
		return newStringValHolder(val)
	case *parser.RawStringAtomContext:
		val := ""
		l := len(v.GetText())
		if l > 2 {
			val = v.GetText()[1 : l-1]
		}
		return newStringValHolder(val)
	case *parser.ParExprContext:
		return l.handleExpr(v.Expr())
	case *parser.NilAtomContext:
		return newNilValHolder()
	case *parser.BooleanAtomContext:
		val, err := strconv.ParseBool(v.GetText())
		if err != nil {
			return valHolder{}
		}
		return newBooleanValHolder(val)
	case *parser.IdAtomContext:
		//TODO: verify this
		return valHolder{
			dataType: ID,
			value:    v.GetText(),
			name:     v.GetText(),
		}
	case *parser.JsonAtomContext:
		return newStringValHolder(v.GetText())
	default:
		return valHolder{}
	}
}

func (l *KlangListener) handleKubectl_command(ctx parser.IKubectl_commandContext) valHolder {
	switch v := ctx.(type) {
	case *parser.ApplyKubectlCommandContext:
		patterns := []string{"apiVersion", "kind", "metadata.name"}
		namespace := "default"
		if len(v.AllNs()) != 0 {
			namespace = v.Ns(len(v.AllNs()) - 1).GetText()
		}
		updateConfig := ""
		if len(v.AllKubernetes_object_config()) != 0 {
			updateConfigObj := v.Kubernetes_object_config(len(v.AllKubernetes_object_config()) - 1).(*parser.Kubernetes_object_configContext)
			updateConfig = l.GetTextFromStringOrId(updateConfigObj.String_or_id().(*parser.String_or_idContext))
		}
		fNames := v.AllString_or_id()
		k := NewKubectl()
		var resources []Resource
		returnVal := true
		for _, fName := range fNames {
			initialManifest := l.GetTextFromStringOrId(fName.(*parser.String_or_idContext))
			finalManifest := updateMultipleKubernetesObjectsYaml(updateConfig, initialManifest, patterns)
			ar := ApplyRequest{
				Manifest:  finalManifest,
				Namespace: namespace,
				Force:     nil,
				Validate:  nil,
			}
			//TODO: handle response
			responses, err := k.ApplyResource(context.Background(), &ar)
			if err != nil {
				//Continue appyling others
				fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
				continue
				//return newErrHolder(err)
			}
			for _, response := range responses {
				resource := Resource{
					Operation: APPLY,
					Group:     response.GroupVersionKind.Group,
					Version:   response.GroupVersionKind.Version,
					Kind:      response.GroupVersionKind.Kind,
					Namespace: response.Namespace,
					Name:      response.Name,
					Message:   response.Message,
					Status:    ResourceSyncStatusCodeSynced,
				}
				//TODO: update status and health
				if len(response.Err) > 0 {
					resource.Message = response.Err
					resource.Status = ResourceSyncStatusCodeOutOfSync
					returnVal = false
				}
				resources = append(resources, resource)
			}
		}
		for _, resource := range resources {
			resourceKey := kube.NewResourceKey(resource.Group, resource.Kind, resource.Namespace, resource.Name)
			if _, ok := l.kubernetesResources[resourceKey.String()]; !ok {
				l.kubernetesResources[resourceKey.String()] = make([]Resource, 0)
			}
			l.kubernetesResources[resourceKey.String()] = append(l.kubernetesResources[resourceKey.String()], resource)
		}
		return newBooleanValHolder(returnVal)
	case *parser.PatchKubectlCommandContext:
		a := NewFactory(l.mapper)
		k := NewKubectl()
		namespace := "default"
		if len(v.AllNs()) > 0 {
			namespace = v.Ns(len(v.AllNs()) - 1).GetText()
		}
		if len(v.AllString_or_id()) == 0 {
			return newErrHolder(fmt.Errorf("patch cannot be empty (found size %d)", len(v.AllString_or_id())))
		}
		patch := l.GetTextFromStringOrId(v.AllString_or_id()[0].(*parser.String_or_idContext))
		patchType := ""
		if len(v.AllPatch_type()) > 0 {
			patchType = StripQuotes(v.AllPatch_type()[0].GetText())
		}
		resources := v.AllResource()
		var args []string
		for _, resource := range resources {
			args = append(args, l.resolveResource(resource))
		}
		a.ResourceTypeOrNameArgs(args...)
		if len(a.errs) != 0 {
			var errs []string
			for _, err := range a.errs {
				errs = append(errs, err.Error())
			}
			return newErrHolder(fmt.Errorf(strings.Join(errs, "\n")))
		}
		if len(a.ResourceTuples()) != 0 {
			if len(a.ResourceTuples()) != 1 {
				return newErrHolder(fmt.Errorf("can only patch one resource at a time"))
			}
			rts := a.ResourceTuples()
			rt := rts[0]
			resource, err := a.mappingFor(rt.Resource)
			if err != nil {
				return newErrHolder(err)
			}
			if len(rt.Name) == 0 {
				return newErrHolder(fmt.Errorf("name is mandatory for patch operation"))
			}
			pr := PatchRequest{
				Name:             rt.Name,
				Namespace:        namespace,
				GroupVersionKind: resource.GroupVersionKind,
				Patch:            patch,
			}
			if len(patchType) != 0 {
				pr.PatchType = patchType
			}
			resp, err := k.PatchResource(context.Background(), &pr)

			res := Resource{
				Operation: PATCH,
				Group:     pr.GroupVersionKind.Group,
				Version:   pr.GroupVersionKind.Version,
				Kind:      pr.GroupVersionKind.Kind,
				Namespace: pr.Namespace,
				Name:      pr.Name,
				Message:   "",
				Status:    ResourceSyncStatusCodeSynced,
			}
			resourceKey := kube.NewResourceKey(res.Group, res.Kind, res.Namespace, res.Name)
			if _, ok := l.kubernetesResources[resourceKey.String()]; !ok {
				l.kubernetesResources[resourceKey.String()] = make([]Resource, 0)
			}
			if err != nil {
				res.Message = err.Error()
				res.Status = ResourceSyncStatusCodeOutOfSync
				l.kubernetesResources[resourceKey.String()] = append(l.kubernetesResources[resourceKey.String()], res)
				return newErrHolder(err)
			}
			l.kubernetesResources[resourceKey.String()] = append(l.kubernetesResources[resourceKey.String()], res)
			return newStringValHolder(resp.Manifest)
		}
		return newErrHolder(fmt.Errorf("unable to identify resources"))
	case *parser.GetKubectlCommandContext:
		a := NewFactory(l.mapper)
		k := NewKubectl()
		var resp []string
		namespace := "default"
		if len(v.AllNs()) > 0 {
			namespace = v.Ns(len(v.AllNs()) - 1).GetText()
		}
		resources := v.AllResource()
		var args []string
		for _, resource := range resources {
			args = append(args, l.resolveResource(resource))
		}
		a.ResourceTypeOrNameArgs(args...)
		if len(a.errs) != 0 {
			var errs []string
			for _, err := range a.errs {
				errs = append(errs, err.Error())
			}
			return newErrHolder(fmt.Errorf(strings.Join(errs, "\n")))
		}
		if len(a.ResourceTuples()) != 0 {
			rts := a.ResourceTuples()
			for _, rt := range rts {
				resource, err := a.mappingFor(rt.Resource)
				if err != nil {
					return newErrHolder(err)
				}
				gr := GetRequest{
					Namespace:        namespace,
					GroupVersionKind: resource.GroupVersionKind,
				}
				if len(rt.Name) > 0 {
					gr.Name = rt.Name
				}
				mresp, err := k.GetResource(context.Background(), &gr)
				if err != nil {
					//TODO: return error along with retrieved manifests
					return newErrHolder(err)
				}
				resp = append(resp, mresp.Manifest)
			}
		}
		out := strings.Join(resp, ",")
		if len(resp) > 1 {
			out = "{\"apiVersion\": \"v1\",    \"items\": [" + out + "], \"kind\": \"List\", \"metadata\": { \"resourceVersion\": \"\", \"selfLink\": \"\" }}"
		}
		return newStringValHolder(out)
	case *parser.DeleteKubectlCommandContext:
		a := NewFactory(l.mapper)
		k := NewKubectl()
		//var resp []string
		namespace := "default"
		if len(v.AllNs()) > 0 {
			namespace = v.Ns(len(v.AllNs()) - 1).GetText()
		}
		resources := v.AllResource()
		var args []string
		for _, resource := range resources {
			args = append(args, l.resolveResource(resource))
		}
		a.ResourceTypeOrNameArgs(args...)
		if len(a.errs) != 0 {
			var errs []string
			for _, err := range a.errs {
				errs = append(errs, err.Error())
			}
			return newErrHolder(fmt.Errorf(strings.Join(errs, "\n")))
		}
		var err error
		if len(a.ResourceTuples()) != 0 {
			rts := a.ResourceTuples()
			for _, rt := range rts {
				var resource *meta.RESTMapping
				resource, err = a.mappingFor(rt.Resource)
				if err != nil {
					fmt.Printf("{\"err\": \"%s\"}\n", err.Error())
					continue
					//return newErrHolder(err)
				}
				if len(rt.Name) == 0 {
					fmt.Printf("{\"err\": \"%s\"}\n", "resource name cannot be empty with delete statement")
					continue
				}
				dr := DeleteRequest{
					Name:             rt.Name,
					Namespace:        namespace,
					GroupVersionKind: resource.GroupVersionKind,
				}
				_, err = k.DeleteResource(context.Background(), &dr)
				res := Resource{
					Operation: DELETE,
					Group:     dr.GroupVersionKind.Group,
					Version:   dr.GroupVersionKind.Version,
					Kind:      dr.GroupVersionKind.Kind,
					Namespace: dr.Namespace,
					Name:      dr.Name,
					Message:   "",
					Status:    ResourceSyncStatusCodeSynced,
				}
				resourceKey := kube.NewResourceKey(res.Group, res.Kind, res.Namespace, res.Name)
				if _, ok := l.kubernetesResources[resourceKey.String()]; !ok {
					l.kubernetesResources[resourceKey.String()] = make([]Resource, 0)
				}
				//If its an error still try others
				if err != nil {
					res.Message = err.Error()
					res.Status = ResourceSyncStatusCodeOutOfSync
					//l.kubernetesResources[resourceKey.String()] = append(l.kubernetesResources[resourceKey.String()], res)
					//return newErrHolder(err)
				}
				l.kubernetesResources[resourceKey.String()] = append(l.kubernetesResources[resourceKey.String()], res)
				//resp = append(resp, mresp.Manifest)
			}
		}
		if err != nil {
			return newBooleanValHolder(false)
		}
		return newBooleanValHolder(true)
	default:
		break
	}
	return valHolder{}
}

func (l *KlangListener) resolveResource(val parser.IResourceContext) string {
	rc := val.(*parser.ResourceContext)
	if rc.String_or_id() != nil {
		stringOrId := rc.String_or_id().(*parser.String_or_idContext)
		v := l.GetTextFromStringOrId(stringOrId)
		if len(v) != 0 {
			return v
		}
		//If ID doesnt have any value associated then return ID text
		if stringOrId.ID() != nil {
			return stringOrId.GetText()
		}
	}
	if rc.PATH() != nil {
		return rc.PATH().GetText()
	}
	return ""
}

func updateMultipleKubernetesObjectsYaml(from, to string, patterns []string) string {
	tos := strings.Split(to, yamlSeperator)
	froms := strings.Split(from, yamlSeperator)
	var outs []string
	for i, _ := range tos {
		t := tos[i]
		for _, f := range froms {
			t = updateKubernetesObjectsYaml(f, t, patterns)
		}
		outs = append(outs, t)
	}
	return strings.Join(outs, yamlSeperator)
}
