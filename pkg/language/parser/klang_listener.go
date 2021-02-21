// Generated from Klang.g4 by ANTLR 4.7.

package parser // Klang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// KlangListener is a complete listener for a parse tree produced by KlangParser.
type KlangListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterShell_script is called when entering the shell_script production.
	EnterShell_script(c *Shell_scriptContext)

	// EnterJson_edit_fn is called when entering the json_edit_fn production.
	EnterJson_edit_fn(c *Json_edit_fnContext)

	// EnterJson_delete_fn is called when entering the json_delete_fn production.
	EnterJson_delete_fn(c *Json_delete_fnContext)

	// EnterYaml_edit_fn is called when entering the yaml_edit_fn production.
	EnterYaml_edit_fn(c *Yaml_edit_fnContext)

	// EnterYaml_delete_fn is called when entering the yaml_delete_fn production.
	EnterYaml_delete_fn(c *Yaml_delete_fnContext)

	// EnterKube_json_edit_fn is called when entering the kube_json_edit_fn production.
	EnterKube_json_edit_fn(c *Kube_json_edit_fnContext)

	// EnterKube_json_delete_fn is called when entering the kube_json_delete_fn production.
	EnterKube_json_delete_fn(c *Kube_json_delete_fnContext)

	// EnterKube_yaml_edit_fn is called when entering the kube_yaml_edit_fn production.
	EnterKube_yaml_edit_fn(c *Kube_yaml_edit_fnContext)

	// EnterKube_yaml_delete_fn is called when entering the kube_yaml_delete_fn production.
	EnterKube_yaml_delete_fn(c *Kube_yaml_delete_fnContext)

	// EnterSleep_fn is called when entering the sleep_fn production.
	EnterSleep_fn(c *Sleep_fnContext)

	// EnterExit_fn is called when entering the exit_fn production.
	EnterExit_fn(c *Exit_fnContext)

	// EnterIf_stat is called when entering the if_stat production.
	EnterIf_stat(c *If_statContext)

	// EnterCondition_block is called when entering the condition_block production.
	EnterCondition_block(c *Condition_blockContext)

	// EnterStat_block is called when entering the stat_block production.
	EnterStat_block(c *Stat_blockContext)

	// EnterWhile_stat is called when entering the while_stat production.
	EnterWhile_stat(c *While_statContext)

	// EnterLog is called when entering the log production.
	EnterLog(c *LogContext)

	// EnterApplyKubectlCommand is called when entering the applyKubectlCommand production.
	EnterApplyKubectlCommand(c *ApplyKubectlCommandContext)

	// EnterPatchKubectlCommand is called when entering the patchKubectlCommand production.
	EnterPatchKubectlCommand(c *PatchKubectlCommandContext)

	// EnterGetKubectlCommand is called when entering the getKubectlCommand production.
	EnterGetKubectlCommand(c *GetKubectlCommandContext)

	// EnterDeleteKubectlCommand is called when entering the deleteKubectlCommand production.
	EnterDeleteKubectlCommand(c *DeleteKubectlCommandContext)

	// EnterDownload_fn is called when entering the download_fn production.
	EnterDownload_fn(c *Download_fnContext)

	// EnterJson_select_fn is called when entering the json_select_fn production.
	EnterJson_select_fn(c *Json_select_fnContext)

	// EnterYaml_select_fn is called when entering the yaml_select_fn production.
	EnterYaml_select_fn(c *Yaml_select_fnContext)

	// EnterLoad_fn is called when entering the load_fn production.
	EnterLoad_fn(c *Load_fnContext)

	// EnterStepInfo is called when entering the stepInfo production.
	EnterStepInfo(c *StepInfoContext)

	// EnterNs is called when entering the ns production.
	EnterNs(c *NsContext)

	// EnterAsObject is called when entering the asObject production.
	EnterAsObject(c *AsObjectContext)

	// EnterPatch_type is called when entering the patch_type production.
	EnterPatch_type(c *Patch_typeContext)

	// EnterString_or_id is called when entering the string_or_id production.
	EnterString_or_id(c *String_or_idContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterKubernetes_object_config is called when entering the kubernetes_object_config production.
	EnterKubernetes_object_config(c *Kubernetes_object_configContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterJsonSelectFn is called when entering the jsonSelectFn production.
	EnterJsonSelectFn(c *JsonSelectFnContext)

	// EnterShellScript is called when entering the shellScript production.
	EnterShellScript(c *ShellScriptContext)

	// EnterYamlSelectFn is called when entering the yamlSelectFn production.
	EnterYamlSelectFn(c *YamlSelectFnContext)

	// EnterAtomExpr is called when entering the atomExpr production.
	EnterAtomExpr(c *AtomExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterRelationalExpr is called when entering the relationalExpr production.
	EnterRelationalExpr(c *RelationalExprContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterUnaryMinusExpr is called when entering the unaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// EnterMultiplicationExpr is called when entering the multiplicationExpr production.
	EnterMultiplicationExpr(c *MultiplicationExprContext)

	// EnterDownloadFn is called when entering the downloadFn production.
	EnterDownloadFn(c *DownloadFnContext)

	// EnterPowExpr is called when entering the powExpr production.
	EnterPowExpr(c *PowExprContext)

	// EnterKubectlExpr is called when entering the kubectlExpr production.
	EnterKubectlExpr(c *KubectlExprContext)

	// EnterEqualityExpr is called when entering the equalityExpr production.
	EnterEqualityExpr(c *EqualityExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterParExpr is called when entering the parExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterNumberAtom is called when entering the numberAtom production.
	EnterNumberAtom(c *NumberAtomContext)

	// EnterBooleanAtom is called when entering the booleanAtom production.
	EnterBooleanAtom(c *BooleanAtomContext)

	// EnterRawStringAtom is called when entering the rawStringAtom production.
	EnterRawStringAtom(c *RawStringAtomContext)

	// EnterIdAtom is called when entering the idAtom production.
	EnterIdAtom(c *IdAtomContext)

	// EnterStringAtom is called when entering the stringAtom production.
	EnterStringAtom(c *StringAtomContext)

	// EnterJsonAtom is called when entering the jsonAtom production.
	EnterJsonAtom(c *JsonAtomContext)

	// EnterNilAtom is called when entering the nilAtom production.
	EnterNilAtom(c *NilAtomContext)

	// EnterJson is called when entering the json production.
	EnterJson(c *JsonContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArr is called when entering the arr production.
	EnterArr(c *ArrContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitShell_script is called when exiting the shell_script production.
	ExitShell_script(c *Shell_scriptContext)

	// ExitJson_edit_fn is called when exiting the json_edit_fn production.
	ExitJson_edit_fn(c *Json_edit_fnContext)

	// ExitJson_delete_fn is called when exiting the json_delete_fn production.
	ExitJson_delete_fn(c *Json_delete_fnContext)

	// ExitYaml_edit_fn is called when exiting the yaml_edit_fn production.
	ExitYaml_edit_fn(c *Yaml_edit_fnContext)

	// ExitYaml_delete_fn is called when exiting the yaml_delete_fn production.
	ExitYaml_delete_fn(c *Yaml_delete_fnContext)

	// ExitKube_json_edit_fn is called when exiting the kube_json_edit_fn production.
	ExitKube_json_edit_fn(c *Kube_json_edit_fnContext)

	// ExitKube_json_delete_fn is called when exiting the kube_json_delete_fn production.
	ExitKube_json_delete_fn(c *Kube_json_delete_fnContext)

	// ExitKube_yaml_edit_fn is called when exiting the kube_yaml_edit_fn production.
	ExitKube_yaml_edit_fn(c *Kube_yaml_edit_fnContext)

	// ExitKube_yaml_delete_fn is called when exiting the kube_yaml_delete_fn production.
	ExitKube_yaml_delete_fn(c *Kube_yaml_delete_fnContext)

	// ExitSleep_fn is called when exiting the sleep_fn production.
	ExitSleep_fn(c *Sleep_fnContext)

	// ExitExit_fn is called when exiting the exit_fn production.
	ExitExit_fn(c *Exit_fnContext)

	// ExitIf_stat is called when exiting the if_stat production.
	ExitIf_stat(c *If_statContext)

	// ExitCondition_block is called when exiting the condition_block production.
	ExitCondition_block(c *Condition_blockContext)

	// ExitStat_block is called when exiting the stat_block production.
	ExitStat_block(c *Stat_blockContext)

	// ExitWhile_stat is called when exiting the while_stat production.
	ExitWhile_stat(c *While_statContext)

	// ExitLog is called when exiting the log production.
	ExitLog(c *LogContext)

	// ExitApplyKubectlCommand is called when exiting the applyKubectlCommand production.
	ExitApplyKubectlCommand(c *ApplyKubectlCommandContext)

	// ExitPatchKubectlCommand is called when exiting the patchKubectlCommand production.
	ExitPatchKubectlCommand(c *PatchKubectlCommandContext)

	// ExitGetKubectlCommand is called when exiting the getKubectlCommand production.
	ExitGetKubectlCommand(c *GetKubectlCommandContext)

	// ExitDeleteKubectlCommand is called when exiting the deleteKubectlCommand production.
	ExitDeleteKubectlCommand(c *DeleteKubectlCommandContext)

	// ExitDownload_fn is called when exiting the download_fn production.
	ExitDownload_fn(c *Download_fnContext)

	// ExitJson_select_fn is called when exiting the json_select_fn production.
	ExitJson_select_fn(c *Json_select_fnContext)

	// ExitYaml_select_fn is called when exiting the yaml_select_fn production.
	ExitYaml_select_fn(c *Yaml_select_fnContext)

	// ExitLoad_fn is called when exiting the load_fn production.
	ExitLoad_fn(c *Load_fnContext)

	// ExitStepInfo is called when exiting the stepInfo production.
	ExitStepInfo(c *StepInfoContext)

	// ExitNs is called when exiting the ns production.
	ExitNs(c *NsContext)

	// ExitAsObject is called when exiting the asObject production.
	ExitAsObject(c *AsObjectContext)

	// ExitPatch_type is called when exiting the patch_type production.
	ExitPatch_type(c *Patch_typeContext)

	// ExitString_or_id is called when exiting the string_or_id production.
	ExitString_or_id(c *String_or_idContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitKubernetes_object_config is called when exiting the kubernetes_object_config production.
	ExitKubernetes_object_config(c *Kubernetes_object_configContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitJsonSelectFn is called when exiting the jsonSelectFn production.
	ExitJsonSelectFn(c *JsonSelectFnContext)

	// ExitShellScript is called when exiting the shellScript production.
	ExitShellScript(c *ShellScriptContext)

	// ExitYamlSelectFn is called when exiting the yamlSelectFn production.
	ExitYamlSelectFn(c *YamlSelectFnContext)

	// ExitAtomExpr is called when exiting the atomExpr production.
	ExitAtomExpr(c *AtomExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitRelationalExpr is called when exiting the relationalExpr production.
	ExitRelationalExpr(c *RelationalExprContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitUnaryMinusExpr is called when exiting the unaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitMultiplicationExpr is called when exiting the multiplicationExpr production.
	ExitMultiplicationExpr(c *MultiplicationExprContext)

	// ExitDownloadFn is called when exiting the downloadFn production.
	ExitDownloadFn(c *DownloadFnContext)

	// ExitPowExpr is called when exiting the powExpr production.
	ExitPowExpr(c *PowExprContext)

	// ExitKubectlExpr is called when exiting the kubectlExpr production.
	ExitKubectlExpr(c *KubectlExprContext)

	// ExitEqualityExpr is called when exiting the equalityExpr production.
	ExitEqualityExpr(c *EqualityExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitParExpr is called when exiting the parExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitNumberAtom is called when exiting the numberAtom production.
	ExitNumberAtom(c *NumberAtomContext)

	// ExitBooleanAtom is called when exiting the booleanAtom production.
	ExitBooleanAtom(c *BooleanAtomContext)

	// ExitRawStringAtom is called when exiting the rawStringAtom production.
	ExitRawStringAtom(c *RawStringAtomContext)

	// ExitIdAtom is called when exiting the idAtom production.
	ExitIdAtom(c *IdAtomContext)

	// ExitStringAtom is called when exiting the stringAtom production.
	ExitStringAtom(c *StringAtomContext)

	// ExitJsonAtom is called when exiting the jsonAtom production.
	ExitJsonAtom(c *JsonAtomContext)

	// ExitNilAtom is called when exiting the nilAtom production.
	ExitNilAtom(c *NilAtomContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArr is called when exiting the arr production.
	ExitArr(c *ArrContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
