// Generated from /Users/pghildiy/go/src/github.com/devtron-labs/inception/pkg/language/grammar/Klang.g4 by ANTLR 4.9
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class KlangParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, OR=5, AND=6, EQ=7, NEQ=8, GT=9, LT=10, 
		GTEQ=11, LTEQ=12, PLUS=13, MINUS=14, MULT=15, DIV=16, MOD=17, POW=18, 
		NOT=19, SCOL=20, ASSIGN=21, OPAR=22, CPAR=23, OBRACE=24, CBRACE=25, COMMA=26, 
		TRUE=27, FALSE=28, NIL=29, IF=30, ELSE=31, WHILE=32, LOG=33, KUBECTL=34, 
		APPLY=35, PATCH=36, GET=37, REPLACE=38, DELETE=39, NAMESPACE=40, PATCHTYPE=41, 
		PATCHLOAD=42, UPDATELOAD=43, JSONPATH=44, LOAD=45, EXIT=46, JSONSELECT=47, 
		JSONEDIT=48, JSONDELETE=49, YAMLSELECT=50, YAMLEDIT=51, YAMLDELETE=52, 
		KUBEJSONEDIT=53, KUBEJSONDELETE=54, KUBEYAMLEDIT=55, KUBEYAMLDELETE=56, 
		SHELLSCRIPT=57, DOWNLOAD=58, SLEEP=59, STEPINFO=60, FILTER=61, PATTERN=62, 
		ID=63, NUMBER=64, PATH=65, RAW_STRING_LIT=66, STRING=67, COMMENT=68, SPACE=69, 
		OTHER=70;
	public static final int
		RULE_parse = 0, RULE_block = 1, RULE_stat = 2, RULE_assignment = 3, RULE_shell_script = 4, 
		RULE_json_edit_fn = 5, RULE_json_delete_fn = 6, RULE_yaml_edit_fn = 7, 
		RULE_yaml_delete_fn = 8, RULE_kube_json_edit_fn = 9, RULE_kube_json_delete_fn = 10, 
		RULE_kube_yaml_edit_fn = 11, RULE_kube_yaml_delete_fn = 12, RULE_sleep_fn = 13, 
		RULE_exit_fn = 14, RULE_if_stat = 15, RULE_condition_block = 16, RULE_stat_block = 17, 
		RULE_while_stat = 18, RULE_log = 19, RULE_kubectl_command = 20, RULE_download_fn = 21, 
		RULE_json_select_fn = 22, RULE_yaml_select_fn = 23, RULE_load_fn = 24, 
		RULE_stepInfo = 25, RULE_ns = 26, RULE_patch_type = 27, RULE_string_or_id = 28, 
		RULE_resource = 29, RULE_kubernetes_object_config = 30, RULE_filter = 31, 
		RULE_pattern = 32, RULE_expr = 33, RULE_atom = 34, RULE_json = 35, RULE_obj = 36, 
		RULE_pair = 37, RULE_arr = 38, RULE_value = 39;
	private static String[] makeRuleNames() {
		return new String[] {
			"parse", "block", "stat", "assignment", "shell_script", "json_edit_fn", 
			"json_delete_fn", "yaml_edit_fn", "yaml_delete_fn", "kube_json_edit_fn", 
			"kube_json_delete_fn", "kube_yaml_edit_fn", "kube_yaml_delete_fn", "sleep_fn", 
			"exit_fn", "if_stat", "condition_block", "stat_block", "while_stat", 
			"log", "kubectl_command", "download_fn", "json_select_fn", "yaml_select_fn", 
			"load_fn", "stepInfo", "ns", "patch_type", "string_or_id", "resource", 
			"kubernetes_object_config", "filter", "pattern", "expr", "atom", "json", 
			"obj", "pair", "arr", "value"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':'", "'['", "']'", "'null'", "'||'", "'&&'", "'=='", "'!='", 
			"'>'", "'<'", "'>='", "'<='", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", 
			"'!'", "';'", "'='", "'('", "')'", "'{'", "'}'", "','", "'true'", "'false'", 
			"'nil'", "'if'", "'else'", "'while'", "'log'", "'kubectl'", "'apply'", 
			"'patch'", "'get'", "'replace'", "'delete'", "'-n'", "'--type'", "'-p'", 
			"'-u'", "'-jsonpath'", "'load'", "'exit'", "'jsonSelect'", "'jsonEdit'", 
			"'jsonDelete'", "'yamlSelect'", "'yamlEdit'", "'yamlDelete'", "'kubeJsonEdit'", 
			"'kubeJsonDelete'", "'kubeYamlEdit'", "'kubeYamlDelete'", "'shellScript'", 
			"'download'", "'sleep'", "'stepInfo'", "'filter'", "'pattern'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", 
			"LTEQ", "PLUS", "MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "SCOL", 
			"ASSIGN", "OPAR", "CPAR", "OBRACE", "CBRACE", "COMMA", "TRUE", "FALSE", 
			"NIL", "IF", "ELSE", "WHILE", "LOG", "KUBECTL", "APPLY", "PATCH", "GET", 
			"REPLACE", "DELETE", "NAMESPACE", "PATCHTYPE", "PATCHLOAD", "UPDATELOAD", 
			"JSONPATH", "LOAD", "EXIT", "JSONSELECT", "JSONEDIT", "JSONDELETE", "YAMLSELECT", 
			"YAMLEDIT", "YAMLDELETE", "KUBEJSONEDIT", "KUBEJSONDELETE", "KUBEYAMLEDIT", 
			"KUBEYAMLDELETE", "SHELLSCRIPT", "DOWNLOAD", "SLEEP", "STEPINFO", "FILTER", 
			"PATTERN", "ID", "NUMBER", "PATH", "RAW_STRING_LIT", "STRING", "COMMENT", 
			"SPACE", "OTHER"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Klang.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public KlangParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ParseContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(KlangParser.EOF, 0); }
		public ParseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parse; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterParse(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitParse(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitParse(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ParseContext parse() throws RecognitionException {
		ParseContext _localctx = new ParseContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_parse);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			block();
			setState(81);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public List<StatContext> stat() {
			return getRuleContexts(StatContext.class);
		}
		public StatContext stat(int i) {
			return getRuleContext(StatContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitBlock(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitBlock(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 30)) & ~0x3f) == 0 && ((1L << (_la - 30)) & ((1L << (IF - 30)) | (1L << (WHILE - 30)) | (1L << (LOG - 30)) | (1L << (EXIT - 30)) | (1L << (JSONEDIT - 30)) | (1L << (JSONDELETE - 30)) | (1L << (YAMLEDIT - 30)) | (1L << (YAMLDELETE - 30)) | (1L << (KUBEJSONEDIT - 30)) | (1L << (KUBEJSONDELETE - 30)) | (1L << (KUBEYAMLEDIT - 30)) | (1L << (KUBEYAMLDELETE - 30)) | (1L << (SLEEP - 30)) | (1L << (ID - 30)) | (1L << (OTHER - 30)))) != 0)) {
				{
				{
				setState(83);
				stat();
				}
				}
				setState(88);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatContext extends ParserRuleContext {
		public Token OTHER;
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public Json_edit_fnContext json_edit_fn() {
			return getRuleContext(Json_edit_fnContext.class,0);
		}
		public Json_delete_fnContext json_delete_fn() {
			return getRuleContext(Json_delete_fnContext.class,0);
		}
		public Yaml_edit_fnContext yaml_edit_fn() {
			return getRuleContext(Yaml_edit_fnContext.class,0);
		}
		public Yaml_delete_fnContext yaml_delete_fn() {
			return getRuleContext(Yaml_delete_fnContext.class,0);
		}
		public Kube_json_delete_fnContext kube_json_delete_fn() {
			return getRuleContext(Kube_json_delete_fnContext.class,0);
		}
		public Kube_json_edit_fnContext kube_json_edit_fn() {
			return getRuleContext(Kube_json_edit_fnContext.class,0);
		}
		public Kube_yaml_delete_fnContext kube_yaml_delete_fn() {
			return getRuleContext(Kube_yaml_delete_fnContext.class,0);
		}
		public Kube_yaml_edit_fnContext kube_yaml_edit_fn() {
			return getRuleContext(Kube_yaml_edit_fnContext.class,0);
		}
		public If_statContext if_stat() {
			return getRuleContext(If_statContext.class,0);
		}
		public While_statContext while_stat() {
			return getRuleContext(While_statContext.class,0);
		}
		public Sleep_fnContext sleep_fn() {
			return getRuleContext(Sleep_fnContext.class,0);
		}
		public Exit_fnContext exit_fn() {
			return getRuleContext(Exit_fnContext.class,0);
		}
		public LogContext log() {
			return getRuleContext(LogContext.class,0);
		}
		public TerminalNode OTHER() { return getToken(KlangParser.OTHER, 0); }
		public StatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stat; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterStat(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitStat(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitStat(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StatContext stat() throws RecognitionException {
		StatContext _localctx = new StatContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_stat);
		try {
			setState(105);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				setState(89);
				assignment();
				}
				break;
			case JSONEDIT:
				enterOuterAlt(_localctx, 2);
				{
				setState(90);
				json_edit_fn();
				}
				break;
			case JSONDELETE:
				enterOuterAlt(_localctx, 3);
				{
				setState(91);
				json_delete_fn();
				}
				break;
			case YAMLEDIT:
				enterOuterAlt(_localctx, 4);
				{
				setState(92);
				yaml_edit_fn();
				}
				break;
			case YAMLDELETE:
				enterOuterAlt(_localctx, 5);
				{
				setState(93);
				yaml_delete_fn();
				}
				break;
			case KUBEJSONDELETE:
				enterOuterAlt(_localctx, 6);
				{
				setState(94);
				kube_json_delete_fn();
				}
				break;
			case KUBEJSONEDIT:
				enterOuterAlt(_localctx, 7);
				{
				setState(95);
				kube_json_edit_fn();
				}
				break;
			case KUBEYAMLDELETE:
				enterOuterAlt(_localctx, 8);
				{
				setState(96);
				kube_yaml_delete_fn();
				}
				break;
			case KUBEYAMLEDIT:
				enterOuterAlt(_localctx, 9);
				{
				setState(97);
				kube_yaml_edit_fn();
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 10);
				{
				setState(98);
				if_stat();
				}
				break;
			case WHILE:
				enterOuterAlt(_localctx, 11);
				{
				setState(99);
				while_stat();
				}
				break;
			case SLEEP:
				enterOuterAlt(_localctx, 12);
				{
				setState(100);
				sleep_fn();
				}
				break;
			case EXIT:
				enterOuterAlt(_localctx, 13);
				{
				setState(101);
				exit_fn();
				}
				break;
			case LOG:
				enterOuterAlt(_localctx, 14);
				{
				setState(102);
				log();
				}
				break;
			case OTHER:
				enterOuterAlt(_localctx, 15);
				{
				setState(103);
				((StatContext)_localctx).OTHER = match(OTHER);
				fmt.Println("unknown char: " + (((StatContext)_localctx).OTHER!=null?((StatContext)_localctx).OTHER.getText():null));
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(KlangParser.ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public Load_fnContext load_fn() {
			return getRuleContext(Load_fnContext.class,0);
		}
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterAssignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitAssignment(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitAssignment(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_assignment);
		try {
			setState(117);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(107);
				match(ID);
				setState(108);
				match(ASSIGN);
				setState(109);
				expr(0);
				setState(110);
				match(SCOL);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				match(ID);
				setState(113);
				match(ASSIGN);
				setState(114);
				load_fn();
				setState(115);
				match(SCOL);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Shell_scriptContext extends ParserRuleContext {
		public TerminalNode SHELLSCRIPT() { return getToken(KlangParser.SHELLSCRIPT, 0); }
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public Shell_scriptContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_shell_script; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterShell_script(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitShell_script(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitShell_script(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Shell_scriptContext shell_script() throws RecognitionException {
		Shell_scriptContext _localctx = new Shell_scriptContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_shell_script);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			match(SHELLSCRIPT);
			setState(120);
			string_or_id();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Json_edit_fnContext extends ParserRuleContext {
		public TerminalNode JSONEDIT() { return getToken(KlangParser.JSONEDIT, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public List<TerminalNode> COMMA() { return getTokens(KlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(KlangParser.COMMA, i);
		}
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public Json_edit_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_json_edit_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterJson_edit_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitJson_edit_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitJson_edit_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Json_edit_fnContext json_edit_fn() throws RecognitionException {
		Json_edit_fnContext _localctx = new Json_edit_fnContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_json_edit_fn);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(122);
			match(JSONEDIT);
			setState(123);
			match(OPAR);
			setState(124);
			match(ID);
			setState(125);
			match(COMMA);
			setState(126);
			string_or_id();
			setState(127);
			match(COMMA);
			setState(128);
			expr(0);
			setState(129);
			match(CPAR);
			setState(130);
			match(SCOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Json_delete_fnContext extends ParserRuleContext {
		public TerminalNode JSONDELETE() { return getToken(KlangParser.JSONDELETE, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public TerminalNode COMMA() { return getToken(KlangParser.COMMA, 0); }
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public Json_delete_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_json_delete_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterJson_delete_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitJson_delete_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitJson_delete_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Json_delete_fnContext json_delete_fn() throws RecognitionException {
		Json_delete_fnContext _localctx = new Json_delete_fnContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_json_delete_fn);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			match(JSONDELETE);
			setState(133);
			match(OPAR);
			setState(134);
			match(ID);
			setState(135);
			match(COMMA);
			setState(136);
			string_or_id();
			setState(137);
			match(CPAR);
			setState(138);
			match(SCOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Yaml_edit_fnContext extends ParserRuleContext {
		public TerminalNode YAMLEDIT() { return getToken(KlangParser.YAMLEDIT, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public List<TerminalNode> COMMA() { return getTokens(KlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(KlangParser.COMMA, i);
		}
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public TerminalNode NUMBER() { return getToken(KlangParser.NUMBER, 0); }
		public Yaml_edit_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_yaml_edit_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterYaml_edit_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitYaml_edit_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitYaml_edit_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Yaml_edit_fnContext yaml_edit_fn() throws RecognitionException {
		Yaml_edit_fnContext _localctx = new Yaml_edit_fnContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_yaml_edit_fn);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			match(YAMLEDIT);
			setState(141);
			match(OPAR);
			setState(142);
			match(ID);
			setState(143);
			match(COMMA);
			setState(144);
			string_or_id();
			setState(145);
			match(COMMA);
			setState(146);
			expr(0);
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(147);
				match(COMMA);
				setState(148);
				match(NUMBER);
				}
			}

			setState(151);
			match(CPAR);
			setState(152);
			match(SCOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Yaml_delete_fnContext extends ParserRuleContext {
		public TerminalNode YAMLDELETE() { return getToken(KlangParser.YAMLDELETE, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public List<TerminalNode> COMMA() { return getTokens(KlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(KlangParser.COMMA, i);
		}
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public TerminalNode NUMBER() { return getToken(KlangParser.NUMBER, 0); }
		public Yaml_delete_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_yaml_delete_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterYaml_delete_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitYaml_delete_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitYaml_delete_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Yaml_delete_fnContext yaml_delete_fn() throws RecognitionException {
		Yaml_delete_fnContext _localctx = new Yaml_delete_fnContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_yaml_delete_fn);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			match(YAMLDELETE);
			setState(155);
			match(OPAR);
			setState(156);
			match(ID);
			setState(157);
			match(COMMA);
			setState(158);
			string_or_id();
			setState(161);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(159);
				match(COMMA);
				setState(160);
				match(NUMBER);
				}
			}

			setState(163);
			match(CPAR);
			setState(164);
			match(SCOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Kube_json_edit_fnContext extends ParserRuleContext {
		public TerminalNode KUBEJSONEDIT() { return getToken(KlangParser.KUBEJSONEDIT, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public List<TerminalNode> COMMA() { return getTokens(KlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(KlangParser.COMMA, i);
		}
		public List<String_or_idContext> string_or_id() {
			return getRuleContexts(String_or_idContext.class);
		}
		public String_or_idContext string_or_id(int i) {
			return getRuleContext(String_or_idContext.class,i);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public Kube_json_edit_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_kube_json_edit_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterKube_json_edit_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitKube_json_edit_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitKube_json_edit_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Kube_json_edit_fnContext kube_json_edit_fn() throws RecognitionException {
		Kube_json_edit_fnContext _localctx = new Kube_json_edit_fnContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_kube_json_edit_fn);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(166);
			match(KUBEJSONEDIT);
			setState(167);
			match(OPAR);
			setState(168);
			match(ID);
			setState(169);
			match(COMMA);
			setState(170);
			string_or_id();
			setState(171);
			match(COMMA);
			setState(172);
			expr(0);
			setState(175);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(173);
				match(COMMA);
				setState(174);
				string_or_id();
				}
			}

			setState(177);
			match(CPAR);
			setState(178);
			match(SCOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Kube_json_delete_fnContext extends ParserRuleContext {
		public TerminalNode KUBEJSONDELETE() { return getToken(KlangParser.KUBEJSONDELETE, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public List<TerminalNode> COMMA() { return getTokens(KlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(KlangParser.COMMA, i);
		}
		public FilterContext filter() {
			return getRuleContext(FilterContext.class,0);
		}
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public Kube_json_delete_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_kube_json_delete_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterKube_json_delete_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitKube_json_delete_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitKube_json_delete_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Kube_json_delete_fnContext kube_json_delete_fn() throws RecognitionException {
		Kube_json_delete_fnContext _localctx = new Kube_json_delete_fnContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_kube_json_delete_fn);
		try {
			setState(216);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(180);
				match(KUBEJSONDELETE);
				setState(181);
				match(OPAR);
				setState(182);
				match(ID);
				setState(183);
				match(COMMA);
				setState(184);
				filter();
				setState(185);
				match(COMMA);
				setState(186);
				pattern();
				setState(187);
				match(CPAR);
				setState(188);
				match(SCOL);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(190);
				match(KUBEJSONDELETE);
				setState(191);
				match(OPAR);
				setState(192);
				match(ID);
				setState(193);
				match(COMMA);
				setState(194);
				pattern();
				setState(195);
				match(COMMA);
				setState(196);
				filter();
				setState(197);
				match(CPAR);
				setState(198);
				match(SCOL);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(200);
				match(KUBEJSONDELETE);
				setState(201);
				match(OPAR);
				setState(202);
				match(ID);
				setState(203);
				match(COMMA);
				setState(204);
				filter();
				setState(205);
				match(CPAR);
				setState(206);
				match(SCOL);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(208);
				match(KUBEJSONDELETE);
				setState(209);
				match(OPAR);
				setState(210);
				match(ID);
				setState(211);
				match(COMMA);
				setState(212);
				pattern();
				setState(213);
				match(CPAR);
				setState(214);
				match(SCOL);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Kube_yaml_edit_fnContext extends ParserRuleContext {
		public TerminalNode KUBEYAMLEDIT() { return getToken(KlangParser.KUBEYAMLEDIT, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public List<TerminalNode> COMMA() { return getTokens(KlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(KlangParser.COMMA, i);
		}
		public List<String_or_idContext> string_or_id() {
			return getRuleContexts(String_or_idContext.class);
		}
		public String_or_idContext string_or_id(int i) {
			return getRuleContext(String_or_idContext.class,i);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public Kube_yaml_edit_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_kube_yaml_edit_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterKube_yaml_edit_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitKube_yaml_edit_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitKube_yaml_edit_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Kube_yaml_edit_fnContext kube_yaml_edit_fn() throws RecognitionException {
		Kube_yaml_edit_fnContext _localctx = new Kube_yaml_edit_fnContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_kube_yaml_edit_fn);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(218);
			match(KUBEYAMLEDIT);
			setState(219);
			match(OPAR);
			setState(220);
			match(ID);
			setState(221);
			match(COMMA);
			setState(222);
			string_or_id();
			setState(223);
			match(COMMA);
			setState(224);
			expr(0);
			setState(227);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(225);
				match(COMMA);
				setState(226);
				string_or_id();
				}
			}

			setState(229);
			match(CPAR);
			setState(230);
			match(SCOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Kube_yaml_delete_fnContext extends ParserRuleContext {
		public TerminalNode KUBEYAMLDELETE() { return getToken(KlangParser.KUBEYAMLDELETE, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public List<TerminalNode> COMMA() { return getTokens(KlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(KlangParser.COMMA, i);
		}
		public FilterContext filter() {
			return getRuleContext(FilterContext.class,0);
		}
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public Kube_yaml_delete_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_kube_yaml_delete_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterKube_yaml_delete_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitKube_yaml_delete_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitKube_yaml_delete_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Kube_yaml_delete_fnContext kube_yaml_delete_fn() throws RecognitionException {
		Kube_yaml_delete_fnContext _localctx = new Kube_yaml_delete_fnContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_kube_yaml_delete_fn);
		try {
			setState(268);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(232);
				match(KUBEYAMLDELETE);
				setState(233);
				match(OPAR);
				setState(234);
				match(ID);
				setState(235);
				match(COMMA);
				setState(236);
				filter();
				setState(237);
				match(COMMA);
				setState(238);
				pattern();
				setState(239);
				match(CPAR);
				setState(240);
				match(SCOL);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(242);
				match(KUBEYAMLDELETE);
				setState(243);
				match(OPAR);
				setState(244);
				match(ID);
				setState(245);
				match(COMMA);
				setState(246);
				pattern();
				setState(247);
				match(COMMA);
				setState(248);
				filter();
				setState(249);
				match(CPAR);
				setState(250);
				match(SCOL);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(252);
				match(KUBEYAMLDELETE);
				setState(253);
				match(OPAR);
				setState(254);
				match(ID);
				setState(255);
				match(COMMA);
				setState(256);
				filter();
				setState(257);
				match(CPAR);
				setState(258);
				match(SCOL);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(260);
				match(KUBEYAMLDELETE);
				setState(261);
				match(OPAR);
				setState(262);
				match(ID);
				setState(263);
				match(COMMA);
				setState(264);
				pattern();
				setState(265);
				match(CPAR);
				setState(266);
				match(SCOL);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Sleep_fnContext extends ParserRuleContext {
		public TerminalNode SLEEP() { return getToken(KlangParser.SLEEP, 0); }
		public TerminalNode NUMBER() { return getToken(KlangParser.NUMBER, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public Sleep_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sleep_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterSleep_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitSleep_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitSleep_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Sleep_fnContext sleep_fn() throws RecognitionException {
		Sleep_fnContext _localctx = new Sleep_fnContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_sleep_fn);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(270);
			match(SLEEP);
			setState(271);
			match(NUMBER);
			setState(272);
			match(SCOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Exit_fnContext extends ParserRuleContext {
		public TerminalNode EXIT() { return getToken(KlangParser.EXIT, 0); }
		public TerminalNode NUMBER() { return getToken(KlangParser.NUMBER, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public Exit_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exit_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterExit_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitExit_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitExit_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Exit_fnContext exit_fn() throws RecognitionException {
		Exit_fnContext _localctx = new Exit_fnContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_exit_fn);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(274);
			match(EXIT);
			setState(275);
			match(NUMBER);
			setState(276);
			match(SCOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class If_statContext extends ParserRuleContext {
		public List<TerminalNode> IF() { return getTokens(KlangParser.IF); }
		public TerminalNode IF(int i) {
			return getToken(KlangParser.IF, i);
		}
		public List<Condition_blockContext> condition_block() {
			return getRuleContexts(Condition_blockContext.class);
		}
		public Condition_blockContext condition_block(int i) {
			return getRuleContext(Condition_blockContext.class,i);
		}
		public List<TerminalNode> ELSE() { return getTokens(KlangParser.ELSE); }
		public TerminalNode ELSE(int i) {
			return getToken(KlangParser.ELSE, i);
		}
		public Stat_blockContext stat_block() {
			return getRuleContext(Stat_blockContext.class,0);
		}
		public If_statContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_stat; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterIf_stat(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitIf_stat(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitIf_stat(this);
			else return visitor.visitChildren(this);
		}
	}

	public final If_statContext if_stat() throws RecognitionException {
		If_statContext _localctx = new If_statContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_if_stat);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(278);
			match(IF);
			setState(279);
			condition_block();
			setState(285);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(280);
					match(ELSE);
					setState(281);
					match(IF);
					setState(282);
					condition_block();
					}
					} 
				}
				setState(287);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			setState(290);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				setState(288);
				match(ELSE);
				setState(289);
				stat_block();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Condition_blockContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public Stat_blockContext stat_block() {
			return getRuleContext(Stat_blockContext.class,0);
		}
		public Condition_blockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condition_block; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterCondition_block(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitCondition_block(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitCondition_block(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Condition_blockContext condition_block() throws RecognitionException {
		Condition_blockContext _localctx = new Condition_blockContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_condition_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(292);
			expr(0);
			setState(293);
			stat_block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Stat_blockContext extends ParserRuleContext {
		public TerminalNode OBRACE() { return getToken(KlangParser.OBRACE, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode CBRACE() { return getToken(KlangParser.CBRACE, 0); }
		public StatContext stat() {
			return getRuleContext(StatContext.class,0);
		}
		public Stat_blockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stat_block; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterStat_block(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitStat_block(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitStat_block(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Stat_blockContext stat_block() throws RecognitionException {
		Stat_blockContext _localctx = new Stat_blockContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_stat_block);
		try {
			setState(300);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OBRACE:
				enterOuterAlt(_localctx, 1);
				{
				setState(295);
				match(OBRACE);
				setState(296);
				block();
				setState(297);
				match(CBRACE);
				}
				break;
			case IF:
			case WHILE:
			case LOG:
			case EXIT:
			case JSONEDIT:
			case JSONDELETE:
			case YAMLEDIT:
			case YAMLDELETE:
			case KUBEJSONEDIT:
			case KUBEJSONDELETE:
			case KUBEYAMLEDIT:
			case KUBEYAMLDELETE:
			case SLEEP:
			case ID:
			case OTHER:
				enterOuterAlt(_localctx, 2);
				{
				setState(299);
				stat();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class While_statContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(KlangParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public Stat_blockContext stat_block() {
			return getRuleContext(Stat_blockContext.class,0);
		}
		public While_statContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_stat; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterWhile_stat(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitWhile_stat(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitWhile_stat(this);
			else return visitor.visitChildren(this);
		}
	}

	public final While_statContext while_stat() throws RecognitionException {
		While_statContext _localctx = new While_statContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_while_stat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(302);
			match(WHILE);
			setState(303);
			expr(0);
			setState(304);
			stat_block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LogContext extends ParserRuleContext {
		public TerminalNode LOG() { return getToken(KlangParser.LOG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public LogContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_log; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterLog(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitLog(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitLog(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LogContext log() throws RecognitionException {
		LogContext _localctx = new LogContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_log);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(306);
			match(LOG);
			setState(307);
			expr(0);
			setState(308);
			match(SCOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Kubectl_commandContext extends ParserRuleContext {
		public Kubectl_commandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_kubectl_command; }
	 
		public Kubectl_commandContext() { }
		public void copyFrom(Kubectl_commandContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DeleteKubectlCommandContext extends Kubectl_commandContext {
		public TerminalNode KUBECTL() { return getToken(KlangParser.KUBECTL, 0); }
		public TerminalNode DELETE() { return getToken(KlangParser.DELETE, 0); }
		public List<TerminalNode> NAMESPACE() { return getTokens(KlangParser.NAMESPACE); }
		public TerminalNode NAMESPACE(int i) {
			return getToken(KlangParser.NAMESPACE, i);
		}
		public List<NsContext> ns() {
			return getRuleContexts(NsContext.class);
		}
		public NsContext ns(int i) {
			return getRuleContext(NsContext.class,i);
		}
		public List<ResourceContext> resource() {
			return getRuleContexts(ResourceContext.class);
		}
		public ResourceContext resource(int i) {
			return getRuleContext(ResourceContext.class,i);
		}
		public DeleteKubectlCommandContext(Kubectl_commandContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterDeleteKubectlCommand(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitDeleteKubectlCommand(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitDeleteKubectlCommand(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class PatchKubectlCommandContext extends Kubectl_commandContext {
		public TerminalNode KUBECTL() { return getToken(KlangParser.KUBECTL, 0); }
		public TerminalNode PATCH() { return getToken(KlangParser.PATCH, 0); }
		public List<TerminalNode> NAMESPACE() { return getTokens(KlangParser.NAMESPACE); }
		public TerminalNode NAMESPACE(int i) {
			return getToken(KlangParser.NAMESPACE, i);
		}
		public List<NsContext> ns() {
			return getRuleContexts(NsContext.class);
		}
		public NsContext ns(int i) {
			return getRuleContext(NsContext.class,i);
		}
		public List<ResourceContext> resource() {
			return getRuleContexts(ResourceContext.class);
		}
		public ResourceContext resource(int i) {
			return getRuleContext(ResourceContext.class,i);
		}
		public List<TerminalNode> PATCHTYPE() { return getTokens(KlangParser.PATCHTYPE); }
		public TerminalNode PATCHTYPE(int i) {
			return getToken(KlangParser.PATCHTYPE, i);
		}
		public List<Patch_typeContext> patch_type() {
			return getRuleContexts(Patch_typeContext.class);
		}
		public Patch_typeContext patch_type(int i) {
			return getRuleContext(Patch_typeContext.class,i);
		}
		public List<TerminalNode> PATCHLOAD() { return getTokens(KlangParser.PATCHLOAD); }
		public TerminalNode PATCHLOAD(int i) {
			return getToken(KlangParser.PATCHLOAD, i);
		}
		public List<String_or_idContext> string_or_id() {
			return getRuleContexts(String_or_idContext.class);
		}
		public String_or_idContext string_or_id(int i) {
			return getRuleContext(String_or_idContext.class,i);
		}
		public PatchKubectlCommandContext(Kubectl_commandContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterPatchKubectlCommand(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitPatchKubectlCommand(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitPatchKubectlCommand(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class ApplyKubectlCommandContext extends Kubectl_commandContext {
		public TerminalNode KUBECTL() { return getToken(KlangParser.KUBECTL, 0); }
		public TerminalNode APPLY() { return getToken(KlangParser.APPLY, 0); }
		public List<TerminalNode> NAMESPACE() { return getTokens(KlangParser.NAMESPACE); }
		public TerminalNode NAMESPACE(int i) {
			return getToken(KlangParser.NAMESPACE, i);
		}
		public List<NsContext> ns() {
			return getRuleContexts(NsContext.class);
		}
		public NsContext ns(int i) {
			return getRuleContext(NsContext.class,i);
		}
		public List<String_or_idContext> string_or_id() {
			return getRuleContexts(String_or_idContext.class);
		}
		public String_or_idContext string_or_id(int i) {
			return getRuleContext(String_or_idContext.class,i);
		}
		public List<TerminalNode> UPDATELOAD() { return getTokens(KlangParser.UPDATELOAD); }
		public TerminalNode UPDATELOAD(int i) {
			return getToken(KlangParser.UPDATELOAD, i);
		}
		public List<Kubernetes_object_configContext> kubernetes_object_config() {
			return getRuleContexts(Kubernetes_object_configContext.class);
		}
		public Kubernetes_object_configContext kubernetes_object_config(int i) {
			return getRuleContext(Kubernetes_object_configContext.class,i);
		}
		public ApplyKubectlCommandContext(Kubectl_commandContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterApplyKubectlCommand(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitApplyKubectlCommand(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitApplyKubectlCommand(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class GetKubectlCommandContext extends Kubectl_commandContext {
		public TerminalNode KUBECTL() { return getToken(KlangParser.KUBECTL, 0); }
		public TerminalNode GET() { return getToken(KlangParser.GET, 0); }
		public List<TerminalNode> NAMESPACE() { return getTokens(KlangParser.NAMESPACE); }
		public TerminalNode NAMESPACE(int i) {
			return getToken(KlangParser.NAMESPACE, i);
		}
		public List<NsContext> ns() {
			return getRuleContexts(NsContext.class);
		}
		public NsContext ns(int i) {
			return getRuleContext(NsContext.class,i);
		}
		public List<ResourceContext> resource() {
			return getRuleContexts(ResourceContext.class);
		}
		public ResourceContext resource(int i) {
			return getRuleContext(ResourceContext.class,i);
		}
		public GetKubectlCommandContext(Kubectl_commandContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterGetKubectlCommand(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitGetKubectlCommand(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitGetKubectlCommand(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Kubectl_commandContext kubectl_command() throws RecognitionException {
		Kubectl_commandContext _localctx = new Kubectl_commandContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_kubectl_command);
		try {
			int _alt;
			setState(352);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				_localctx = new ApplyKubectlCommandContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(310);
				match(KUBECTL);
				setState(311);
				match(APPLY);
				setState(317); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						setState(317);
						_errHandler.sync(this);
						switch (_input.LA(1)) {
						case NAMESPACE:
							{
							setState(312);
							match(NAMESPACE);
							setState(313);
							ns();
							}
							break;
						case ID:
						case RAW_STRING_LIT:
						case STRING:
							{
							setState(314);
							string_or_id();
							}
							break;
						case UPDATELOAD:
							{
							setState(315);
							match(UPDATELOAD);
							setState(316);
							kubernetes_object_config();
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(319); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 2:
				_localctx = new PatchKubectlCommandContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(321);
				match(KUBECTL);
				setState(322);
				match(PATCH);
				setState(330); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						setState(330);
						_errHandler.sync(this);
						switch (_input.LA(1)) {
						case NAMESPACE:
							{
							setState(323);
							match(NAMESPACE);
							setState(324);
							ns();
							}
							break;
						case ID:
						case PATH:
						case RAW_STRING_LIT:
						case STRING:
							{
							setState(325);
							resource();
							}
							break;
						case PATCHTYPE:
							{
							setState(326);
							match(PATCHTYPE);
							setState(327);
							patch_type();
							}
							break;
						case PATCHLOAD:
							{
							setState(328);
							match(PATCHLOAD);
							setState(329);
							string_or_id();
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(332); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 3:
				_localctx = new GetKubectlCommandContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(334);
				match(KUBECTL);
				setState(335);
				match(GET);
				setState(339); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						setState(339);
						_errHandler.sync(this);
						switch (_input.LA(1)) {
						case NAMESPACE:
							{
							setState(336);
							match(NAMESPACE);
							setState(337);
							ns();
							}
							break;
						case ID:
						case PATH:
						case RAW_STRING_LIT:
						case STRING:
							{
							setState(338);
							resource();
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(341); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 4:
				_localctx = new DeleteKubectlCommandContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(343);
				match(KUBECTL);
				setState(344);
				match(DELETE);
				setState(348); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						setState(348);
						_errHandler.sync(this);
						switch (_input.LA(1)) {
						case NAMESPACE:
							{
							setState(345);
							match(NAMESPACE);
							setState(346);
							ns();
							}
							break;
						case ID:
						case PATH:
						case RAW_STRING_LIT:
						case STRING:
							{
							setState(347);
							resource();
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(350); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Download_fnContext extends ParserRuleContext {
		public TerminalNode DOWNLOAD() { return getToken(KlangParser.DOWNLOAD, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public List<String_or_idContext> string_or_id() {
			return getRuleContexts(String_or_idContext.class);
		}
		public String_or_idContext string_or_id(int i) {
			return getRuleContext(String_or_idContext.class,i);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode COMMA() { return getToken(KlangParser.COMMA, 0); }
		public Download_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_download_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterDownload_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitDownload_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitDownload_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Download_fnContext download_fn() throws RecognitionException {
		Download_fnContext _localctx = new Download_fnContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_download_fn);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(354);
			match(DOWNLOAD);
			setState(355);
			match(OPAR);
			setState(356);
			string_or_id();
			setState(359);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(357);
				match(COMMA);
				setState(358);
				string_or_id();
				}
			}

			setState(361);
			match(CPAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Json_select_fnContext extends ParserRuleContext {
		public TerminalNode JSONSELECT() { return getToken(KlangParser.JSONSELECT, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public TerminalNode COMMA() { return getToken(KlangParser.COMMA, 0); }
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public Json_select_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_json_select_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterJson_select_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitJson_select_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitJson_select_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Json_select_fnContext json_select_fn() throws RecognitionException {
		Json_select_fnContext _localctx = new Json_select_fnContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_json_select_fn);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(JSONSELECT);
			setState(364);
			match(OPAR);
			setState(365);
			match(ID);
			setState(366);
			match(COMMA);
			setState(367);
			string_or_id();
			setState(368);
			match(CPAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Yaml_select_fnContext extends ParserRuleContext {
		public TerminalNode YAMLSELECT() { return getToken(KlangParser.YAMLSELECT, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public List<TerminalNode> COMMA() { return getTokens(KlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(KlangParser.COMMA, i);
		}
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode NUMBER() { return getToken(KlangParser.NUMBER, 0); }
		public Yaml_select_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_yaml_select_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterYaml_select_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitYaml_select_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitYaml_select_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Yaml_select_fnContext yaml_select_fn() throws RecognitionException {
		Yaml_select_fnContext _localctx = new Yaml_select_fnContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_yaml_select_fn);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(370);
			match(YAMLSELECT);
			setState(371);
			match(OPAR);
			setState(372);
			match(ID);
			setState(373);
			match(COMMA);
			setState(374);
			string_or_id();
			setState(377);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(375);
				match(COMMA);
				setState(376);
				match(NUMBER);
				}
			}

			setState(379);
			match(CPAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Load_fnContext extends ParserRuleContext {
		public TerminalNode LOAD() { return getToken(KlangParser.LOAD, 0); }
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public TerminalNode COMMA() { return getToken(KlangParser.COMMA, 0); }
		public TerminalNode STRING() { return getToken(KlangParser.STRING, 0); }
		public Load_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_load_fn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterLoad_fn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitLoad_fn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitLoad_fn(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Load_fnContext load_fn() throws RecognitionException {
		Load_fnContext _localctx = new Load_fnContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_load_fn);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(381);
			match(LOAD);
			setState(382);
			match(OPAR);
			setState(383);
			string_or_id();
			setState(386);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(384);
				match(COMMA);
				setState(385);
				match(STRING);
				}
			}

			setState(388);
			match(CPAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StepInfoContext extends ParserRuleContext {
		public TerminalNode STEPINFO() { return getToken(KlangParser.STEPINFO, 0); }
		public TerminalNode STRING() { return getToken(KlangParser.STRING, 0); }
		public TerminalNode SCOL() { return getToken(KlangParser.SCOL, 0); }
		public TerminalNode RAW_STRING_LIT() { return getToken(KlangParser.RAW_STRING_LIT, 0); }
		public StepInfoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stepInfo; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterStepInfo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitStepInfo(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitStepInfo(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StepInfoContext stepInfo() throws RecognitionException {
		StepInfoContext _localctx = new StepInfoContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_stepInfo);
		try {
			setState(396);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(390);
				match(STEPINFO);
				setState(391);
				match(STRING);
				setState(392);
				match(SCOL);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(393);
				match(STEPINFO);
				setState(394);
				match(RAW_STRING_LIT);
				setState(395);
				match(SCOL);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NsContext extends ParserRuleContext {
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public TerminalNode PATH() { return getToken(KlangParser.PATH, 0); }
		public NsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ns; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterNs(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitNs(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitNs(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NsContext ns() throws RecognitionException {
		NsContext _localctx = new NsContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_ns);
		try {
			setState(400);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
			case RAW_STRING_LIT:
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(398);
				string_or_id();
				}
				break;
			case PATH:
				enterOuterAlt(_localctx, 2);
				{
				setState(399);
				match(PATH);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Patch_typeContext extends ParserRuleContext {
		public TerminalNode PATH() { return getToken(KlangParser.PATH, 0); }
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public Patch_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_patch_type; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterPatch_type(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitPatch_type(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitPatch_type(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Patch_typeContext patch_type() throws RecognitionException {
		Patch_typeContext _localctx = new Patch_typeContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_patch_type);
		try {
			setState(404);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PATH:
				enterOuterAlt(_localctx, 1);
				{
				setState(402);
				match(PATH);
				}
				break;
			case ID:
			case RAW_STRING_LIT:
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(403);
				string_or_id();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class String_or_idContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public TerminalNode STRING() { return getToken(KlangParser.STRING, 0); }
		public TerminalNode RAW_STRING_LIT() { return getToken(KlangParser.RAW_STRING_LIT, 0); }
		public String_or_idContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string_or_id; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterString_or_id(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitString_or_id(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitString_or_id(this);
			else return visitor.visitChildren(this);
		}
	}

	public final String_or_idContext string_or_id() throws RecognitionException {
		String_or_idContext _localctx = new String_or_idContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_string_or_id);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(406);
			_la = _input.LA(1);
			if ( !(((((_la - 63)) & ~0x3f) == 0 && ((1L << (_la - 63)) & ((1L << (ID - 63)) | (1L << (RAW_STRING_LIT - 63)) | (1L << (STRING - 63)))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ResourceContext extends ParserRuleContext {
		public TerminalNode PATH() { return getToken(KlangParser.PATH, 0); }
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public ResourceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_resource; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterResource(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitResource(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitResource(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ResourceContext resource() throws RecognitionException {
		ResourceContext _localctx = new ResourceContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_resource);
		try {
			setState(410);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PATH:
				enterOuterAlt(_localctx, 1);
				{
				setState(408);
				match(PATH);
				}
				break;
			case ID:
			case RAW_STRING_LIT:
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(409);
				string_or_id();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Kubernetes_object_configContext extends ParserRuleContext {
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public Kubernetes_object_configContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_kubernetes_object_config; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterKubernetes_object_config(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitKubernetes_object_config(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitKubernetes_object_config(this);
			else return visitor.visitChildren(this);
		}
	}

	public final Kubernetes_object_configContext kubernetes_object_config() throws RecognitionException {
		Kubernetes_object_configContext _localctx = new Kubernetes_object_configContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_kubernetes_object_config);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(412);
			string_or_id();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FilterContext extends ParserRuleContext {
		public TerminalNode FILTER() { return getToken(KlangParser.FILTER, 0); }
		public TerminalNode ASSIGN() { return getToken(KlangParser.ASSIGN, 0); }
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public FilterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_filter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterFilter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitFilter(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitFilter(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FilterContext filter() throws RecognitionException {
		FilterContext _localctx = new FilterContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_filter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(414);
			match(FILTER);
			setState(415);
			match(ASSIGN);
			setState(416);
			string_or_id();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PatternContext extends ParserRuleContext {
		public TerminalNode PATTERN() { return getToken(KlangParser.PATTERN, 0); }
		public TerminalNode ASSIGN() { return getToken(KlangParser.ASSIGN, 0); }
		public String_or_idContext string_or_id() {
			return getRuleContext(String_or_idContext.class,0);
		}
		public PatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterPattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitPattern(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitPattern(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PatternContext pattern() throws RecognitionException {
		PatternContext _localctx = new PatternContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_pattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(418);
			match(PATTERN);
			setState(419);
			match(ASSIGN);
			setState(420);
			string_or_id();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class JsonSelectFnContext extends ExprContext {
		public Json_select_fnContext json_select_fn() {
			return getRuleContext(Json_select_fnContext.class,0);
		}
		public JsonSelectFnContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterJsonSelectFn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitJsonSelectFn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitJsonSelectFn(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class ShellScriptContext extends ExprContext {
		public Shell_scriptContext shell_script() {
			return getRuleContext(Shell_scriptContext.class,0);
		}
		public ShellScriptContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterShellScript(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitShellScript(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitShellScript(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class YamlSelectFnContext extends ExprContext {
		public Yaml_select_fnContext yaml_select_fn() {
			return getRuleContext(Yaml_select_fnContext.class,0);
		}
		public YamlSelectFnContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterYamlSelectFn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitYamlSelectFn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitYamlSelectFn(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class AtomExprContext extends ExprContext {
		public AtomContext atom() {
			return getRuleContext(AtomContext.class,0);
		}
		public AtomExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterAtomExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitAtomExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitAtomExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class OrExprContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OR() { return getToken(KlangParser.OR, 0); }
		public OrExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterOrExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitOrExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitOrExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class AdditiveExprContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode PLUS() { return getToken(KlangParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(KlangParser.MINUS, 0); }
		public AdditiveExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterAdditiveExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitAdditiveExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitAdditiveExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class RelationalExprContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode LTEQ() { return getToken(KlangParser.LTEQ, 0); }
		public TerminalNode GTEQ() { return getToken(KlangParser.GTEQ, 0); }
		public TerminalNode LT() { return getToken(KlangParser.LT, 0); }
		public TerminalNode GT() { return getToken(KlangParser.GT, 0); }
		public RelationalExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterRelationalExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitRelationalExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitRelationalExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class NotExprContext extends ExprContext {
		public TerminalNode NOT() { return getToken(KlangParser.NOT, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NotExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterNotExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitNotExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitNotExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class UnaryMinusExprContext extends ExprContext {
		public TerminalNode MINUS() { return getToken(KlangParser.MINUS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public UnaryMinusExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterUnaryMinusExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitUnaryMinusExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitUnaryMinusExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class MultiplicationExprContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode MULT() { return getToken(KlangParser.MULT, 0); }
		public TerminalNode DIV() { return getToken(KlangParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(KlangParser.MOD, 0); }
		public MultiplicationExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterMultiplicationExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitMultiplicationExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitMultiplicationExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class DownloadFnContext extends ExprContext {
		public Download_fnContext download_fn() {
			return getRuleContext(Download_fnContext.class,0);
		}
		public DownloadFnContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterDownloadFn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitDownloadFn(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitDownloadFn(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class PowExprContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode POW() { return getToken(KlangParser.POW, 0); }
		public PowExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterPowExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitPowExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitPowExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class KubectlExprContext extends ExprContext {
		public Kubectl_commandContext kubectl_command() {
			return getRuleContext(Kubectl_commandContext.class,0);
		}
		public KubectlExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterKubectlExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitKubectlExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitKubectlExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class EqualityExprContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode EQ() { return getToken(KlangParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(KlangParser.NEQ, 0); }
		public EqualityExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterEqualityExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitEqualityExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitEqualityExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class AndExprContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode AND() { return getToken(KlangParser.AND, 0); }
		public AndExprContext(ExprContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterAndExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitAndExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitAndExpr(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 66;
		enterRecursionRule(_localctx, 66, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(433);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MINUS:
				{
				_localctx = new UnaryMinusExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(423);
				match(MINUS);
				setState(424);
				expr(14);
				}
				break;
			case NOT:
				{
				_localctx = new NotExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(425);
				match(NOT);
				setState(426);
				expr(13);
				}
				break;
			case KUBECTL:
				{
				_localctx = new KubectlExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(427);
				kubectl_command();
				}
				break;
			case JSONSELECT:
				{
				_localctx = new JsonSelectFnContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(428);
				json_select_fn();
				}
				break;
			case YAMLSELECT:
				{
				_localctx = new YamlSelectFnContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(429);
				yaml_select_fn();
				}
				break;
			case SHELLSCRIPT:
				{
				_localctx = new ShellScriptContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(430);
				shell_script();
				}
				break;
			case DOWNLOAD:
				{
				_localctx = new DownloadFnContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(431);
				download_fn();
				}
				break;
			case T__1:
			case T__3:
			case OPAR:
			case OBRACE:
			case TRUE:
			case FALSE:
			case NIL:
			case ID:
			case NUMBER:
			case RAW_STRING_LIT:
			case STRING:
				{
				_localctx = new AtomExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(432);
				atom();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(458);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(456);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
					case 1:
						{
						_localctx = new PowExprContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(435);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(436);
						match(POW);
						setState(437);
						expr(15);
						}
						break;
					case 2:
						{
						_localctx = new MultiplicationExprContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(438);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(439);
						((MultiplicationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MULT) | (1L << DIV) | (1L << MOD))) != 0)) ) {
							((MultiplicationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(440);
						expr(13);
						}
						break;
					case 3:
						{
						_localctx = new AdditiveExprContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(441);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(442);
						((AdditiveExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((AdditiveExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(443);
						expr(12);
						}
						break;
					case 4:
						{
						_localctx = new RelationalExprContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(444);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(445);
						((RelationalExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << GT) | (1L << LT) | (1L << GTEQ) | (1L << LTEQ))) != 0)) ) {
							((RelationalExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(446);
						expr(11);
						}
						break;
					case 5:
						{
						_localctx = new EqualityExprContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(447);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(448);
						((EqualityExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==EQ || _la==NEQ) ) {
							((EqualityExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(449);
						expr(10);
						}
						break;
					case 6:
						{
						_localctx = new AndExprContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(450);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(451);
						match(AND);
						setState(452);
						expr(9);
						}
						break;
					case 7:
						{
						_localctx = new OrExprContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(453);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(454);
						match(OR);
						setState(455);
						expr(8);
						}
						break;
					}
					} 
				}
				setState(460);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AtomContext extends ParserRuleContext {
		public AtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atom; }
	 
		public AtomContext() { }
		public void copyFrom(AtomContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ParExprContext extends AtomContext {
		public TerminalNode OPAR() { return getToken(KlangParser.OPAR, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CPAR() { return getToken(KlangParser.CPAR, 0); }
		public ParExprContext(AtomContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterParExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitParExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitParExpr(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class BooleanAtomContext extends AtomContext {
		public TerminalNode TRUE() { return getToken(KlangParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(KlangParser.FALSE, 0); }
		public BooleanAtomContext(AtomContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterBooleanAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitBooleanAtom(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitBooleanAtom(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class IdAtomContext extends AtomContext {
		public TerminalNode ID() { return getToken(KlangParser.ID, 0); }
		public IdAtomContext(AtomContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterIdAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitIdAtom(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitIdAtom(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class StringAtomContext extends AtomContext {
		public TerminalNode STRING() { return getToken(KlangParser.STRING, 0); }
		public StringAtomContext(AtomContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterStringAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitStringAtom(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitStringAtom(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class RawStringAtomContext extends AtomContext {
		public TerminalNode RAW_STRING_LIT() { return getToken(KlangParser.RAW_STRING_LIT, 0); }
		public RawStringAtomContext(AtomContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterRawStringAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitRawStringAtom(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitRawStringAtom(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class JsonAtomContext extends AtomContext {
		public JsonContext json() {
			return getRuleContext(JsonContext.class,0);
		}
		public JsonAtomContext(AtomContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterJsonAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitJsonAtom(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitJsonAtom(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class NilAtomContext extends AtomContext {
		public TerminalNode NIL() { return getToken(KlangParser.NIL, 0); }
		public NilAtomContext(AtomContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterNilAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitNilAtom(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitNilAtom(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class NumberAtomContext extends AtomContext {
		public TerminalNode NUMBER() { return getToken(KlangParser.NUMBER, 0); }
		public NumberAtomContext(AtomContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterNumberAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitNumberAtom(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitNumberAtom(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AtomContext atom() throws RecognitionException {
		AtomContext _localctx = new AtomContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_atom);
		int _la;
		try {
			setState(472);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				_localctx = new ParExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(461);
				match(OPAR);
				setState(462);
				expr(0);
				setState(463);
				match(CPAR);
				}
				break;
			case 2:
				_localctx = new NumberAtomContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(465);
				match(NUMBER);
				}
				break;
			case 3:
				_localctx = new BooleanAtomContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(466);
				_la = _input.LA(1);
				if ( !(_la==TRUE || _la==FALSE) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 4:
				_localctx = new RawStringAtomContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(467);
				match(RAW_STRING_LIT);
				}
				break;
			case 5:
				_localctx = new IdAtomContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(468);
				match(ID);
				}
				break;
			case 6:
				_localctx = new StringAtomContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(469);
				match(STRING);
				}
				break;
			case 7:
				_localctx = new JsonAtomContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(470);
				json();
				}
				break;
			case 8:
				_localctx = new NilAtomContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(471);
				match(NIL);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class JsonContext extends ParserRuleContext {
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public JsonContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_json; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterJson(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitJson(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitJson(this);
			else return visitor.visitChildren(this);
		}
	}

	public final JsonContext json() throws RecognitionException {
		JsonContext _localctx = new JsonContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_json);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(474);
			value();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ObjContext extends ParserRuleContext {
		public TerminalNode OBRACE() { return getToken(KlangParser.OBRACE, 0); }
		public List<PairContext> pair() {
			return getRuleContexts(PairContext.class);
		}
		public PairContext pair(int i) {
			return getRuleContext(PairContext.class,i);
		}
		public TerminalNode CBRACE() { return getToken(KlangParser.CBRACE, 0); }
		public List<TerminalNode> COMMA() { return getTokens(KlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(KlangParser.COMMA, i);
		}
		public ObjContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_obj; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterObj(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitObj(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitObj(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ObjContext obj() throws RecognitionException {
		ObjContext _localctx = new ObjContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_obj);
		int _la;
		try {
			setState(489);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(476);
				match(OBRACE);
				setState(477);
				pair();
				setState(482);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(478);
					match(COMMA);
					setState(479);
					pair();
					}
					}
					setState(484);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(485);
				match(CBRACE);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(487);
				match(OBRACE);
				setState(488);
				match(CBRACE);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PairContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(KlangParser.STRING, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public PairContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pair; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterPair(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitPair(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitPair(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PairContext pair() throws RecognitionException {
		PairContext _localctx = new PairContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_pair);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(491);
			match(STRING);
			setState(492);
			match(T__0);
			setState(493);
			value();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArrContext extends ParserRuleContext {
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(KlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(KlangParser.COMMA, i);
		}
		public ArrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterArr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitArr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitArr(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArrContext arr() throws RecognitionException {
		ArrContext _localctx = new ArrContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_arr);
		int _la;
		try {
			setState(508);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(495);
				match(T__1);
				setState(496);
				value();
				setState(501);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(497);
					match(COMMA);
					setState(498);
					value();
					}
					}
					setState(503);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(504);
				match(T__2);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(506);
				match(T__1);
				setState(507);
				match(T__2);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValueContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(KlangParser.STRING, 0); }
		public TerminalNode NUMBER() { return getToken(KlangParser.NUMBER, 0); }
		public ObjContext obj() {
			return getRuleContext(ObjContext.class,0);
		}
		public ArrContext arr() {
			return getRuleContext(ArrContext.class,0);
		}
		public TerminalNode TRUE() { return getToken(KlangParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(KlangParser.FALSE, 0); }
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).enterValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof KlangListener ) ((KlangListener)listener).exitValue(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof KlangVisitor ) return ((KlangVisitor<? extends T>)visitor).visitValue(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_value);
		try {
			setState(517);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(510);
				match(STRING);
				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 2);
				{
				setState(511);
				match(NUMBER);
				}
				break;
			case OBRACE:
				enterOuterAlt(_localctx, 3);
				{
				setState(512);
				obj();
				}
				break;
			case T__1:
				enterOuterAlt(_localctx, 4);
				{
				setState(513);
				arr();
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 5);
				{
				setState(514);
				match(TRUE);
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 6);
				{
				setState(515);
				match(FALSE);
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 7);
				{
				setState(516);
				match(T__3);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 33:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 15);
		case 1:
			return precpred(_ctx, 12);
		case 2:
			return precpred(_ctx, 11);
		case 3:
			return precpred(_ctx, 10);
		case 4:
			return precpred(_ctx, 9);
		case 5:
			return precpred(_ctx, 8);
		case 6:
			return precpred(_ctx, 7);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3H\u020a\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\3\2\3\2\3\2\3"+
		"\3\7\3W\n\3\f\3\16\3Z\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\5\4l\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\5\5x\n\5\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u0098"+
		"\n\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u00a4\n\n\3\n\3\n\3\n"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u00b2\n\13\3\13\3\13"+
		"\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\5\f\u00db\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00e6"+
		"\n\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u010f\n\16"+
		"\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\7\21"+
		"\u011e\n\21\f\21\16\21\u0121\13\21\3\21\3\21\5\21\u0125\n\21\3\22\3\22"+
		"\3\22\3\23\3\23\3\23\3\23\3\23\5\23\u012f\n\23\3\24\3\24\3\24\3\24\3\25"+
		"\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\6\26\u0140\n\26\r\26"+
		"\16\26\u0141\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\6\26\u014d\n"+
		"\26\r\26\16\26\u014e\3\26\3\26\3\26\3\26\3\26\6\26\u0156\n\26\r\26\16"+
		"\26\u0157\3\26\3\26\3\26\3\26\3\26\6\26\u015f\n\26\r\26\16\26\u0160\5"+
		"\26\u0163\n\26\3\27\3\27\3\27\3\27\3\27\5\27\u016a\n\27\3\27\3\27\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31"+
		"\u017c\n\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\5\32\u0185\n\32\3\32\3"+
		"\32\3\33\3\33\3\33\3\33\3\33\3\33\5\33\u018f\n\33\3\34\3\34\5\34\u0193"+
		"\n\34\3\35\3\35\5\35\u0197\n\35\3\36\3\36\3\37\3\37\5\37\u019d\n\37\3"+
		" \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\5#\u01b4"+
		"\n#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\7#"+
		"\u01cb\n#\f#\16#\u01ce\13#\3$\3$\3$\3$\3$\3$\3$\3$\3$\3$\3$\5$\u01db\n"+
		"$\3%\3%\3&\3&\3&\3&\7&\u01e3\n&\f&\16&\u01e6\13&\3&\3&\3&\3&\5&\u01ec"+
		"\n&\3\'\3\'\3\'\3\'\3(\3(\3(\3(\7(\u01f6\n(\f(\16(\u01f9\13(\3(\3(\3("+
		"\3(\5(\u01ff\n(\3)\3)\3)\3)\3)\3)\3)\5)\u0208\n)\3)\2\3D*\2\4\6\b\n\f"+
		"\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNP\2\b\4\2"+
		"AADE\3\2\21\23\3\2\17\20\3\2\13\16\3\2\t\n\3\2\35\36\2\u0232\2R\3\2\2"+
		"\2\4X\3\2\2\2\6k\3\2\2\2\bw\3\2\2\2\ny\3\2\2\2\f|\3\2\2\2\16\u0086\3\2"+
		"\2\2\20\u008e\3\2\2\2\22\u009c\3\2\2\2\24\u00a8\3\2\2\2\26\u00da\3\2\2"+
		"\2\30\u00dc\3\2\2\2\32\u010e\3\2\2\2\34\u0110\3\2\2\2\36\u0114\3\2\2\2"+
		" \u0118\3\2\2\2\"\u0126\3\2\2\2$\u012e\3\2\2\2&\u0130\3\2\2\2(\u0134\3"+
		"\2\2\2*\u0162\3\2\2\2,\u0164\3\2\2\2.\u016d\3\2\2\2\60\u0174\3\2\2\2\62"+
		"\u017f\3\2\2\2\64\u018e\3\2\2\2\66\u0192\3\2\2\28\u0196\3\2\2\2:\u0198"+
		"\3\2\2\2<\u019c\3\2\2\2>\u019e\3\2\2\2@\u01a0\3\2\2\2B\u01a4\3\2\2\2D"+
		"\u01b3\3\2\2\2F\u01da\3\2\2\2H\u01dc\3\2\2\2J\u01eb\3\2\2\2L\u01ed\3\2"+
		"\2\2N\u01fe\3\2\2\2P\u0207\3\2\2\2RS\5\4\3\2ST\7\2\2\3T\3\3\2\2\2UW\5"+
		"\6\4\2VU\3\2\2\2WZ\3\2\2\2XV\3\2\2\2XY\3\2\2\2Y\5\3\2\2\2ZX\3\2\2\2[l"+
		"\5\b\5\2\\l\5\f\7\2]l\5\16\b\2^l\5\20\t\2_l\5\22\n\2`l\5\26\f\2al\5\24"+
		"\13\2bl\5\32\16\2cl\5\30\r\2dl\5 \21\2el\5&\24\2fl\5\34\17\2gl\5\36\20"+
		"\2hl\5(\25\2ij\7H\2\2jl\b\4\1\2k[\3\2\2\2k\\\3\2\2\2k]\3\2\2\2k^\3\2\2"+
		"\2k_\3\2\2\2k`\3\2\2\2ka\3\2\2\2kb\3\2\2\2kc\3\2\2\2kd\3\2\2\2ke\3\2\2"+
		"\2kf\3\2\2\2kg\3\2\2\2kh\3\2\2\2ki\3\2\2\2l\7\3\2\2\2mn\7A\2\2no\7\27"+
		"\2\2op\5D#\2pq\7\26\2\2qx\3\2\2\2rs\7A\2\2st\7\27\2\2tu\5\62\32\2uv\7"+
		"\26\2\2vx\3\2\2\2wm\3\2\2\2wr\3\2\2\2x\t\3\2\2\2yz\7;\2\2z{\5:\36\2{\13"+
		"\3\2\2\2|}\7\62\2\2}~\7\30\2\2~\177\7A\2\2\177\u0080\7\34\2\2\u0080\u0081"+
		"\5:\36\2\u0081\u0082\7\34\2\2\u0082\u0083\5D#\2\u0083\u0084\7\31\2\2\u0084"+
		"\u0085\7\26\2\2\u0085\r\3\2\2\2\u0086\u0087\7\63\2\2\u0087\u0088\7\30"+
		"\2\2\u0088\u0089\7A\2\2\u0089\u008a\7\34\2\2\u008a\u008b\5:\36\2\u008b"+
		"\u008c\7\31\2\2\u008c\u008d\7\26\2\2\u008d\17\3\2\2\2\u008e\u008f\7\65"+
		"\2\2\u008f\u0090\7\30\2\2\u0090\u0091\7A\2\2\u0091\u0092\7\34\2\2\u0092"+
		"\u0093\5:\36\2\u0093\u0094\7\34\2\2\u0094\u0097\5D#\2\u0095\u0096\7\34"+
		"\2\2\u0096\u0098\7B\2\2\u0097\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098"+
		"\u0099\3\2\2\2\u0099\u009a\7\31\2\2\u009a\u009b\7\26\2\2\u009b\21\3\2"+
		"\2\2\u009c\u009d\7\66\2\2\u009d\u009e\7\30\2\2\u009e\u009f\7A\2\2\u009f"+
		"\u00a0\7\34\2\2\u00a0\u00a3\5:\36\2\u00a1\u00a2\7\34\2\2\u00a2\u00a4\7"+
		"B\2\2\u00a3\u00a1\3\2\2\2\u00a3\u00a4\3\2\2\2\u00a4\u00a5\3\2\2\2\u00a5"+
		"\u00a6\7\31\2\2\u00a6\u00a7\7\26\2\2\u00a7\23\3\2\2\2\u00a8\u00a9\7\67"+
		"\2\2\u00a9\u00aa\7\30\2\2\u00aa\u00ab\7A\2\2\u00ab\u00ac\7\34\2\2\u00ac"+
		"\u00ad\5:\36\2\u00ad\u00ae\7\34\2\2\u00ae\u00b1\5D#\2\u00af\u00b0\7\34"+
		"\2\2\u00b0\u00b2\5:\36\2\u00b1\u00af\3\2\2\2\u00b1\u00b2\3\2\2\2\u00b2"+
		"\u00b3\3\2\2\2\u00b3\u00b4\7\31\2\2\u00b4\u00b5\7\26\2\2\u00b5\25\3\2"+
		"\2\2\u00b6\u00b7\78\2\2\u00b7\u00b8\7\30\2\2\u00b8\u00b9\7A\2\2\u00b9"+
		"\u00ba\7\34\2\2\u00ba\u00bb\5@!\2\u00bb\u00bc\7\34\2\2\u00bc\u00bd\5B"+
		"\"\2\u00bd\u00be\7\31\2\2\u00be\u00bf\7\26\2\2\u00bf\u00db\3\2\2\2\u00c0"+
		"\u00c1\78\2\2\u00c1\u00c2\7\30\2\2\u00c2\u00c3\7A\2\2\u00c3\u00c4\7\34"+
		"\2\2\u00c4\u00c5\5B\"\2\u00c5\u00c6\7\34\2\2\u00c6\u00c7\5@!\2\u00c7\u00c8"+
		"\7\31\2\2\u00c8\u00c9\7\26\2\2\u00c9\u00db\3\2\2\2\u00ca\u00cb\78\2\2"+
		"\u00cb\u00cc\7\30\2\2\u00cc\u00cd\7A\2\2\u00cd\u00ce\7\34\2\2\u00ce\u00cf"+
		"\5@!\2\u00cf\u00d0\7\31\2\2\u00d0\u00d1\7\26\2\2\u00d1\u00db\3\2\2\2\u00d2"+
		"\u00d3\78\2\2\u00d3\u00d4\7\30\2\2\u00d4\u00d5\7A\2\2\u00d5\u00d6\7\34"+
		"\2\2\u00d6\u00d7\5B\"\2\u00d7\u00d8\7\31\2\2\u00d8\u00d9\7\26\2\2\u00d9"+
		"\u00db\3\2\2\2\u00da\u00b6\3\2\2\2\u00da\u00c0\3\2\2\2\u00da\u00ca\3\2"+
		"\2\2\u00da\u00d2\3\2\2\2\u00db\27\3\2\2\2\u00dc\u00dd\79\2\2\u00dd\u00de"+
		"\7\30\2\2\u00de\u00df\7A\2\2\u00df\u00e0\7\34\2\2\u00e0\u00e1\5:\36\2"+
		"\u00e1\u00e2\7\34\2\2\u00e2\u00e5\5D#\2\u00e3\u00e4\7\34\2\2\u00e4\u00e6"+
		"\5:\36\2\u00e5\u00e3\3\2\2\2\u00e5\u00e6\3\2\2\2\u00e6\u00e7\3\2\2\2\u00e7"+
		"\u00e8\7\31\2\2\u00e8\u00e9\7\26\2\2\u00e9\31\3\2\2\2\u00ea\u00eb\7:\2"+
		"\2\u00eb\u00ec\7\30\2\2\u00ec\u00ed\7A\2\2\u00ed\u00ee\7\34\2\2\u00ee"+
		"\u00ef\5@!\2\u00ef\u00f0\7\34\2\2\u00f0\u00f1\5B\"\2\u00f1\u00f2\7\31"+
		"\2\2\u00f2\u00f3\7\26\2\2\u00f3\u010f\3\2\2\2\u00f4\u00f5\7:\2\2\u00f5"+
		"\u00f6\7\30\2\2\u00f6\u00f7\7A\2\2\u00f7\u00f8\7\34\2\2\u00f8\u00f9\5"+
		"B\"\2\u00f9\u00fa\7\34\2\2\u00fa\u00fb\5@!\2\u00fb\u00fc\7\31\2\2\u00fc"+
		"\u00fd\7\26\2\2\u00fd\u010f\3\2\2\2\u00fe\u00ff\7:\2\2\u00ff\u0100\7\30"+
		"\2\2\u0100\u0101\7A\2\2\u0101\u0102\7\34\2\2\u0102\u0103\5@!\2\u0103\u0104"+
		"\7\31\2\2\u0104\u0105\7\26\2\2\u0105\u010f\3\2\2\2\u0106\u0107\7:\2\2"+
		"\u0107\u0108\7\30\2\2\u0108\u0109\7A\2\2\u0109\u010a\7\34\2\2\u010a\u010b"+
		"\5B\"\2\u010b\u010c\7\31\2\2\u010c\u010d\7\26\2\2\u010d\u010f\3\2\2\2"+
		"\u010e\u00ea\3\2\2\2\u010e\u00f4\3\2\2\2\u010e\u00fe\3\2\2\2\u010e\u0106"+
		"\3\2\2\2\u010f\33\3\2\2\2\u0110\u0111\7=\2\2\u0111\u0112\7B\2\2\u0112"+
		"\u0113\7\26\2\2\u0113\35\3\2\2\2\u0114\u0115\7\60\2\2\u0115\u0116\7B\2"+
		"\2\u0116\u0117\7\26\2\2\u0117\37\3\2\2\2\u0118\u0119\7 \2\2\u0119\u011f"+
		"\5\"\22\2\u011a\u011b\7!\2\2\u011b\u011c\7 \2\2\u011c\u011e\5\"\22\2\u011d"+
		"\u011a\3\2\2\2\u011e\u0121\3\2\2\2\u011f\u011d\3\2\2\2\u011f\u0120\3\2"+
		"\2\2\u0120\u0124\3\2\2\2\u0121\u011f\3\2\2\2\u0122\u0123\7!\2\2\u0123"+
		"\u0125\5$\23\2\u0124\u0122\3\2\2\2\u0124\u0125\3\2\2\2\u0125!\3\2\2\2"+
		"\u0126\u0127\5D#\2\u0127\u0128\5$\23\2\u0128#\3\2\2\2\u0129\u012a\7\32"+
		"\2\2\u012a\u012b\5\4\3\2\u012b\u012c\7\33\2\2\u012c\u012f\3\2\2\2\u012d"+
		"\u012f\5\6\4\2\u012e\u0129\3\2\2\2\u012e\u012d\3\2\2\2\u012f%\3\2\2\2"+
		"\u0130\u0131\7\"\2\2\u0131\u0132\5D#\2\u0132\u0133\5$\23\2\u0133\'\3\2"+
		"\2\2\u0134\u0135\7#\2\2\u0135\u0136\5D#\2\u0136\u0137\7\26\2\2\u0137)"+
		"\3\2\2\2\u0138\u0139\7$\2\2\u0139\u013f\7%\2\2\u013a\u013b\7*\2\2\u013b"+
		"\u0140\5\66\34\2\u013c\u0140\5:\36\2\u013d\u013e\7-\2\2\u013e\u0140\5"+
		"> \2\u013f\u013a\3\2\2\2\u013f\u013c\3\2\2\2\u013f\u013d\3\2\2\2\u0140"+
		"\u0141\3\2\2\2\u0141\u013f\3\2\2\2\u0141\u0142\3\2\2\2\u0142\u0163\3\2"+
		"\2\2\u0143\u0144\7$\2\2\u0144\u014c\7&\2\2\u0145\u0146\7*\2\2\u0146\u014d"+
		"\5\66\34\2\u0147\u014d\5<\37\2\u0148\u0149\7+\2\2\u0149\u014d\58\35\2"+
		"\u014a\u014b\7,\2\2\u014b\u014d\5:\36\2\u014c\u0145\3\2\2\2\u014c\u0147"+
		"\3\2\2\2\u014c\u0148\3\2\2\2\u014c\u014a\3\2\2\2\u014d\u014e\3\2\2\2\u014e"+
		"\u014c\3\2\2\2\u014e\u014f\3\2\2\2\u014f\u0163\3\2\2\2\u0150\u0151\7$"+
		"\2\2\u0151\u0155\7\'\2\2\u0152\u0153\7*\2\2\u0153\u0156\5\66\34\2\u0154"+
		"\u0156\5<\37\2\u0155\u0152\3\2\2\2\u0155\u0154\3\2\2\2\u0156\u0157\3\2"+
		"\2\2\u0157\u0155\3\2\2\2\u0157\u0158\3\2\2\2\u0158\u0163\3\2\2\2\u0159"+
		"\u015a\7$\2\2\u015a\u015e\7)\2\2\u015b\u015c\7*\2\2\u015c\u015f\5\66\34"+
		"\2\u015d\u015f\5<\37\2\u015e\u015b\3\2\2\2\u015e\u015d\3\2\2\2\u015f\u0160"+
		"\3\2\2\2\u0160\u015e\3\2\2\2\u0160\u0161\3\2\2\2\u0161\u0163\3\2\2\2\u0162"+
		"\u0138\3\2\2\2\u0162\u0143\3\2\2\2\u0162\u0150\3\2\2\2\u0162\u0159\3\2"+
		"\2\2\u0163+\3\2\2\2\u0164\u0165\7<\2\2\u0165\u0166\7\30\2\2\u0166\u0169"+
		"\5:\36\2\u0167\u0168\7\34\2\2\u0168\u016a\5:\36\2\u0169\u0167\3\2\2\2"+
		"\u0169\u016a\3\2\2\2\u016a\u016b\3\2\2\2\u016b\u016c\7\31\2\2\u016c-\3"+
		"\2\2\2\u016d\u016e\7\61\2\2\u016e\u016f\7\30\2\2\u016f\u0170\7A\2\2\u0170"+
		"\u0171\7\34\2\2\u0171\u0172\5:\36\2\u0172\u0173\7\31\2\2\u0173/\3\2\2"+
		"\2\u0174\u0175\7\64\2\2\u0175\u0176\7\30\2\2\u0176\u0177\7A\2\2\u0177"+
		"\u0178\7\34\2\2\u0178\u017b\5:\36\2\u0179\u017a\7\34\2\2\u017a\u017c\7"+
		"B\2\2\u017b\u0179\3\2\2\2\u017b\u017c\3\2\2\2\u017c\u017d\3\2\2\2\u017d"+
		"\u017e\7\31\2\2\u017e\61\3\2\2\2\u017f\u0180\7/\2\2\u0180\u0181\7\30\2"+
		"\2\u0181\u0184\5:\36\2\u0182\u0183\7\34\2\2\u0183\u0185\7E\2\2\u0184\u0182"+
		"\3\2\2\2\u0184\u0185\3\2\2\2\u0185\u0186\3\2\2\2\u0186\u0187\7\31\2\2"+
		"\u0187\63\3\2\2\2\u0188\u0189\7>\2\2\u0189\u018a\7E\2\2\u018a\u018f\7"+
		"\26\2\2\u018b\u018c\7>\2\2\u018c\u018d\7D\2\2\u018d\u018f\7\26\2\2\u018e"+
		"\u0188\3\2\2\2\u018e\u018b\3\2\2\2\u018f\65\3\2\2\2\u0190\u0193\5:\36"+
		"\2\u0191\u0193\7C\2\2\u0192\u0190\3\2\2\2\u0192\u0191\3\2\2\2\u0193\67"+
		"\3\2\2\2\u0194\u0197\7C\2\2\u0195\u0197\5:\36\2\u0196\u0194\3\2\2\2\u0196"+
		"\u0195\3\2\2\2\u01979\3\2\2\2\u0198\u0199\t\2\2\2\u0199;\3\2\2\2\u019a"+
		"\u019d\7C\2\2\u019b\u019d\5:\36\2\u019c\u019a\3\2\2\2\u019c\u019b\3\2"+
		"\2\2\u019d=\3\2\2\2\u019e\u019f\5:\36\2\u019f?\3\2\2\2\u01a0\u01a1\7?"+
		"\2\2\u01a1\u01a2\7\27\2\2\u01a2\u01a3\5:\36\2\u01a3A\3\2\2\2\u01a4\u01a5"+
		"\7@\2\2\u01a5\u01a6\7\27\2\2\u01a6\u01a7\5:\36\2\u01a7C\3\2\2\2\u01a8"+
		"\u01a9\b#\1\2\u01a9\u01aa\7\20\2\2\u01aa\u01b4\5D#\20\u01ab\u01ac\7\25"+
		"\2\2\u01ac\u01b4\5D#\17\u01ad\u01b4\5*\26\2\u01ae\u01b4\5.\30\2\u01af"+
		"\u01b4\5\60\31\2\u01b0\u01b4\5\n\6\2\u01b1\u01b4\5,\27\2\u01b2\u01b4\5"+
		"F$\2\u01b3\u01a8\3\2\2\2\u01b3\u01ab\3\2\2\2\u01b3\u01ad\3\2\2\2\u01b3"+
		"\u01ae\3\2\2\2\u01b3\u01af\3\2\2\2\u01b3\u01b0\3\2\2\2\u01b3\u01b1\3\2"+
		"\2\2\u01b3\u01b2\3\2\2\2\u01b4\u01cc\3\2\2\2\u01b5\u01b6\f\21\2\2\u01b6"+
		"\u01b7\7\24\2\2\u01b7\u01cb\5D#\21\u01b8\u01b9\f\16\2\2\u01b9\u01ba\t"+
		"\3\2\2\u01ba\u01cb\5D#\17\u01bb\u01bc\f\r\2\2\u01bc\u01bd\t\4\2\2\u01bd"+
		"\u01cb\5D#\16\u01be\u01bf\f\f\2\2\u01bf\u01c0\t\5\2\2\u01c0\u01cb\5D#"+
		"\r\u01c1\u01c2\f\13\2\2\u01c2\u01c3\t\6\2\2\u01c3\u01cb\5D#\f\u01c4\u01c5"+
		"\f\n\2\2\u01c5\u01c6\7\b\2\2\u01c6\u01cb\5D#\13\u01c7\u01c8\f\t\2\2\u01c8"+
		"\u01c9\7\7\2\2\u01c9\u01cb\5D#\n\u01ca\u01b5\3\2\2\2\u01ca\u01b8\3\2\2"+
		"\2\u01ca\u01bb\3\2\2\2\u01ca\u01be\3\2\2\2\u01ca\u01c1\3\2\2\2\u01ca\u01c4"+
		"\3\2\2\2\u01ca\u01c7\3\2\2\2\u01cb\u01ce\3\2\2\2\u01cc\u01ca\3\2\2\2\u01cc"+
		"\u01cd\3\2\2\2\u01cdE\3\2\2\2\u01ce\u01cc\3\2\2\2\u01cf\u01d0\7\30\2\2"+
		"\u01d0\u01d1\5D#\2\u01d1\u01d2\7\31\2\2\u01d2\u01db\3\2\2\2\u01d3\u01db"+
		"\7B\2\2\u01d4\u01db\t\7\2\2\u01d5\u01db\7D\2\2\u01d6\u01db\7A\2\2\u01d7"+
		"\u01db\7E\2\2\u01d8\u01db\5H%\2\u01d9\u01db\7\37\2\2\u01da\u01cf\3\2\2"+
		"\2\u01da\u01d3\3\2\2\2\u01da\u01d4\3\2\2\2\u01da\u01d5\3\2\2\2\u01da\u01d6"+
		"\3\2\2\2\u01da\u01d7\3\2\2\2\u01da\u01d8\3\2\2\2\u01da\u01d9\3\2\2\2\u01db"+
		"G\3\2\2\2\u01dc\u01dd\5P)\2\u01ddI\3\2\2\2\u01de\u01df\7\32\2\2\u01df"+
		"\u01e4\5L\'\2\u01e0\u01e1\7\34\2\2\u01e1\u01e3\5L\'\2\u01e2\u01e0\3\2"+
		"\2\2\u01e3\u01e6\3\2\2\2\u01e4\u01e2\3\2\2\2\u01e4\u01e5\3\2\2\2\u01e5"+
		"\u01e7\3\2\2\2\u01e6\u01e4\3\2\2\2\u01e7\u01e8\7\33\2\2\u01e8\u01ec\3"+
		"\2\2\2\u01e9\u01ea\7\32\2\2\u01ea\u01ec\7\33\2\2\u01eb\u01de\3\2\2\2\u01eb"+
		"\u01e9\3\2\2\2\u01ecK\3\2\2\2\u01ed\u01ee\7E\2\2\u01ee\u01ef\7\3\2\2\u01ef"+
		"\u01f0\5P)\2\u01f0M\3\2\2\2\u01f1\u01f2\7\4\2\2\u01f2\u01f7\5P)\2\u01f3"+
		"\u01f4\7\34\2\2\u01f4\u01f6\5P)\2\u01f5\u01f3\3\2\2\2\u01f6\u01f9\3\2"+
		"\2\2\u01f7\u01f5\3\2\2\2\u01f7\u01f8\3\2\2\2\u01f8\u01fa\3\2\2\2\u01f9"+
		"\u01f7\3\2\2\2\u01fa\u01fb\7\5\2\2\u01fb\u01ff\3\2\2\2\u01fc\u01fd\7\4"+
		"\2\2\u01fd\u01ff\7\5\2\2\u01fe\u01f1\3\2\2\2\u01fe\u01fc\3\2\2\2\u01ff"+
		"O\3\2\2\2\u0200\u0208\7E\2\2\u0201\u0208\7B\2\2\u0202\u0208\5J&\2\u0203"+
		"\u0208\5N(\2\u0204\u0208\7\35\2\2\u0205\u0208\7\36\2\2\u0206\u0208\7\6"+
		"\2\2\u0207\u0200\3\2\2\2\u0207\u0201\3\2\2\2\u0207\u0202\3\2\2\2\u0207"+
		"\u0203\3\2\2\2\u0207\u0204\3\2\2\2\u0207\u0205\3\2\2\2\u0207\u0206\3\2"+
		"\2\2\u0208Q\3\2\2\2\'Xkw\u0097\u00a3\u00b1\u00da\u00e5\u010e\u011f\u0124"+
		"\u012e\u013f\u0141\u014c\u014e\u0155\u0157\u015e\u0160\u0162\u0169\u017b"+
		"\u0184\u018e\u0192\u0196\u019c\u01b3\u01ca\u01cc\u01da\u01e4\u01eb\u01f7"+
		"\u01fe\u0207";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}