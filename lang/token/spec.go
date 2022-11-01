package token

import "regexp"

func r(pattern string) *regexp.Regexp {
	return regexp.MustCompile("^(" + pattern + ")")
}

// Spec defines a pattern for each token
var Spec = map[*regexp.Regexp]TokenType{
	r("number"): NumType,
	r("string"): StrType,
	r("bool"):   BoolType,
	r("void"):   VoidType,
	r("true"):   True,
	r("false"):  False,
	r("def"):    Def,
	r("return"): Return,
	r("if"):     If,
	r("elif"):   Elif,
	r("while"):  While,
	r("=="):     Equality,
	r("!="):     NotEqual,
	r("<="):     LessEq,
	r(">="):     GreaterEq,
	r(">"):      Less,
	r("<"):      Greater,

	// long
	r("\"[^\n\"]*\"|'[^\n']*'"): StrLiteral,
	r("\\d+(\\.\\d+)?"):         NumLiteral,
	r("[a-zA-Z_][a-zA-Z0-9_]*"): Id,

	// short
	r("[+-]"):  AddSub,
	r("[*/%]"): MulDivMod,
	r(","):     Comma,
	r(";"):     Semi,
	r("\\("):   Lp,
	r("\\)"):   Rp,
	r("\\["):   Lb,
	r("\\]"):   Rb,
	r("{"):     Lc,
	r("}"):     Rc,
	r("="):     Equal,
	r("!"):     Bang,
}

// WhiteSpace is a regex for matching whitespaces
var WhiteSpace = r("[^\\S]+")
