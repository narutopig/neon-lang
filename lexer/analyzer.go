package lexer

import "strings"

func typeToken(str string) Token {
	switch str {
	case "int":
		return NewToken(INTTYPE, "")
	case "string":
		return NewToken(STRINGTYPE, "")
	case "float":
		return NewToken(FLOATTYPE, "")
	case "bool":
		return NewToken(BOOLEANTYPE, "")
	}
	return NewToken(NONE, "")
}

func identifierMagic(token Token) Token {
	if token.Value == "true" || token.Value == "false" {
		return NewToken(BOOLEANVALUE, token.Value)
	}
	tt := typeToken(token.Value)

	if tt.Type != NONE {
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

func singleCharToken(str string) Token {
	switch str {
	case "(":
		return NewToken(LEFTPAREN, "")
	case ")":
		return NewToken(RIGHTPAREN, "")
	case "[":
		return NewToken(LEFTBRACKET, "")
	case "]":
		return NewToken(RIGHTBRACKET, "")
	case "{":
		return NewToken(LEFTCURLY, "")
	case "}":
		return NewToken(RIGHTCURLY, "")
	case ",":
		return NewToken(COMMA, "")
	case ".":
		return NewToken(PERIOD, "")
	case ";":
		return NewToken(SEMI, "")
	case "+":
		return NewToken(ADD, "")
	case "-":
		return NewToken(SUBTRACT, "")
	case "*":
		return NewToken(MULTIPLY, "")
	case "/":
		return NewToken(DIVIDE, "")
	case "%":
		return NewToken(MODULUS, "")
	case "=":
		return NewToken(ASSIGN, "")
	case "<":
		return NewToken(LESS, "")
	case ">":
		return NewToken(GREATER, "")
	case "!":
		return NewToken(NOT, "")
	case "&":
		return NewToken(AND, "")
	case "|":
		return NewToken(OR, "")
	}
	return NewToken(NONE, "")
}
