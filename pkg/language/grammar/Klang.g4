grammar Klang;
 
parse
 : block EOF
 ;
 
block
 : stat*
 ;
 
stat
 : assignment
 | json_edit_fn
 | json_delete_fn
 | yaml_edit_fn
 | yaml_delete_fn
 | kube_json_delete_fn
 | kube_json_edit_fn
 | kube_yaml_delete_fn
 | kube_yaml_edit_fn
 | if_stat
 | while_stat
 | sleep_fn
 | exit_fn
 | log
 | OTHER {fmt.Println("unknown char: " + $OTHER.text);}
 ;

assignment
 : ID ASSIGN expr SCOL
 | ID ASSIGN load_fn SCOL
 ;

shell_script
 : SHELLSCRIPT string_or_id
 ;

json_edit_fn
 : JSONEDIT OPAR ID COMMA string_or_id COMMA expr CPAR SCOL
 ;

json_delete_fn
 : JSONDELETE OPAR ID COMMA string_or_id CPAR SCOL
 ;

yaml_edit_fn
 : YAMLEDIT OPAR ID COMMA string_or_id COMMA expr (COMMA NUMBER)? CPAR SCOL
 ;

yaml_delete_fn
 : YAMLDELETE OPAR ID COMMA string_or_id (COMMA NUMBER)? CPAR SCOL
 ;

kube_json_edit_fn
 : KUBEJSONEDIT OPAR ID COMMA string_or_id COMMA expr (COMMA string_or_id)? (COMMA asObject)? CPAR SCOL
 ;

kube_json_delete_fn
 : KUBEJSONDELETE OPAR ID COMMA filter COMMA pattern CPAR SCOL
 | KUBEJSONDELETE OPAR ID COMMA pattern COMMA filter CPAR SCOL
 | KUBEJSONDELETE OPAR ID COMMA filter CPAR SCOL
 | KUBEJSONDELETE OPAR ID COMMA pattern CPAR SCOL
 ;

kube_yaml_edit_fn
 : KUBEYAMLEDIT OPAR ID COMMA string_or_id COMMA expr (COMMA string_or_id)? (COMMA asObject)? CPAR SCOL
 ;

kube_yaml_delete_fn
 : KUBEYAMLDELETE OPAR ID COMMA filter  COMMA pattern CPAR SCOL
 | KUBEYAMLDELETE OPAR ID COMMA pattern COMMA filter CPAR SCOL
 | KUBEYAMLDELETE OPAR ID COMMA filter CPAR SCOL
 | KUBEYAMLDELETE OPAR ID COMMA pattern CPAR SCOL
 ;

sleep_fn
 : SLEEP NUMBER SCOL
 ;

exit_fn
 : EXIT NUMBER SCOL
 ;

if_stat
 : IF condition_block (ELSE IF condition_block)* (ELSE stat_block)?
 ;
 
condition_block
 : expr stat_block
 ;
 
stat_block
 : OBRACE block CBRACE
 | stat
 ;
 
while_stat
 : WHILE expr stat_block
 ;
 
log
 : LOG expr SCOL
 ;

kubectl_command
 : KUBECTL APPLY (NAMESPACE ns | string_or_id | UPDATELOAD kubernetes_object_config)+ #applyKubectlCommand
 | KUBECTL PATCH (NAMESPACE ns | resource | PATCHTYPE patch_type | PATCHLOAD string_or_id)+ #patchKubectlCommand
 | KUBECTL GET (NAMESPACE ns | resource)+ #getKubectlCommand
 | KUBECTL DELETE (NAMESPACE ns | resource)+ #deleteKubectlCommand
 ;

download_fn
 : DOWNLOAD OPAR string_or_id (COMMA string_or_id)? CPAR
 ;

json_select_fn
 : JSONSELECT OPAR ID COMMA string_or_id CPAR
 ;

yaml_select_fn
 : YAMLSELECT OPAR ID COMMA string_or_id (COMMA NUMBER)? CPAR
 ;

load_fn
 : LOAD OPAR string_or_id (COMMA STRING)? CPAR
 ;

stepInfo
 : STEPINFO STRING SCOL
 | STEPINFO RAW_STRING_LIT SCOL
 ;

ns
 : string_or_id
 | PATH
 ;

asObject
  : '"asObject"'
  ;

patch_type
 : PATH
 | string_or_id
 ;

string_or_id
 : ID
 | STRING
 | RAW_STRING_LIT
 ;

resource
 : PATH
 | string_or_id
 ;

kubernetes_object_config
 : string_or_id
 ;

filter
 : FILTER ASSIGN string_or_id
 ;

pattern
 : PATTERN ASSIGN string_or_id
 ;

expr
 :<assoc=right> expr POW expr           #powExpr
 | MINUS expr                           #unaryMinusExpr
 | NOT expr                             #notExpr
 | expr op=(MULT | DIV | MOD) expr      #multiplicationExpr
 | expr op=(PLUS | MINUS) expr          #additiveExpr
 | expr op=(LTEQ | GTEQ | LT | GT) expr #relationalExpr
 | expr op=(EQ | NEQ) expr              #equalityExpr
 | expr AND expr                        #andExpr
 | expr OR expr                         #orExpr
 | kubectl_command                      #kubectlExpr
 | json_select_fn                       #jsonSelectFn
 | yaml_select_fn                       #yamlSelectFn
 | shell_script                         #shellScript
 | download_fn                          #downloadFn
 | atom                                 #atomExpr
 ;
 
atom
 : OPAR expr CPAR #parExpr
 | NUMBER         #numberAtom
 | (TRUE | FALSE) #booleanAtom
 | RAW_STRING_LIT #rawStringAtom
 | ID             #idAtom
 | STRING         #stringAtom
 | json           #jsonAtom
 | NIL            #nilAtom
 ;

json
   : value
   ;

obj
   : '{' pair (',' pair)* '}'
   | '{' '}'
   ;

pair
   : STRING ':' value
   ;

arr
   : '[' value (',' value)* ']'
   | '[' ']'
   ;

value
   : STRING
   | NUMBER
   | obj
   | arr
   | 'true'
   | 'false'
   | 'null'
   ;
 
OR : '||';
AND : '&&';
EQ : '==';
NEQ : '!=';
GT : '>';
LT : '<';
GTEQ : '>=';
LTEQ : '<=';
PLUS : '+';
MINUS : '-';
MULT : '*';
DIV : '/';
MOD : '%';
POW : '^';
NOT : '!';
 
SCOL : ';';
ASSIGN : '=';
OPAR : '(';
CPAR : ')';
OBRACE : '{';
CBRACE : '}';
COMMA : ',';
 
TRUE : 'true';
FALSE : 'false';
NIL : 'nil';
IF : 'if';
ELSE : 'else';
WHILE : 'while';
LOG : 'log';
KUBECTL : 'kubectl';
APPLY : 'apply';
PATCH : 'patch';
GET : 'get';
REPLACE : 'replace';
DELETE : 'delete';
NAMESPACE : '-n';
PATCHTYPE : '--type';
PATCHLOAD: '-p';
UPDATELOAD: '-u';
JSONPATH : '-jsonpath';
LOAD : 'load';
EXIT : 'exit';
JSONSELECT : 'jsonSelect';
JSONEDIT : 'jsonEdit';
JSONDELETE: 'jsonDelete';
YAMLSELECT : 'yamlSelect';
YAMLEDIT : 'yamlEdit';
YAMLDELETE: 'yamlDelete';
KUBEJSONEDIT : 'kubeJsonEdit';
KUBEJSONDELETE: 'kubeJsonDelete';
KUBEYAMLEDIT : 'kubeYamlEdit';
KUBEYAMLDELETE: 'kubeYamlDelete';
SHELLSCRIPT : 'shellScript';
DOWNLOAD: 'download';
SLEEP: 'sleep';
STEPINFO: 'stepInfo';
FILTER: 'filter';
PATTERN: 'pattern';

ID
 : [a-zA-Z_] [a-zA-Z_0-9]*
 ;

NUMBER
 : '-'? INT ('.' [0-9] +)? EXP?
 ;

fragment EXP
 : [Ee] [+\-]? INT
 ;

fragment INT
 : '0' | [1-9] [0-9]*
 ;

PATH
 : [a-zA-Z] [a-zA-Z_0-9/\\.-]*
 ;

RAW_STRING_LIT
 : '`' ~'`'*                      '`'
 ;

STRING
 : '"' (ESC | SAFECODEPOINT)* '"'
 | '\'' (ESCQUOTE | SAFECODEPOINTQUOTE)* '\''
 ;

fragment ESCQUOTE
   : '\\' (['\\/bfnrt] | UNICODE)
   ;
fragment SAFECODEPOINTQUOTE
: ~ ['\\\u0000-\u001F]
;
fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;
fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;
fragment HEX
 : [0-9a-fA-F]
 ;
fragment SAFECODEPOINT
 : ~ ["\\\u0000-\u001F]
 ;

COMMENT
 : '#' ~[\r\n]* -> skip
 ;
SPACE
 : [ \t\r\n] + -> skip
 ;
OTHER
 : . 
 ;