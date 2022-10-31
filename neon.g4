grammar neon;

options {tokenVocab=NeonLexer;}

program
    : ((stat | func))* EOF
    ;

stat : 
    (
    decl
    | assign
    )
    SEMI;

func : DEF ID '(' funcarg* ')' '{' stat* '}';

type    : NUMTYPE | STRTYPE | BOOLTYPE;
arithop : ADD_SUB | MUL_DIV;
compop  : LESS | GREATER | LESSEQ | GREATEREQ;
op      : arithop | compop;

funcarg : type ID;
decl : type ID EQUAL expr;
assign: ID EQUAL expr;

expr : expr op expr
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

// compound operators e.g. !=
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
MUL_DIV: '*' | '/';
COMMA  : ',';
Lb     : '(';
Rb     : ')';
Lc     : '{';
Rc     : '}';
EQUAL  : '=';
LESS   : '<';
GREATER: '>';