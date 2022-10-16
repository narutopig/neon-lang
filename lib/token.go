//go:generate stringer -type=TokenType -output token_string.go

package lib

import (
	"fmt"
)

// TokenType is a enum type for Token
type TokenType byte

const (
	// variables

	IDENTIFIER  TokenType = iota // num
	INTTYPE                      // int
	FLOATTYPE                    // float
	STRINGTYPE                   // string
	BOOLEANTYPE                  // boolean

	// values

	STRINGVALUE  // "some string"
	NUMVALUE     // 0
	BOOLEANVALUE // true / false

	// structure symbols

	LEFTPAREN    // (
	RIGHTPAREN   // )
	LEFTBRACKET  // [
	RIGHTBRACKET // ]
	LEFTCURLY    // {
	RIGHTCURLY   // }
	COMMA        // ,
	NEWLINE      // \n
	TAB          //
	PERIOD       // .

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

func (t Token) String() string {
	if t.Value != "" {
		return fmt.Sprintf("{Type: %s, Value: %s}", t.Type, t.Value)
	}
	return fmt.Sprintf("{Type: %s}", t.Type)
}

// NewToken returns a Token
func NewToken(tokenType TokenType, value string) Token {
	return Token{tokenType, value}
}
