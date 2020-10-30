// Generated from Klang.g4 by ANTLR 4.7.

package parser // Klang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseKlangListener is a complete listener for a parse tree produced by KlangParser.
type BaseKlangListener struct{}

var _ KlangListener = &BaseKlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseKlangListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseKlangListener) ExitParse(ctx *ParseContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseKlangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseKlangListener) ExitBlock(ctx *BlockContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseKlangListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseKlangListener) ExitStat(ctx *StatContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseKlangListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseKlangListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterShell_script is called when production shell_script is entered.
func (s *BaseKlangListener) EnterShell_script(ctx *Shell_scriptContext) {}

// ExitShell_script is called when production shell_script is exited.
func (s *BaseKlangListener) ExitShell_script(ctx *Shell_scriptContext) {}

// EnterJson_edit_fn is called when production json_edit_fn is entered.
func (s *BaseKlangListener) EnterJson_edit_fn(ctx *Json_edit_fnContext) {}

// ExitJson_edit_fn is called when production json_edit_fn is exited.
func (s *BaseKlangListener) ExitJson_edit_fn(ctx *Json_edit_fnContext) {}

// EnterJson_delete_fn is called when production json_delete_fn is entered.
func (s *BaseKlangListener) EnterJson_delete_fn(ctx *Json_delete_fnContext) {}

// ExitJson_delete_fn is called when production json_delete_fn is exited.
func (s *BaseKlangListener) ExitJson_delete_fn(ctx *Json_delete_fnContext) {}

// EnterYaml_edit_fn is called when production yaml_edit_fn is entered.
func (s *BaseKlangListener) EnterYaml_edit_fn(ctx *Yaml_edit_fnContext) {}

// ExitYaml_edit_fn is called when production yaml_edit_fn is exited.
func (s *BaseKlangListener) ExitYaml_edit_fn(ctx *Yaml_edit_fnContext) {}

// EnterYaml_delete_fn is called when production yaml_delete_fn is entered.
func (s *BaseKlangListener) EnterYaml_delete_fn(ctx *Yaml_delete_fnContext) {}

// ExitYaml_delete_fn is called when production yaml_delete_fn is exited.
func (s *BaseKlangListener) ExitYaml_delete_fn(ctx *Yaml_delete_fnContext) {}

// EnterSleep_fn is called when production sleep_fn is entered.
func (s *BaseKlangListener) EnterSleep_fn(ctx *Sleep_fnContext) {}

// ExitSleep_fn is called when production sleep_fn is exited.
func (s *BaseKlangListener) ExitSleep_fn(ctx *Sleep_fnContext) {}

// EnterIf_stat is called when production if_stat is entered.
func (s *BaseKlangListener) EnterIf_stat(ctx *If_statContext) {}

// ExitIf_stat is called when production if_stat is exited.
func (s *BaseKlangListener) ExitIf_stat(ctx *If_statContext) {}

// EnterCondition_block is called when production condition_block is entered.
func (s *BaseKlangListener) EnterCondition_block(ctx *Condition_blockContext) {}

// ExitCondition_block is called when production condition_block is exited.
func (s *BaseKlangListener) ExitCondition_block(ctx *Condition_blockContext) {}

// EnterStat_block is called when production stat_block is entered.
func (s *BaseKlangListener) EnterStat_block(ctx *Stat_blockContext) {}

// ExitStat_block is called when production stat_block is exited.
func (s *BaseKlangListener) ExitStat_block(ctx *Stat_blockContext) {}

// EnterWhile_stat is called when production while_stat is entered.
func (s *BaseKlangListener) EnterWhile_stat(ctx *While_statContext) {}

// ExitWhile_stat is called when production while_stat is exited.
func (s *BaseKlangListener) ExitWhile_stat(ctx *While_statContext) {}

// EnterLog is called when production log is entered.
func (s *BaseKlangListener) EnterLog(ctx *LogContext) {}

// ExitLog is called when production log is exited.
func (s *BaseKlangListener) ExitLog(ctx *LogContext) {}

// EnterApplyKubectlCommand is called when production applyKubectlCommand is entered.
func (s *BaseKlangListener) EnterApplyKubectlCommand(ctx *ApplyKubectlCommandContext) {}

// ExitApplyKubectlCommand is called when production applyKubectlCommand is exited.
func (s *BaseKlangListener) ExitApplyKubectlCommand(ctx *ApplyKubectlCommandContext) {}

// EnterPatchKubectlCommand is called when production patchKubectlCommand is entered.
func (s *BaseKlangListener) EnterPatchKubectlCommand(ctx *PatchKubectlCommandContext) {}

// ExitPatchKubectlCommand is called when production patchKubectlCommand is exited.
func (s *BaseKlangListener) ExitPatchKubectlCommand(ctx *PatchKubectlCommandContext) {}

// EnterGetKubectlCommand is called when production getKubectlCommand is entered.
func (s *BaseKlangListener) EnterGetKubectlCommand(ctx *GetKubectlCommandContext) {}

// ExitGetKubectlCommand is called when production getKubectlCommand is exited.
func (s *BaseKlangListener) ExitGetKubectlCommand(ctx *GetKubectlCommandContext) {}

// EnterDeleteKubectlCommand is called when production deleteKubectlCommand is entered.
func (s *BaseKlangListener) EnterDeleteKubectlCommand(ctx *DeleteKubectlCommandContext) {}

// ExitDeleteKubectlCommand is called when production deleteKubectlCommand is exited.
func (s *BaseKlangListener) ExitDeleteKubectlCommand(ctx *DeleteKubectlCommandContext) {}

// EnterDownload_fn is called when production download_fn is entered.
func (s *BaseKlangListener) EnterDownload_fn(ctx *Download_fnContext) {}

// ExitDownload_fn is called when production download_fn is exited.
func (s *BaseKlangListener) ExitDownload_fn(ctx *Download_fnContext) {}

// EnterJson_select_fn is called when production json_select_fn is entered.
func (s *BaseKlangListener) EnterJson_select_fn(ctx *Json_select_fnContext) {}

// ExitJson_select_fn is called when production json_select_fn is exited.
func (s *BaseKlangListener) ExitJson_select_fn(ctx *Json_select_fnContext) {}

// EnterYaml_select_fn is called when production yaml_select_fn is entered.
func (s *BaseKlangListener) EnterYaml_select_fn(ctx *Yaml_select_fnContext) {}

// ExitYaml_select_fn is called when production yaml_select_fn is exited.
func (s *BaseKlangListener) ExitYaml_select_fn(ctx *Yaml_select_fnContext) {}

// EnterLoad_fn is called when production load_fn is entered.
func (s *BaseKlangListener) EnterLoad_fn(ctx *Load_fnContext) {}

// ExitLoad_fn is called when production load_fn is exited.
func (s *BaseKlangListener) ExitLoad_fn(ctx *Load_fnContext) {}

// EnterNs is called when production ns is entered.
func (s *BaseKlangListener) EnterNs(ctx *NsContext) {}

// ExitNs is called when production ns is exited.
func (s *BaseKlangListener) ExitNs(ctx *NsContext) {}

// EnterPatch_type is called when production patch_type is entered.
func (s *BaseKlangListener) EnterPatch_type(ctx *Patch_typeContext) {}

// ExitPatch_type is called when production patch_type is exited.
func (s *BaseKlangListener) ExitPatch_type(ctx *Patch_typeContext) {}

// EnterString_or_id is called when production string_or_id is entered.
func (s *BaseKlangListener) EnterString_or_id(ctx *String_or_idContext) {}

// ExitString_or_id is called when production string_or_id is exited.
func (s *BaseKlangListener) ExitString_or_id(ctx *String_or_idContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseKlangListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseKlangListener) ExitResource(ctx *ResourceContext) {}

// EnterKubernetes_object_config is called when production kubernetes_object_config is entered.
func (s *BaseKlangListener) EnterKubernetes_object_config(ctx *Kubernetes_object_configContext) {}

// ExitKubernetes_object_config is called when production kubernetes_object_config is exited.
func (s *BaseKlangListener) ExitKubernetes_object_config(ctx *Kubernetes_object_configContext) {}

// EnterJsonSelectFn is called when production jsonSelectFn is entered.
func (s *BaseKlangListener) EnterJsonSelectFn(ctx *JsonSelectFnContext) {}

// ExitJsonSelectFn is called when production jsonSelectFn is exited.
func (s *BaseKlangListener) ExitJsonSelectFn(ctx *JsonSelectFnContext) {}

// EnterShellScript is called when production shellScript is entered.
func (s *BaseKlangListener) EnterShellScript(ctx *ShellScriptContext) {}

// ExitShellScript is called when production shellScript is exited.
func (s *BaseKlangListener) ExitShellScript(ctx *ShellScriptContext) {}

// EnterYamlSelectFn is called when production yamlSelectFn is entered.
func (s *BaseKlangListener) EnterYamlSelectFn(ctx *YamlSelectFnContext) {}

// ExitYamlSelectFn is called when production yamlSelectFn is exited.
func (s *BaseKlangListener) ExitYamlSelectFn(ctx *YamlSelectFnContext) {}

// EnterAtomExpr is called when production atomExpr is entered.
func (s *BaseKlangListener) EnterAtomExpr(ctx *AtomExprContext) {}

// ExitAtomExpr is called when production atomExpr is exited.
func (s *BaseKlangListener) ExitAtomExpr(ctx *AtomExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseKlangListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseKlangListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BaseKlangListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BaseKlangListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterRelationalExpr is called when production relationalExpr is entered.
func (s *BaseKlangListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production relationalExpr is exited.
func (s *BaseKlangListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BaseKlangListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BaseKlangListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterUnaryMinusExpr is called when production unaryMinusExpr is entered.
func (s *BaseKlangListener) EnterUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// ExitUnaryMinusExpr is called when production unaryMinusExpr is exited.
func (s *BaseKlangListener) ExitUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// EnterMultiplicationExpr is called when production multiplicationExpr is entered.
func (s *BaseKlangListener) EnterMultiplicationExpr(ctx *MultiplicationExprContext) {}

// ExitMultiplicationExpr is called when production multiplicationExpr is exited.
func (s *BaseKlangListener) ExitMultiplicationExpr(ctx *MultiplicationExprContext) {}

// EnterDownloadFn is called when production downloadFn is entered.
func (s *BaseKlangListener) EnterDownloadFn(ctx *DownloadFnContext) {}

// ExitDownloadFn is called when production downloadFn is exited.
func (s *BaseKlangListener) ExitDownloadFn(ctx *DownloadFnContext) {}

// EnterPowExpr is called when production powExpr is entered.
func (s *BaseKlangListener) EnterPowExpr(ctx *PowExprContext) {}

// ExitPowExpr is called when production powExpr is exited.
func (s *BaseKlangListener) ExitPowExpr(ctx *PowExprContext) {}

// EnterKubectlExpr is called when production kubectlExpr is entered.
func (s *BaseKlangListener) EnterKubectlExpr(ctx *KubectlExprContext) {}

// ExitKubectlExpr is called when production kubectlExpr is exited.
func (s *BaseKlangListener) ExitKubectlExpr(ctx *KubectlExprContext) {}

// EnterEqualityExpr is called when production equalityExpr is entered.
func (s *BaseKlangListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production equalityExpr is exited.
func (s *BaseKlangListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseKlangListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseKlangListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterParExpr is called when production parExpr is entered.
func (s *BaseKlangListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production parExpr is exited.
func (s *BaseKlangListener) ExitParExpr(ctx *ParExprContext) {}

// EnterNumberAtom is called when production numberAtom is entered.
func (s *BaseKlangListener) EnterNumberAtom(ctx *NumberAtomContext) {}

// ExitNumberAtom is called when production numberAtom is exited.
func (s *BaseKlangListener) ExitNumberAtom(ctx *NumberAtomContext) {}

// EnterBooleanAtom is called when production booleanAtom is entered.
func (s *BaseKlangListener) EnterBooleanAtom(ctx *BooleanAtomContext) {}

// ExitBooleanAtom is called when production booleanAtom is exited.
func (s *BaseKlangListener) ExitBooleanAtom(ctx *BooleanAtomContext) {}

// EnterRawStringAtom is called when production rawStringAtom is entered.
func (s *BaseKlangListener) EnterRawStringAtom(ctx *RawStringAtomContext) {}

// ExitRawStringAtom is called when production rawStringAtom is exited.
func (s *BaseKlangListener) ExitRawStringAtom(ctx *RawStringAtomContext) {}

// EnterIdAtom is called when production idAtom is entered.
func (s *BaseKlangListener) EnterIdAtom(ctx *IdAtomContext) {}

// ExitIdAtom is called when production idAtom is exited.
func (s *BaseKlangListener) ExitIdAtom(ctx *IdAtomContext) {}

// EnterStringAtom is called when production stringAtom is entered.
func (s *BaseKlangListener) EnterStringAtom(ctx *StringAtomContext) {}

// ExitStringAtom is called when production stringAtom is exited.
func (s *BaseKlangListener) ExitStringAtom(ctx *StringAtomContext) {}

// EnterJsonAtom is called when production jsonAtom is entered.
func (s *BaseKlangListener) EnterJsonAtom(ctx *JsonAtomContext) {}

// ExitJsonAtom is called when production jsonAtom is exited.
func (s *BaseKlangListener) ExitJsonAtom(ctx *JsonAtomContext) {}

// EnterNilAtom is called when production nilAtom is entered.
func (s *BaseKlangListener) EnterNilAtom(ctx *NilAtomContext) {}

// ExitNilAtom is called when production nilAtom is exited.
func (s *BaseKlangListener) ExitNilAtom(ctx *NilAtomContext) {}

// EnterJson is called when production json is entered.
func (s *BaseKlangListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseKlangListener) ExitJson(ctx *JsonContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseKlangListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseKlangListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseKlangListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseKlangListener) ExitPair(ctx *PairContext) {}

// EnterArr is called when production arr is entered.
func (s *BaseKlangListener) EnterArr(ctx *ArrContext) {}

// ExitArr is called when production arr is exited.
func (s *BaseKlangListener) ExitArr(ctx *ArrContext) {}

// EnterValue is called when production value is entered.
func (s *BaseKlangListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseKlangListener) ExitValue(ctx *ValueContext) {}
