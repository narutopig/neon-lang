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

if  : IF '(' expr ')' '{' stat* '}';
while  : WHILE '(' expr ')' '{' stat* '}';
func : DEF type ID '(' (funcarg (COMMA funcarg)*)? ')' '{' stat* '}';
funccall : ID '(' (expr (COMMA expr)*)? ')';

type    : NUMTYPE | STRTYPE | BOOLTYPE;
arithop : ADD_SUB | MDM;
compop  : EQUALITY | LESS | GREATER | LESSEQ | GREATEREQ;
op      : arithop | compop;

funcarg : type ID;
decl : type ID EQUAL expr;
assign: ID EQUAL expr;
return: RETURN expr;

expr : expr op expr
     | funccall
     | INT
     | BOOL
     | STRING
     | ID
     ;

// place all special identifiers here
NUMTYPE  : 'number';
STRTYPE  : 'string';
BOOLTYPE : 'bool';
DEF      : 'def';
RETURN   : 'return';
IF       : 'if';
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