package lib

import (
	"fmt"
)

// TokenType is a enum type for Token
//
// stringer -type=TokenType
//
//go:generate stringer -type=TokenType
type TokenType byte

const (
	// functions

	STDFUNCTION TokenType = iota // print()
	FUNCTION                     // count()

	// variables

	VARIABLE    // num
	INTTYPE     // int
	FLOATTYPE   // float
	STRINGTYPE  // string
	BOOLEANTYPE // boolean

	// values

	STRINGVALUE  // "some string"
	NUMVALUE     // 0
	BOOLEANVALUE // true / false

	// structure symbols

	LEFTPAREN    // (
	RIGHTPAREN   // )
	LEFTBRACKET  // [
	RIGHTBRACKET // ]
	COMMA        // ,
	NEWLINE      // \n
	TAB          // \t

	// math

	ADD      // +
	SUBTRACT // -
	MULTIPLY // *
	DIVIDE   // /
	MODULUS  // %

	// equality and comparison

	ASSIGN    // =
	EQUAL     // ==
	GREATER   // >
	LESS      // <
	GREATEREQ // >=
	LESSEQ    // <=
	NOT       // !
	AND       // &
	OR        // |

	// special values

	UNKNOWNVALUE // used when token type is unkown
	NONE
)

// Token encapsulates a TokenType and a value
type Token struct {
	Type  TokenType
	Value string
}

func (t TokenType) String() string {
	tokenTypes := []string{
		"STDFUNCTION",
		"FUNCTION",
		"VARIABLE",
		"INTTYPE",
		"FLOATTYPE",
		"STRINGTYPE",
		"BOOLEANTYPE",
		"STRINGVALUE",
		"NUMVALUE",
		"BOOLEANVALUE",
		"LEFTPAREN",
		"RIGHTPAREN",
		"LEFTBRACKET",
		"RIGHTBRACKET",
		"COMMA",
		"NEWLINE",
		"TAB",
		"ADD",
		"SUBTRACT",
		"MULTIPLY",
		"DIVIDE",
		"MODULUS",
		"ASSIGN",
		"EQUAL",
		"GREATER",
		"LESS",
		"GREATEREQ",
		"LESSEQ",
		"NOT",
		"AND",
		"OR",
		"UNKNOWNVALUE",
		"NONE",
	}

	return tokenTypes[t]
}

func (t Token) String() string {
	return fmt.Sprintf("{Type: %s, Value: %s}", t.Type, t.Value)
}

// NewToken returns a Token
func NewToken(tokenType TokenType, value string) Token {
	return Token{tokenType, value}
}
