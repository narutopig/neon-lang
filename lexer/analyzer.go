package lexer

import (
	"strings"

	t "github.com/narutopig/neon-lang/token"
)

func typeToken(str string) t.Token {
	switch str {
	case "int":
		return t.New(t.INTTYPE, "")
	case "string":
		return t.New(t.STRINGTYPE, "")
	case "float":
		return t.New(t.FLOATTYPE, "")
	case "bool":
		return t.New(t.BOOLEANTYPE, "")
	}
	return t.New(t.NONE, "")
}

func identifierMagic(token t.Token) t.Token {
	if token.Value == "true" || token.Value == "false" {
		return t.New(t.BOOLEANVALUE, token.Value)
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
		return t.New(t.LEFTPAREN, "")
	case ")":
		return t.New(t.RIGHTPAREN, "")
	case "[":
		return t.New(t.LEFTBRACKET, "")
	case "]":
		return t.New(t.RIGHTBRACKET, "")
	case "{":
		return t.New(t.LEFTCURLY, "")
	case "}":
		return t.New(t.RIGHTCURLY, "")
	case ",":
		return t.New(t.COMMA, "")
	case ".":
		return t.New(t.PERIOD, "")
	case ";":
		return t.New(t.SEMI, "")
	case "+":
		return t.New(t.ADD, "")
	case "-":
		return t.New(t.SUBTRACT, "")
	case "*":
		return t.New(t.MULTIPLY, "")
	case "/":
		return t.New(t.DIVIDE, "")
	case "%":
		return t.New(t.MODULUS, "")
	case "=":
		return t.New(t.ASSIGN, "")
	case "<":
		return t.New(t.LESS, "")
	case ">":
		return t.New(t.GREATER, "")
	case "!":
		return t.New(t.NOT, "")
	case "&":
		return t.New(t.AND, "")
	case "|":
		return t.New(t.OR, "")
	}
	return t.New(t.NONE, "")
}
