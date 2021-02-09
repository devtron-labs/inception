// Generated from /Users/pghildiy/go/src/github.com/devtron-labs/inception/pkg/language/grammar/Klang.g4 by ANTLR 4.9
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class KlangLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "OR", "AND", "EQ", "NEQ", "GT", "LT", 
			"GTEQ", "LTEQ", "PLUS", "MINUS", "MULT", "DIV", "MOD", "POW", "NOT", 
			"SCOL", "ASSIGN", "OPAR", "CPAR", "OBRACE", "CBRACE", "COMMA", "TRUE", 
			"FALSE", "NIL", "IF", "ELSE", "WHILE", "LOG", "KUBECTL", "APPLY", "PATCH", 
			"GET", "REPLACE", "DELETE", "NAMESPACE", "PATCHTYPE", "PATCHLOAD", "UPDATELOAD", 
			"JSONPATH", "LOAD", "EXIT", "JSONSELECT", "JSONEDIT", "JSONDELETE", "YAMLSELECT", 
			"YAMLEDIT", "YAMLDELETE", "KUBEJSONEDIT", "KUBEJSONDELETE", "KUBEYAMLEDIT", 
			"KUBEYAMLDELETE", "SHELLSCRIPT", "DOWNLOAD", "SLEEP", "STEPINFO", "FILTER", 
			"PATTERN", "ID", "NUMBER", "EXP", "INT", "PATH", "RAW_STRING_LIT", "STRING", 
			"ESCQUOTE", "SAFECODEPOINTQUOTE", "ESC", "UNICODE", "HEX", "SAFECODEPOINT", 
			"COMMENT", "SPACE", "OTHER"
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


	public KlangLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Klang.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2H\u0263\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3"+
		"\5\3\5\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\t\3\n\3\n"+
		"\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21"+
		"\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30"+
		"\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35\3\35"+
		"\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3 \3"+
		" \3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3#\3#\3$\3$\3$\3"+
		"$\3$\3$\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'"+
		"\3(\3(\3(\3(\3(\3(\3(\3)\3)\3)\3*\3*\3*\3*\3*\3*\3*\3+\3+\3+\3,\3,\3,"+
		"\3-\3-\3-\3-\3-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3/\3/\3/\3/\3/\3\60\3\60"+
		"\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61"+
		"\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\62\3\62\3\62\3\62\3\62"+
		"\3\62\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\64\3\64"+
		"\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\65\3\65\3\65"+
		"\3\65\3\65\3\65\3\65\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\66"+
		"\3\66\3\66\3\66\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67"+
		"\3\67\3\67\3\67\3\67\38\38\38\38\38\38\38\38\38\38\38\38\38\39\39\39\3"+
		"9\39\39\39\39\39\39\39\39\39\39\39\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3"+
		":\3;\3;\3;\3;\3;\3;\3;\3;\3;\3<\3<\3<\3<\3<\3<\3=\3=\3=\3=\3=\3=\3=\3"+
		"=\3=\3>\3>\3>\3>\3>\3>\3>\3?\3?\3?\3?\3?\3?\3?\3?\3@\3@\7@\u01f4\n@\f"+
		"@\16@\u01f7\13@\3A\5A\u01fa\nA\3A\3A\3A\6A\u01ff\nA\rA\16A\u0200\5A\u0203"+
		"\nA\3A\5A\u0206\nA\3B\3B\5B\u020a\nB\3B\3B\3C\3C\3C\7C\u0211\nC\fC\16"+
		"C\u0214\13C\5C\u0216\nC\3D\3D\7D\u021a\nD\fD\16D\u021d\13D\3E\3E\7E\u0221"+
		"\nE\fE\16E\u0224\13E\3E\3E\3F\3F\3F\7F\u022b\nF\fF\16F\u022e\13F\3F\3"+
		"F\3F\3F\7F\u0234\nF\fF\16F\u0237\13F\3F\5F\u023a\nF\3G\3G\3G\5G\u023f"+
		"\nG\3H\3H\3I\3I\3I\5I\u0246\nI\3J\3J\3J\3J\3J\3J\3K\3K\3L\3L\3M\3M\7M"+
		"\u0254\nM\fM\16M\u0257\13M\3M\3M\3N\6N\u025c\nN\rN\16N\u025d\3N\3N\3O"+
		"\3O\2\2P\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17"+
		"\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\35"+
		"9\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66"+
		"k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083\2\u0085\2\u0087C\u0089D\u008b"+
		"E\u008d\2\u008f\2\u0091\2\u0093\2\u0095\2\u0097\2\u0099F\u009bG\u009d"+
		"H\3\2\22\5\2C\\aac|\6\2\62;C\\aac|\3\2\62;\4\2GGgg\4\2--//\3\2\63;\4\2"+
		"C\\c|\7\2/;C\\^^aac|\3\2bb\n\2))\61\61^^ddhhppttvv\5\2\2!))^^\n\2$$\61"+
		"\61^^ddhhppttvv\5\2\62;CHch\5\2\2!$$^^\4\2\f\f\17\17\5\2\13\f\17\17\""+
		"\"\2\u026d\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2"+
		"\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27"+
		"\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2"+
		"\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2"+
		"\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2"+
		"\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2"+
		"\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S"+
		"\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2"+
		"\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2"+
		"\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y"+
		"\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0087\3"+
		"\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u0099\3\2\2\2\2\u009b\3\2\2\2"+
		"\2\u009d\3\2\2\2\3\u009f\3\2\2\2\5\u00a1\3\2\2\2\7\u00a3\3\2\2\2\t\u00a5"+
		"\3\2\2\2\13\u00aa\3\2\2\2\r\u00ad\3\2\2\2\17\u00b0\3\2\2\2\21\u00b3\3"+
		"\2\2\2\23\u00b6\3\2\2\2\25\u00b8\3\2\2\2\27\u00ba\3\2\2\2\31\u00bd\3\2"+
		"\2\2\33\u00c0\3\2\2\2\35\u00c2\3\2\2\2\37\u00c4\3\2\2\2!\u00c6\3\2\2\2"+
		"#\u00c8\3\2\2\2%\u00ca\3\2\2\2\'\u00cc\3\2\2\2)\u00ce\3\2\2\2+\u00d0\3"+
		"\2\2\2-\u00d2\3\2\2\2/\u00d4\3\2\2\2\61\u00d6\3\2\2\2\63\u00d8\3\2\2\2"+
		"\65\u00da\3\2\2\2\67\u00dc\3\2\2\29\u00e1\3\2\2\2;\u00e7\3\2\2\2=\u00eb"+
		"\3\2\2\2?\u00ee\3\2\2\2A\u00f3\3\2\2\2C\u00f9\3\2\2\2E\u00fd\3\2\2\2G"+
		"\u0105\3\2\2\2I\u010b\3\2\2\2K\u0111\3\2\2\2M\u0115\3\2\2\2O\u011d\3\2"+
		"\2\2Q\u0124\3\2\2\2S\u0127\3\2\2\2U\u012e\3\2\2\2W\u0131\3\2\2\2Y\u0134"+
		"\3\2\2\2[\u013e\3\2\2\2]\u0143\3\2\2\2_\u0148\3\2\2\2a\u0153\3\2\2\2c"+
		"\u015c\3\2\2\2e\u0167\3\2\2\2g\u0172\3\2\2\2i\u017b\3\2\2\2k\u0186\3\2"+
		"\2\2m\u0193\3\2\2\2o\u01a2\3\2\2\2q\u01af\3\2\2\2s\u01be\3\2\2\2u\u01ca"+
		"\3\2\2\2w\u01d3\3\2\2\2y\u01d9\3\2\2\2{\u01e2\3\2\2\2}\u01e9\3\2\2\2\177"+
		"\u01f1\3\2\2\2\u0081\u01f9\3\2\2\2\u0083\u0207\3\2\2\2\u0085\u0215\3\2"+
		"\2\2\u0087\u0217\3\2\2\2\u0089\u021e\3\2\2\2\u008b\u0239\3\2\2\2\u008d"+
		"\u023b\3\2\2\2\u008f\u0240\3\2\2\2\u0091\u0242\3\2\2\2\u0093\u0247\3\2"+
		"\2\2\u0095\u024d\3\2\2\2\u0097\u024f\3\2\2\2\u0099\u0251\3\2\2\2\u009b"+
		"\u025b\3\2\2\2\u009d\u0261\3\2\2\2\u009f\u00a0\7<\2\2\u00a0\4\3\2\2\2"+
		"\u00a1\u00a2\7]\2\2\u00a2\6\3\2\2\2\u00a3\u00a4\7_\2\2\u00a4\b\3\2\2\2"+
		"\u00a5\u00a6\7p\2\2\u00a6\u00a7\7w\2\2\u00a7\u00a8\7n\2\2\u00a8\u00a9"+
		"\7n\2\2\u00a9\n\3\2\2\2\u00aa\u00ab\7~\2\2\u00ab\u00ac\7~\2\2\u00ac\f"+
		"\3\2\2\2\u00ad\u00ae\7(\2\2\u00ae\u00af\7(\2\2\u00af\16\3\2\2\2\u00b0"+
		"\u00b1\7?\2\2\u00b1\u00b2\7?\2\2\u00b2\20\3\2\2\2\u00b3\u00b4\7#\2\2\u00b4"+
		"\u00b5\7?\2\2\u00b5\22\3\2\2\2\u00b6\u00b7\7@\2\2\u00b7\24\3\2\2\2\u00b8"+
		"\u00b9\7>\2\2\u00b9\26\3\2\2\2\u00ba\u00bb\7@\2\2\u00bb\u00bc\7?\2\2\u00bc"+
		"\30\3\2\2\2\u00bd\u00be\7>\2\2\u00be\u00bf\7?\2\2\u00bf\32\3\2\2\2\u00c0"+
		"\u00c1\7-\2\2\u00c1\34\3\2\2\2\u00c2\u00c3\7/\2\2\u00c3\36\3\2\2\2\u00c4"+
		"\u00c5\7,\2\2\u00c5 \3\2\2\2\u00c6\u00c7\7\61\2\2\u00c7\"\3\2\2\2\u00c8"+
		"\u00c9\7\'\2\2\u00c9$\3\2\2\2\u00ca\u00cb\7`\2\2\u00cb&\3\2\2\2\u00cc"+
		"\u00cd\7#\2\2\u00cd(\3\2\2\2\u00ce\u00cf\7=\2\2\u00cf*\3\2\2\2\u00d0\u00d1"+
		"\7?\2\2\u00d1,\3\2\2\2\u00d2\u00d3\7*\2\2\u00d3.\3\2\2\2\u00d4\u00d5\7"+
		"+\2\2\u00d5\60\3\2\2\2\u00d6\u00d7\7}\2\2\u00d7\62\3\2\2\2\u00d8\u00d9"+
		"\7\177\2\2\u00d9\64\3\2\2\2\u00da\u00db\7.\2\2\u00db\66\3\2\2\2\u00dc"+
		"\u00dd\7v\2\2\u00dd\u00de\7t\2\2\u00de\u00df\7w\2\2\u00df\u00e0\7g\2\2"+
		"\u00e08\3\2\2\2\u00e1\u00e2\7h\2\2\u00e2\u00e3\7c\2\2\u00e3\u00e4\7n\2"+
		"\2\u00e4\u00e5\7u\2\2\u00e5\u00e6\7g\2\2\u00e6:\3\2\2\2\u00e7\u00e8\7"+
		"p\2\2\u00e8\u00e9\7k\2\2\u00e9\u00ea\7n\2\2\u00ea<\3\2\2\2\u00eb\u00ec"+
		"\7k\2\2\u00ec\u00ed\7h\2\2\u00ed>\3\2\2\2\u00ee\u00ef\7g\2\2\u00ef\u00f0"+
		"\7n\2\2\u00f0\u00f1\7u\2\2\u00f1\u00f2\7g\2\2\u00f2@\3\2\2\2\u00f3\u00f4"+
		"\7y\2\2\u00f4\u00f5\7j\2\2\u00f5\u00f6\7k\2\2\u00f6\u00f7\7n\2\2\u00f7"+
		"\u00f8\7g\2\2\u00f8B\3\2\2\2\u00f9\u00fa\7n\2\2\u00fa\u00fb\7q\2\2\u00fb"+
		"\u00fc\7i\2\2\u00fcD\3\2\2\2\u00fd\u00fe\7m\2\2\u00fe\u00ff\7w\2\2\u00ff"+
		"\u0100\7d\2\2\u0100\u0101\7g\2\2\u0101\u0102\7e\2\2\u0102\u0103\7v\2\2"+
		"\u0103\u0104\7n\2\2\u0104F\3\2\2\2\u0105\u0106\7c\2\2\u0106\u0107\7r\2"+
		"\2\u0107\u0108\7r\2\2\u0108\u0109\7n\2\2\u0109\u010a\7{\2\2\u010aH\3\2"+
		"\2\2\u010b\u010c\7r\2\2\u010c\u010d\7c\2\2\u010d\u010e\7v\2\2\u010e\u010f"+
		"\7e\2\2\u010f\u0110\7j\2\2\u0110J\3\2\2\2\u0111\u0112\7i\2\2\u0112\u0113"+
		"\7g\2\2\u0113\u0114\7v\2\2\u0114L\3\2\2\2\u0115\u0116\7t\2\2\u0116\u0117"+
		"\7g\2\2\u0117\u0118\7r\2\2\u0118\u0119\7n\2\2\u0119\u011a\7c\2\2\u011a"+
		"\u011b\7e\2\2\u011b\u011c\7g\2\2\u011cN\3\2\2\2\u011d\u011e\7f\2\2\u011e"+
		"\u011f\7g\2\2\u011f\u0120\7n\2\2\u0120\u0121\7g\2\2\u0121\u0122\7v\2\2"+
		"\u0122\u0123\7g\2\2\u0123P\3\2\2\2\u0124\u0125\7/\2\2\u0125\u0126\7p\2"+
		"\2\u0126R\3\2\2\2\u0127\u0128\7/\2\2\u0128\u0129\7/\2\2\u0129\u012a\7"+
		"v\2\2\u012a\u012b\7{\2\2\u012b\u012c\7r\2\2\u012c\u012d\7g\2\2\u012dT"+
		"\3\2\2\2\u012e\u012f\7/\2\2\u012f\u0130\7r\2\2\u0130V\3\2\2\2\u0131\u0132"+
		"\7/\2\2\u0132\u0133\7w\2\2\u0133X\3\2\2\2\u0134\u0135\7/\2\2\u0135\u0136"+
		"\7l\2\2\u0136\u0137\7u\2\2\u0137\u0138\7q\2\2\u0138\u0139\7p\2\2\u0139"+
		"\u013a\7r\2\2\u013a\u013b\7c\2\2\u013b\u013c\7v\2\2\u013c\u013d\7j\2\2"+
		"\u013dZ\3\2\2\2\u013e\u013f\7n\2\2\u013f\u0140\7q\2\2\u0140\u0141\7c\2"+
		"\2\u0141\u0142\7f\2\2\u0142\\\3\2\2\2\u0143\u0144\7g\2\2\u0144\u0145\7"+
		"z\2\2\u0145\u0146\7k\2\2\u0146\u0147\7v\2\2\u0147^\3\2\2\2\u0148\u0149"+
		"\7l\2\2\u0149\u014a\7u\2\2\u014a\u014b\7q\2\2\u014b\u014c\7p\2\2\u014c"+
		"\u014d\7U\2\2\u014d\u014e\7g\2\2\u014e\u014f\7n\2\2\u014f\u0150\7g\2\2"+
		"\u0150\u0151\7e\2\2\u0151\u0152\7v\2\2\u0152`\3\2\2\2\u0153\u0154\7l\2"+
		"\2\u0154\u0155\7u\2\2\u0155\u0156\7q\2\2\u0156\u0157\7p\2\2\u0157\u0158"+
		"\7G\2\2\u0158\u0159\7f\2\2\u0159\u015a\7k\2\2\u015a\u015b\7v\2\2\u015b"+
		"b\3\2\2\2\u015c\u015d\7l\2\2\u015d\u015e\7u\2\2\u015e\u015f\7q\2\2\u015f"+
		"\u0160\7p\2\2\u0160\u0161\7F\2\2\u0161\u0162\7g\2\2\u0162\u0163\7n\2\2"+
		"\u0163\u0164\7g\2\2\u0164\u0165\7v\2\2\u0165\u0166\7g\2\2\u0166d\3\2\2"+
		"\2\u0167\u0168\7{\2\2\u0168\u0169\7c\2\2\u0169\u016a\7o\2\2\u016a\u016b"+
		"\7n\2\2\u016b\u016c\7U\2\2\u016c\u016d\7g\2\2\u016d\u016e\7n\2\2\u016e"+
		"\u016f\7g\2\2\u016f\u0170\7e\2\2\u0170\u0171\7v\2\2\u0171f\3\2\2\2\u0172"+
		"\u0173\7{\2\2\u0173\u0174\7c\2\2\u0174\u0175\7o\2\2\u0175\u0176\7n\2\2"+
		"\u0176\u0177\7G\2\2\u0177\u0178\7f\2\2\u0178\u0179\7k\2\2\u0179\u017a"+
		"\7v\2\2\u017ah\3\2\2\2\u017b\u017c\7{\2\2\u017c\u017d\7c\2\2\u017d\u017e"+
		"\7o\2\2\u017e\u017f\7n\2\2\u017f\u0180\7F\2\2\u0180\u0181\7g\2\2\u0181"+
		"\u0182\7n\2\2\u0182\u0183\7g\2\2\u0183\u0184\7v\2\2\u0184\u0185\7g\2\2"+
		"\u0185j\3\2\2\2\u0186\u0187\7m\2\2\u0187\u0188\7w\2\2\u0188\u0189\7d\2"+
		"\2\u0189\u018a\7g\2\2\u018a\u018b\7L\2\2\u018b\u018c\7u\2\2\u018c\u018d"+
		"\7q\2\2\u018d\u018e\7p\2\2\u018e\u018f\7G\2\2\u018f\u0190\7f\2\2\u0190"+
		"\u0191\7k\2\2\u0191\u0192\7v\2\2\u0192l\3\2\2\2\u0193\u0194\7m\2\2\u0194"+
		"\u0195\7w\2\2\u0195\u0196\7d\2\2\u0196\u0197\7g\2\2\u0197\u0198\7L\2\2"+
		"\u0198\u0199\7u\2\2\u0199\u019a\7q\2\2\u019a\u019b\7p\2\2\u019b\u019c"+
		"\7F\2\2\u019c\u019d\7g\2\2\u019d\u019e\7n\2\2\u019e\u019f\7g\2\2\u019f"+
		"\u01a0\7v\2\2\u01a0\u01a1\7g\2\2\u01a1n\3\2\2\2\u01a2\u01a3\7m\2\2\u01a3"+
		"\u01a4\7w\2\2\u01a4\u01a5\7d\2\2\u01a5\u01a6\7g\2\2\u01a6\u01a7\7[\2\2"+
		"\u01a7\u01a8\7c\2\2\u01a8\u01a9\7o\2\2\u01a9\u01aa\7n\2\2\u01aa\u01ab"+
		"\7G\2\2\u01ab\u01ac\7f\2\2\u01ac\u01ad\7k\2\2\u01ad\u01ae\7v\2\2\u01ae"+
		"p\3\2\2\2\u01af\u01b0\7m\2\2\u01b0\u01b1\7w\2\2\u01b1\u01b2\7d\2\2\u01b2"+
		"\u01b3\7g\2\2\u01b3\u01b4\7[\2\2\u01b4\u01b5\7c\2\2\u01b5\u01b6\7o\2\2"+
		"\u01b6\u01b7\7n\2\2\u01b7\u01b8\7F\2\2\u01b8\u01b9\7g\2\2\u01b9\u01ba"+
		"\7n\2\2\u01ba\u01bb\7g\2\2\u01bb\u01bc\7v\2\2\u01bc\u01bd\7g\2\2\u01bd"+
		"r\3\2\2\2\u01be\u01bf\7u\2\2\u01bf\u01c0\7j\2\2\u01c0\u01c1\7g\2\2\u01c1"+
		"\u01c2\7n\2\2\u01c2\u01c3\7n\2\2\u01c3\u01c4\7U\2\2\u01c4\u01c5\7e\2\2"+
		"\u01c5\u01c6\7t\2\2\u01c6\u01c7\7k\2\2\u01c7\u01c8\7r\2\2\u01c8\u01c9"+
		"\7v\2\2\u01c9t\3\2\2\2\u01ca\u01cb\7f\2\2\u01cb\u01cc\7q\2\2\u01cc\u01cd"+
		"\7y\2\2\u01cd\u01ce\7p\2\2\u01ce\u01cf\7n\2\2\u01cf\u01d0\7q\2\2\u01d0"+
		"\u01d1\7c\2\2\u01d1\u01d2\7f\2\2\u01d2v\3\2\2\2\u01d3\u01d4\7u\2\2\u01d4"+
		"\u01d5\7n\2\2\u01d5\u01d6\7g\2\2\u01d6\u01d7\7g\2\2\u01d7\u01d8\7r\2\2"+
		"\u01d8x\3\2\2\2\u01d9\u01da\7u\2\2\u01da\u01db\7v\2\2\u01db\u01dc\7g\2"+
		"\2\u01dc\u01dd\7r\2\2\u01dd\u01de\7K\2\2\u01de\u01df\7p\2\2\u01df\u01e0"+
		"\7h\2\2\u01e0\u01e1\7q\2\2\u01e1z\3\2\2\2\u01e2\u01e3\7h\2\2\u01e3\u01e4"+
		"\7k\2\2\u01e4\u01e5\7n\2\2\u01e5\u01e6\7v\2\2\u01e6\u01e7\7g\2\2\u01e7"+
		"\u01e8\7t\2\2\u01e8|\3\2\2\2\u01e9\u01ea\7r\2\2\u01ea\u01eb\7c\2\2\u01eb"+
		"\u01ec\7v\2\2\u01ec\u01ed\7v\2\2\u01ed\u01ee\7g\2\2\u01ee\u01ef\7t\2\2"+
		"\u01ef\u01f0\7p\2\2\u01f0~\3\2\2\2\u01f1\u01f5\t\2\2\2\u01f2\u01f4\t\3"+
		"\2\2\u01f3\u01f2\3\2\2\2\u01f4\u01f7\3\2\2\2\u01f5\u01f3\3\2\2\2\u01f5"+
		"\u01f6\3\2\2\2\u01f6\u0080\3\2\2\2\u01f7\u01f5\3\2\2\2\u01f8\u01fa\7/"+
		"\2\2\u01f9\u01f8\3\2\2\2\u01f9\u01fa\3\2\2\2\u01fa\u01fb\3\2\2\2\u01fb"+
		"\u0202\5\u0085C\2\u01fc\u01fe\7\60\2\2\u01fd\u01ff\t\4\2\2\u01fe\u01fd"+
		"\3\2\2\2\u01ff\u0200\3\2\2\2\u0200\u01fe\3\2\2\2\u0200\u0201\3\2\2\2\u0201"+
		"\u0203\3\2\2\2\u0202\u01fc\3\2\2\2\u0202\u0203\3\2\2\2\u0203\u0205\3\2"+
		"\2\2\u0204\u0206\5\u0083B\2\u0205\u0204\3\2\2\2\u0205\u0206\3\2\2\2\u0206"+
		"\u0082\3\2\2\2\u0207\u0209\t\5\2\2\u0208\u020a\t\6\2\2\u0209\u0208\3\2"+
		"\2\2\u0209\u020a\3\2\2\2\u020a\u020b\3\2\2\2\u020b\u020c\5\u0085C\2\u020c"+
		"\u0084\3\2\2\2\u020d\u0216\7\62\2\2\u020e\u0212\t\7\2\2\u020f\u0211\t"+
		"\4\2\2\u0210\u020f\3\2\2\2\u0211\u0214\3\2\2\2\u0212\u0210\3\2\2\2\u0212"+
		"\u0213\3\2\2\2\u0213\u0216\3\2\2\2\u0214\u0212\3\2\2\2\u0215\u020d\3\2"+
		"\2\2\u0215\u020e\3\2\2\2\u0216\u0086\3\2\2\2\u0217\u021b\t\b\2\2\u0218"+
		"\u021a\t\t\2\2\u0219\u0218\3\2\2\2\u021a\u021d\3\2\2\2\u021b\u0219\3\2"+
		"\2\2\u021b\u021c\3\2\2\2\u021c\u0088\3\2\2\2\u021d\u021b\3\2\2\2\u021e"+
		"\u0222\7b\2\2\u021f\u0221\n\n\2\2\u0220\u021f\3\2\2\2\u0221\u0224\3\2"+
		"\2\2\u0222\u0220\3\2\2\2\u0222\u0223\3\2\2\2\u0223\u0225\3\2\2\2\u0224"+
		"\u0222\3\2\2\2\u0225\u0226\7b\2\2\u0226\u008a\3\2\2\2\u0227\u022c\7$\2"+
		"\2\u0228\u022b\5\u0091I\2\u0229\u022b\5\u0097L\2\u022a\u0228\3\2\2\2\u022a"+
		"\u0229\3\2\2\2\u022b\u022e\3\2\2\2\u022c\u022a\3\2\2\2\u022c\u022d\3\2"+
		"\2\2\u022d\u022f\3\2\2\2\u022e\u022c\3\2\2\2\u022f\u023a\7$\2\2\u0230"+
		"\u0235\7)\2\2\u0231\u0234\5\u008dG\2\u0232\u0234\5\u008fH\2\u0233\u0231"+
		"\3\2\2\2\u0233\u0232\3\2\2\2\u0234\u0237\3\2\2\2\u0235\u0233\3\2\2\2\u0235"+
		"\u0236\3\2\2\2\u0236\u0238\3\2\2\2\u0237\u0235\3\2\2\2\u0238\u023a\7)"+
		"\2\2\u0239\u0227\3\2\2\2\u0239\u0230\3\2\2\2\u023a\u008c\3\2\2\2\u023b"+
		"\u023e\7^\2\2\u023c\u023f\t\13\2\2\u023d\u023f\5\u0093J\2\u023e\u023c"+
		"\3\2\2\2\u023e\u023d\3\2\2\2\u023f\u008e\3\2\2\2\u0240\u0241\n\f\2\2\u0241"+
		"\u0090\3\2\2\2\u0242\u0245\7^\2\2\u0243\u0246\t\r\2\2\u0244\u0246\5\u0093"+
		"J\2\u0245\u0243\3\2\2\2\u0245\u0244\3\2\2\2\u0246\u0092\3\2\2\2\u0247"+
		"\u0248\7w\2\2\u0248\u0249\5\u0095K\2\u0249\u024a\5\u0095K\2\u024a\u024b"+
		"\5\u0095K\2\u024b\u024c\5\u0095K\2\u024c\u0094\3\2\2\2\u024d\u024e\t\16"+
		"\2\2\u024e\u0096\3\2\2\2\u024f\u0250\n\17\2\2\u0250\u0098\3\2\2\2\u0251"+
		"\u0255\7%\2\2\u0252\u0254\n\20\2\2\u0253\u0252\3\2\2\2\u0254\u0257\3\2"+
		"\2\2\u0255\u0253\3\2\2\2\u0255\u0256\3\2\2\2\u0256\u0258\3\2\2\2\u0257"+
		"\u0255\3\2\2\2\u0258\u0259\bM\2\2\u0259\u009a\3\2\2\2\u025a\u025c\t\21"+
		"\2\2\u025b\u025a\3\2\2\2\u025c\u025d\3\2\2\2\u025d\u025b\3\2\2\2\u025d"+
		"\u025e\3\2\2\2\u025e\u025f\3\2\2\2\u025f\u0260\bN\2\2\u0260\u009c\3\2"+
		"\2\2\u0261\u0262\13\2\2\2\u0262\u009e\3\2\2\2\26\2\u01f5\u01f9\u0200\u0202"+
		"\u0205\u0209\u0212\u0215\u021b\u0222\u022a\u022c\u0233\u0235\u0239\u023e"+
		"\u0245\u0255\u025d\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}