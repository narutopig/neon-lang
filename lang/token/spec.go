package token

import "regexp"

func r(pattern string) regexp.Regexp {
	return *regexp.MustCompile("^(" + pattern + ")")
}

// Spec defines a pattern for each token
var Spec = map[TokenType]regexp.Regexp{
	NumType:   r("number"),
	StrType:   r("string"),
	BoolType:  r("bool"),
	VoidType:  r("void"),
	True:      r("true"),
	False:     r("false"),
	Def:       r("def"),
	Return:    r("return"),
	If:        r("if"),
	Elif:      r("elif"),
	While:     r("while"),
	Equality:  r("=="),
	NotEqual:  r("!="),
	LessEq:    r("<="),
	GreaterEq: r(">="),
	Less:      r(">"),
	Greater:   r(">"),

	// long
	StrLiteral: r("\"[^\n\"]*\"|'[^\n']*'"),
	NumLiteral: r("\\d+(\\.\\d+)?"),
	Id:         r("[a-zA-Z_][a-zA-Z0-9_]*"),

	// short
	AddSub:    r("[+-]"),
	MulDivMod: r("[*/%]"),
	Comma:     r(","),
	Semi:      r(";"),
	Lp:        r("\\("),
	Rp:        r("\\)"),
	Lb:        r("\\["),
	Rb:        r("\\]"),
	Lc:        r("{"),
	Rc:        r("}"),
	Equal:     r("="),
	Bang:      r("!"),
}

// WhiteSpace is a regex for matching whitespaces
var WhiteSpace = r("[^\\S]+")
