grammar neon;

options {tokenVocab=NeonLexer;}

program
    : ((stat | func))* EOF
    ;

// complicated stuff
stat : 
    // function calls, assign, aka no curlies
    (
        (
            decl
            | assign
            | funccall
            | return
        )
        SEMI
    )
    |
    // stuff with curlies
    (
        if
        | while
    )
    ;

if  : IF '(' expr ')' '{' stat* '}' elif* else?;
elif: ELIF '(' expr ')' '{' stat* '}';
else: ELSE '{' stat* '}';
while  : WHILE '(' expr ')' '{' stat* '}';
func : DEF (type | VOID) ID ('(' (funcarg (COMMA funcarg)*)? ')')? '{' stat* '}';
funccall : ID '(' (expr (COMMA expr)*)? ')';

funcarg : type ID;
decl : type ID EQUAL expr;
assign: ID EQUAL expr;
return: RETURN expr;

expr : expr op expr
     | BANG expr
     | funccall
     | INT
     | BOOL
     | STRING
     | ID
     ;

// grouping
type    : NUMTYPE | STRTYPE | BOOLTYPE;
arithop : MDM | ADD_SUB; // order for implied precedence
compop  : EQUALITY | LESS | GREATER | LESSEQ | GREATEREQ;
op      : arithop | compop;
bool    : TRUE | FALSE;
// place all special identifiers here
NUMTYPE  : 'number';
STRTYPE  : 'string';
BOOLTYPE : 'bool';
TRUE     : 'true';
FALSE    : 'false';
VOID     : 'void';
DEF      : 'def';
RETURN   : 'return';
IF       : 'if';
ELIF     : 'elif';
ELSE     : 'else';
WHILE    : 'while';

// compound operators e.g. !=
EQUALITY: '==';
NOTEQUAL: '!=';
SEMI   : ';';
LESSEQ : '<=';
GREATEREQ : '>=';

// regular stuff
STRING : '"' (' '..'~')* '"';
ID     : ('a'..'z'|'A'..'Z')+;
INT    : '0'..'9'+;
WS     : [ \t\n\r]+ -> skip ;
ADD_SUB: '+' | '-';
MDM: '*' | '/' | '%'; // mul, div, mod
COMMA  : ',';
Lb     : '(';
Rb     : ')';
Lc     : '{';
Rc     : '}';
EQUAL  : '=';
LESS   : '<';
GREATER: '>';
BANG   : '!';