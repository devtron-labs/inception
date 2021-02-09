// Generated from /Users/pghildiy/go/src/github.com/devtron-labs/inception/pkg/language/grammar/Klang.g4 by ANTLR 4.9
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link KlangParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface KlangVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link KlangParser#parse}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParse(KlangParser.ParseContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#block}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBlock(KlangParser.BlockContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#stat}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStat(KlangParser.StatContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#assignment}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAssignment(KlangParser.AssignmentContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#shell_script}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitShell_script(KlangParser.Shell_scriptContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#json_edit_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitJson_edit_fn(KlangParser.Json_edit_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#json_delete_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitJson_delete_fn(KlangParser.Json_delete_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#yaml_edit_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitYaml_edit_fn(KlangParser.Yaml_edit_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#yaml_delete_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitYaml_delete_fn(KlangParser.Yaml_delete_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#kube_json_edit_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKube_json_edit_fn(KlangParser.Kube_json_edit_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#kube_json_delete_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKube_json_delete_fn(KlangParser.Kube_json_delete_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#kube_yaml_edit_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKube_yaml_edit_fn(KlangParser.Kube_yaml_edit_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#kube_yaml_delete_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKube_yaml_delete_fn(KlangParser.Kube_yaml_delete_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#sleep_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSleep_fn(KlangParser.Sleep_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#exit_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExit_fn(KlangParser.Exit_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#if_stat}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIf_stat(KlangParser.If_statContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#condition_block}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCondition_block(KlangParser.Condition_blockContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#stat_block}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStat_block(KlangParser.Stat_blockContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#while_stat}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWhile_stat(KlangParser.While_statContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#log}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLog(KlangParser.LogContext ctx);
	/**
	 * Visit a parse tree produced by the {@code applyKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitApplyKubectlCommand(KlangParser.ApplyKubectlCommandContext ctx);
	/**
	 * Visit a parse tree produced by the {@code patchKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPatchKubectlCommand(KlangParser.PatchKubectlCommandContext ctx);
	/**
	 * Visit a parse tree produced by the {@code getKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGetKubectlCommand(KlangParser.GetKubectlCommandContext ctx);
	/**
	 * Visit a parse tree produced by the {@code deleteKubectlCommand}
	 * labeled alternative in {@link KlangParser#kubectl_command}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDeleteKubectlCommand(KlangParser.DeleteKubectlCommandContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#download_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDownload_fn(KlangParser.Download_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#json_select_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitJson_select_fn(KlangParser.Json_select_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#yaml_select_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitYaml_select_fn(KlangParser.Yaml_select_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#load_fn}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLoad_fn(KlangParser.Load_fnContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#stepInfo}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStepInfo(KlangParser.StepInfoContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#ns}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNs(KlangParser.NsContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#patch_type}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPatch_type(KlangParser.Patch_typeContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#string_or_id}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitString_or_id(KlangParser.String_or_idContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#resource}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitResource(KlangParser.ResourceContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#kubernetes_object_config}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKubernetes_object_config(KlangParser.Kubernetes_object_configContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#filter}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFilter(KlangParser.FilterContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#pattern}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPattern(KlangParser.PatternContext ctx);
	/**
	 * Visit a parse tree produced by the {@code jsonSelectFn}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitJsonSelectFn(KlangParser.JsonSelectFnContext ctx);
	/**
	 * Visit a parse tree produced by the {@code shellScript}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitShellScript(KlangParser.ShellScriptContext ctx);
	/**
	 * Visit a parse tree produced by the {@code yamlSelectFn}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitYamlSelectFn(KlangParser.YamlSelectFnContext ctx);
	/**
	 * Visit a parse tree produced by the {@code atomExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAtomExpr(KlangParser.AtomExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code orExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOrExpr(KlangParser.OrExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code additiveExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAdditiveExpr(KlangParser.AdditiveExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code relationalExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRelationalExpr(KlangParser.RelationalExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code notExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNotExpr(KlangParser.NotExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code unaryMinusExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitUnaryMinusExpr(KlangParser.UnaryMinusExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code multiplicationExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMultiplicationExpr(KlangParser.MultiplicationExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code downloadFn}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDownloadFn(KlangParser.DownloadFnContext ctx);
	/**
	 * Visit a parse tree produced by the {@code powExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPowExpr(KlangParser.PowExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code kubectlExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKubectlExpr(KlangParser.KubectlExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code equalityExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEqualityExpr(KlangParser.EqualityExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code andExpr}
	 * labeled alternative in {@link KlangParser#expr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAndExpr(KlangParser.AndExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code parExpr}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParExpr(KlangParser.ParExprContext ctx);
	/**
	 * Visit a parse tree produced by the {@code numberAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNumberAtom(KlangParser.NumberAtomContext ctx);
	/**
	 * Visit a parse tree produced by the {@code booleanAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBooleanAtom(KlangParser.BooleanAtomContext ctx);
	/**
	 * Visit a parse tree produced by the {@code rawStringAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRawStringAtom(KlangParser.RawStringAtomContext ctx);
	/**
	 * Visit a parse tree produced by the {@code idAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdAtom(KlangParser.IdAtomContext ctx);
	/**
	 * Visit a parse tree produced by the {@code stringAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStringAtom(KlangParser.StringAtomContext ctx);
	/**
	 * Visit a parse tree produced by the {@code jsonAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitJsonAtom(KlangParser.JsonAtomContext ctx);
	/**
	 * Visit a parse tree produced by the {@code nilAtom}
	 * labeled alternative in {@link KlangParser#atom}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNilAtom(KlangParser.NilAtomContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#json}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitJson(KlangParser.JsonContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#obj}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObj(KlangParser.ObjContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#pair}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPair(KlangParser.PairContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#arr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArr(KlangParser.ArrContext ctx);
	/**
	 * Visit a parse tree produced by {@link KlangParser#value}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValue(KlangParser.ValueContext ctx);
}