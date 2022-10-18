package lexer

import (
	"strings"

	t "github.com/narutopig/neon-lang/token"
)

func typeToken(str string) t.Token {
	switch str {
	case "int":
		return t.NewToken(t.INTTYPE, "")
	case "string":
		return t.NewToken(t.STRINGTYPE, "")
	case "float":
		return t.NewToken(t.FLOATTYPE, "")
	case "bool":
		return t.NewToken(t.BOOLEANTYPE, "")
	}
	return t.NewToken(t.NONE, "")
}

func identifierMagic(token t.Token) t.Token {
	if token.Value == "true" || token.Value == "false" {
		return t.NewToken(t.BOOLEANVALUE, token.Value)
	}
	tt := typeToken(token.Value)

	if tt.Type != t.NONE {
		return tt
	}

	return token
}

func isDigit(str string) bool {
	return strings.Contains("123456780", str)
}

func isAlpha(str string) bool {
	return strings.Contains("qwertyuiopasdfghjklzxcvbnm", strings.ToLower(str))
}

func isAlphaNumeric(str string) bool {
	s := strings.ToLower(str)

	return isDigit(s) || isAlpha(s)
}

func singleCharToken(str string) t.Token {
	switch str {
	case "(":
		return t.NewToken(t.LEFTPAREN, "")
	case ")":
		return t.NewToken(t.RIGHTPAREN, "")
	case "[":
		return t.NewToken(t.LEFTBRACKET, "")
	case "]":
		return t.NewToken(t.RIGHTBRACKET, "")
	case "{":
		return t.NewToken(t.LEFTCURLY, "")
	case "}":
		return t.NewToken(t.RIGHTCURLY, "")
	case ",":
		return t.NewToken(t.COMMA, "")
	case ".":
		return t.NewToken(t.PERIOD, "")
	case ";":
		return t.NewToken(t.SEMI, "")
	case "+":
		return t.NewToken(t.ADD, "")
	case "-":
		return t.NewToken(t.SUBTRACT, "")
	case "*":
		return t.NewToken(t.MULTIPLY, "")
	case "/":
		return t.NewToken(t.DIVIDE, "")
	case "%":
		return t.NewToken(t.MODULUS, "")
	case "=":
		return t.NewToken(t.ASSIGN, "")
	case "<":
		return t.NewToken(t.LESS, "")
	case ">":
		return t.NewToken(t.GREATER, "")
	case "!":
		return t.NewToken(t.NOT, "")
	case "&":
		return t.NewToken(t.AND, "")
	case "|":
		return t.NewToken(t.OR, "")
	}
	return t.NewToken(t.NONE, "")
}
