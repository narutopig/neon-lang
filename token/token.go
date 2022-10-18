//go:generate stringer -type=TokenType -output token_string.go

package tokens

import (
	"fmt"
)

// TokenType is a enum type for Token
type TokenType byte

const (
	// identifier

	IDENTIFIER TokenType = iota // num

	// keywords

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
	LEFTCURLY    // {
	RIGHTCURLY   // }
	COMMA        // ,

	// NEWLINE      // \n

	TAB    //
	PERIOD // .
	SEMI   // ;

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
	NOTEQ     // !=
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

// NoneToken describes an empty/null token
var NoneToken = New(NONE, "")

// New returns a Token
func New(tokenType TokenType, value string) Token {
	return Token{tokenType, value}
}

func (t Token) String() string {
	if t.Value != "" {
		return fmt.Sprintf("{Type: %s, Value: %s}", t.Type, t.Value)
	}
	return fmt.Sprintf("{Type: %s}", t.Type)
}

// ContainsTT checks if a TokenType is within a TokenType[]
func ContainsTT(tt TokenType, ttl []TokenType) bool {
	for _, t := range ttl {
		if tt == t {
			return true
		}
	}

	return false
}

// Precedence calculates precedence of the TokenType
func Precedence(tt TokenType) int {
	if tt == ADD || tt == SUBTRACT {
		return 0
	} else if tt == MULTIPLY || tt == DIVIDE || tt == MODULUS {
		return 1
	}
	return -1
}
