// Generated from /Users/pghildiy/go/src/github.com/devtron-labs/inception/pkg/language/grammar/Klang.g4 by ANTLR 4.9
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link KlangParser}.
 */
public interface KlangListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link KlangParser#parse}.
	 * @param ctx the parse tree
	 */
	void enterParse(KlangParser.ParseContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#parse}.
	 * @param ctx the parse tree
	 */
	void exitParse(KlangParser.ParseContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(KlangParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(KlangParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#stat}.
	 * @param ctx the parse tree
	 */
	void enterStat(KlangParser.StatContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#stat}.
	 * @param ctx the parse tree
	 */
	void exitStat(KlangParser.StatContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#assignment}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(KlangParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#assignment}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(KlangParser.AssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#shell_script}.
	 * @param ctx the parse tree
	 */
	void enterShell_script(KlangParser.Shell_scriptContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#shell_script}.
	 * @param ctx the parse tree
	 */
	void exitShell_script(KlangParser.Shell_scriptContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#json_edit_fn}.
	 * @param ctx the parse tree
	 */
	void enterJson_edit_fn(KlangParser.Json_edit_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#json_edit_fn}.
	 * @param ctx the parse tree
	 */
	void exitJson_edit_fn(KlangParser.Json_edit_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#json_delete_fn}.
	 * @param ctx the parse tree
	 */
	void enterJson_delete_fn(KlangParser.Json_delete_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#json_delete_fn}.
	 * @param ctx the parse tree
	 */
	void exitJson_delete_fn(KlangParser.Json_delete_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#yaml_edit_fn}.
	 * @param ctx the parse tree
	 */
	void enterYaml_edit_fn(KlangParser.Yaml_edit_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#yaml_edit_fn}.
	 * @param ctx the parse tree
	 */
	void exitYaml_edit_fn(KlangParser.Yaml_edit_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#yaml_delete_fn}.
	 * @param ctx the parse tree
	 */
	void enterYaml_delete_fn(KlangParser.Yaml_delete_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#yaml_delete_fn}.
	 * @param ctx the parse tree
	 */
	void exitYaml_delete_fn(KlangParser.Yaml_delete_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#kube_json_edit_fn}.
	 * @param ctx the parse tree
	 */
	void enterKube_json_edit_fn(KlangParser.Kube_json_edit_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#kube_json_edit_fn}.
	 * @param ctx the parse tree
	 */
	void exitKube_json_edit_fn(KlangParser.Kube_json_edit_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#kube_json_delete_fn}.
	 * @param ctx the parse tree
	 */
	void enterKube_json_delete_fn(KlangParser.Kube_json_delete_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#kube_json_delete_fn}.
	 * @param ctx the parse tree
	 */
	void exitKube_json_delete_fn(KlangParser.Kube_json_delete_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#kube_yaml_edit_fn}.
	 * @param ctx the parse tree
	 */
	void enterKube_yaml_edit_fn(KlangParser.Kube_yaml_edit_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#kube_yaml_edit_fn}.
	 * @param ctx the parse tree
	 */
	void exitKube_yaml_edit_fn(KlangParser.Kube_yaml_edit_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#kube_yaml_delete_fn}.
	 * @param ctx the parse tree
	 */
	void enterKube_yaml_delete_fn(KlangParser.Kube_yaml_delete_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#kube_yaml_delete_fn}.
	 * @param ctx the parse tree
	 */
	void exitKube_yaml_delete_fn(KlangParser.Kube_yaml_delete_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#sleep_fn}.
	 * @param ctx the parse tree
	 */
	void enterSleep_fn(KlangParser.Sleep_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#sleep_fn}.
	 * @param ctx the parse tree
	 */
	void exitSleep_fn(KlangParser.Sleep_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#exit_fn}.
	 * @param ctx the parse tree
	 */
	void enterExit_fn(KlangParser.Exit_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#exit_fn}.
	 * @param ctx the parse tree
	 */
	void exitExit_fn(KlangParser.Exit_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#if_stat}.
	 * @param ctx the parse tree
	 */
	void enterIf_stat(KlangParser.If_statContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#if_stat}.
	 * @param ctx the parse tree
	 */
	void exitIf_stat(KlangParser.If_statContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#condition_block}.
	 * @param ctx the parse tree
	 */
	void enterCondition_block(KlangParser.Condition_blockContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#condition_block}.
	 * @param ctx the parse tree
	 */
	void exitCondition_block(KlangParser.Condition_blockContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#stat_block}.
	 * @param ctx the parse tree
	 */
	void enterStat_block(KlangParser.Stat_blockContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#stat_block}.
	 * @param ctx the parse tree
	 */
	void exitStat_block(KlangParser.Stat_blockContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#while_stat}.
	 * @param ctx the parse tree
	 */
	void enterWhile_stat(KlangParser.While_statContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#while_stat}.
	 * @param ctx the parse tree
	 */
	void exitWhile_stat(KlangParser.While_statContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#log}.
	 * @param ctx the parse tree
	 */
	void enterLog(KlangParser.LogContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#log}.
	 * @param ctx the parse tree
	 */
	void exitLog(KlangParser.LogContext ctx);
	/**
	 * Enter a parse tree produced by the {@code applyKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 */
	void enterApplyKubectlCommand(KlangParser.ApplyKubectlCommandContext ctx);
	/**
	 * Exit a parse tree produced by the {@code applyKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 */
	void exitApplyKubectlCommand(KlangParser.ApplyKubectlCommandContext ctx);
	/**
	 * Enter a parse tree produced by the {@code patchKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 */
	void enterPatchKubectlCommand(KlangParser.PatchKubectlCommandContext ctx);
	/**
	 * Exit a parse tree produced by the {@code patchKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 */
	void exitPatchKubectlCommand(KlangParser.PatchKubectlCommandContext ctx);
	/**
	 * Enter a parse tree produced by the {@code getKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 */
	void enterGetKubectlCommand(KlangParser.GetKubectlCommandContext ctx);
	/**
	 * Exit a parse tree produced by the {@code getKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 */
	void exitGetKubectlCommand(KlangParser.GetKubectlCommandContext ctx);
	/**
	 * Enter a parse tree produced by the {@code deleteKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 */
	void enterDeleteKubectlCommand(KlangParser.DeleteKubectlCommandContext ctx);
	/**
	 * Exit a parse tree produced by the {@code deleteKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 */
	void exitDeleteKubectlCommand(KlangParser.DeleteKubectlCommandContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#download_fn}.
	 * @param ctx the parse tree
	 */
	void enterDownload_fn(KlangParser.Download_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#download_fn}.
	 * @param ctx the parse tree
	 */
	void exitDownload_fn(KlangParser.Download_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#json_select_fn}.
	 * @param ctx the parse tree
	 */
	void enterJson_select_fn(KlangParser.Json_select_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#json_select_fn}.
	 * @param ctx the parse tree
	 */
	void exitJson_select_fn(KlangParser.Json_select_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#yaml_select_fn}.
	 * @param ctx the parse tree
	 */
	void enterYaml_select_fn(KlangParser.Yaml_select_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#yaml_select_fn}.
	 * @param ctx the parse tree
	 */
	void exitYaml_select_fn(KlangParser.Yaml_select_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#load_fn}.
	 * @param ctx the parse tree
	 */
	void enterLoad_fn(KlangParser.Load_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#load_fn}.
	 * @param ctx the parse tree
	 */
	void exitLoad_fn(KlangParser.Load_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#stepInfo}.
	 * @param ctx the parse tree
	 */
	void enterStepInfo(KlangParser.StepInfoContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#stepInfo}.
	 * @param ctx the parse tree
	 */
	void exitStepInfo(KlangParser.StepInfoContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#ns}.
	 * @param ctx the parse tree
	 */
	void enterNs(KlangParser.NsContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#ns}.
	 * @param ctx the parse tree
	 */
	void exitNs(KlangParser.NsContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#patch_type}.
	 * @param ctx the parse tree
	 */
	void enterPatch_type(KlangParser.Patch_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#patch_type}.
	 * @param ctx the parse tree
	 */
	void exitPatch_type(KlangParser.Patch_typeContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#string_or_id}.
	 * @param ctx the parse tree
	 */
	void enterString_or_id(KlangParser.String_or_idContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#string_or_id}.
	 * @param ctx the parse tree
	 */
	void exitString_or_id(KlangParser.String_or_idContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#resource}.
	 * @param ctx the parse tree
	 */
	void enterResource(KlangParser.ResourceContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#resource}.
	 * @param ctx the parse tree
	 */
	void exitResource(KlangParser.ResourceContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#kubernetes_object_config}.
	 * @param ctx the parse tree
	 */
	void enterKubernetes_object_config(KlangParser.Kubernetes_object_configContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#kubernetes_object_config}.
	 * @param ctx the parse tree
	 */
	void exitKubernetes_object_config(KlangParser.Kubernetes_object_configContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#filter}.
	 * @param ctx the parse tree
	 */
	void enterFilter(KlangParser.FilterContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#filter}.
	 * @param ctx the parse tree
	 */
	void exitFilter(KlangParser.FilterContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#pattern}.
	 * @param ctx the parse tree
	 */
	void enterPattern(KlangParser.PatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#pattern}.
	 * @param ctx the parse tree
	 */
	void exitPattern(KlangParser.PatternContext ctx);
	/**
	 * Enter a parse tree produced by the {@code jsonSelectFn}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterJsonSelectFn(KlangParser.JsonSelectFnContext ctx);
	/**
	 * Exit a parse tree produced by the {@code jsonSelectFn}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitJsonSelectFn(KlangParser.JsonSelectFnContext ctx);
	/**
	 * Enter a parse tree produced by the {@code shellScript}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterShellScript(KlangParser.ShellScriptContext ctx);
	/**
	 * Exit a parse tree produced by the {@code shellScript}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitShellScript(KlangParser.ShellScriptContext ctx);
	/**
	 * Enter a parse tree produced by the {@code yamlSelectFn}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterYamlSelectFn(KlangParser.YamlSelectFnContext ctx);
	/**
	 * Exit a parse tree produced by the {@code yamlSelectFn}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitYamlSelectFn(KlangParser.YamlSelectFnContext ctx);
	/**
	 * Enter a parse tree produced by the {@code atomExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAtomExpr(KlangParser.AtomExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code atomExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAtomExpr(KlangParser.AtomExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code orExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterOrExpr(KlangParser.OrExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code orExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitOrExpr(KlangParser.OrExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code additiveExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAdditiveExpr(KlangParser.AdditiveExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code additiveExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAdditiveExpr(KlangParser.AdditiveExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code relationalExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterRelationalExpr(KlangParser.RelationalExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code relationalExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitRelationalExpr(KlangParser.RelationalExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code notExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterNotExpr(KlangParser.NotExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code notExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitNotExpr(KlangParser.NotExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code unaryMinusExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterUnaryMinusExpr(KlangParser.UnaryMinusExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code unaryMinusExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitUnaryMinusExpr(KlangParser.UnaryMinusExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code multiplicationExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterMultiplicationExpr(KlangParser.MultiplicationExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code multiplicationExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitMultiplicationExpr(KlangParser.MultiplicationExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code downloadFn}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterDownloadFn(KlangParser.DownloadFnContext ctx);
	/**
	 * Exit a parse tree produced by the {@code downloadFn}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitDownloadFn(KlangParser.DownloadFnContext ctx);
	/**
	 * Enter a parse tree produced by the {@code powExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterPowExpr(KlangParser.PowExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code powExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitPowExpr(KlangParser.PowExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code kubectlExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterKubectlExpr(KlangParser.KubectlExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code kubectlExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitKubectlExpr(KlangParser.KubectlExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code equalityExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterEqualityExpr(KlangParser.EqualityExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code equalityExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitEqualityExpr(KlangParser.EqualityExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code andExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAndExpr(KlangParser.AndExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code andExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAndExpr(KlangParser.AndExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code parExpr}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterParExpr(KlangParser.ParExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code parExpr}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitParExpr(KlangParser.ParExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code numberAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterNumberAtom(KlangParser.NumberAtomContext ctx);
	/**
	 * Exit a parse tree produced by the {@code numberAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitNumberAtom(KlangParser.NumberAtomContext ctx);
	/**
	 * Enter a parse tree produced by the {@code booleanAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterBooleanAtom(KlangParser.BooleanAtomContext ctx);
	/**
	 * Exit a parse tree produced by the {@code booleanAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitBooleanAtom(KlangParser.BooleanAtomContext ctx);
	/**
	 * Enter a parse tree produced by the {@code rawStringAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterRawStringAtom(KlangParser.RawStringAtomContext ctx);
	/**
	 * Exit a parse tree produced by the {@code rawStringAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitRawStringAtom(KlangParser.RawStringAtomContext ctx);
	/**
	 * Enter a parse tree produced by the {@code idAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterIdAtom(KlangParser.IdAtomContext ctx);
	/**
	 * Exit a parse tree produced by the {@code idAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitIdAtom(KlangParser.IdAtomContext ctx);
	/**
	 * Enter a parse tree produced by the {@code stringAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterStringAtom(KlangParser.StringAtomContext ctx);
	/**
	 * Exit a parse tree produced by the {@code stringAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitStringAtom(KlangParser.StringAtomContext ctx);
	/**
	 * Enter a parse tree produced by the {@code jsonAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterJsonAtom(KlangParser.JsonAtomContext ctx);
	/**
	 * Exit a parse tree produced by the {@code jsonAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitJsonAtom(KlangParser.JsonAtomContext ctx);
	/**
	 * Enter a parse tree produced by the {@code nilAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterNilAtom(KlangParser.NilAtomContext ctx);
	/**
	 * Exit a parse tree produced by the {@code nilAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitNilAtom(KlangParser.NilAtomContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#json}.
	 * @param ctx the parse tree
	 */
	void enterJson(KlangParser.JsonContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#json}.
	 * @param ctx the parse tree
	 */
	void exitJson(KlangParser.JsonContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#obj}.
	 * @param ctx the parse tree
	 */
	void enterObj(KlangParser.ObjContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#obj}.
	 * @param ctx the parse tree
	 */
	void exitObj(KlangParser.ObjContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#pair}.
	 * @param ctx the parse tree
	 */
	void enterPair(KlangParser.PairContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#pair}.
	 * @param ctx the parse tree
	 */
	void exitPair(KlangParser.PairContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#arr}.
	 * @param ctx the parse tree
	 */
	void enterArr(KlangParser.ArrContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#arr}.
	 * @param ctx the parse tree
	 */
	void exitArr(KlangParser.ArrContext ctx);
	/**
	 * Enter a parse tree produced by {@link KlangParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(KlangParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link KlangParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(KlangParser.ValueContext ctx);
}