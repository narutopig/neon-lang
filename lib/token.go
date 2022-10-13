package lib

import (
	"fmt"
	"strconv"
)

// TokenType is a enum type for Token
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

func (t Token) String() string {
	return fmt.Sprintf("{Type: %s, Value: %s}", t.Type, t.Value)
}

// NewToken returns a Token
func NewToken(tokenType TokenType, value string) Token {
	return Token{tokenType, value}
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[STDFUNCTION-0]
	_ = x[FUNCTION-1]
	_ = x[VARIABLE-2]
	_ = x[INTTYPE-3]
	_ = x[FLOATTYPE-4]
	_ = x[STRINGTYPE-5]
	_ = x[BOOLEANTYPE-6]
	_ = x[STRINGVALUE-7]
	_ = x[NUMVALUE-8]
	_ = x[BOOLEANVALUE-9]
	_ = x[LEFTPAREN-10]
	_ = x[RIGHTPAREN-11]
	_ = x[LEFTBRACKET-12]
	_ = x[RIGHTBRACKET-13]
	_ = x[COMMA-14]
	_ = x[NEWLINE-15]
	_ = x[TAB-16]
	_ = x[ADD-17]
	_ = x[SUBTRACT-18]
	_ = x[MULTIPLY-19]
	_ = x[DIVIDE-20]
	_ = x[MODULUS-21]
	_ = x[ASSIGN-22]
	_ = x[EQUAL-23]
	_ = x[GREATER-24]
	_ = x[LESS-25]
	_ = x[GREATEREQ-26]
	_ = x[LESSEQ-27]
	_ = x[NOT-28]
	_ = x[AND-29]
	_ = x[OR-30]
	_ = x[UNKNOWNVALUE-31]
	_ = x[NONE-32]
}

const ttn = "STDFUNCTIONFUNCTIONVARIABLEINTTYPEFLOATTYPESTRINGTYPEBOOLEANTYPESTRINGVALUENUMVALUEBOOLEANVALUELEFTPARENRIGHTPARENLEFTBRACKETRIGHTBRACKETCOMMANEWLINETABADDSUBTRACTMULTIPLYDIVIDEMODULUSASSIGNEQUALGREATERLESSGREATEREQLESSEQNOTANDORUNKNOWNVALUENONE"

var tti = [...]uint8{0, 11, 19, 27, 34, 43, 53, 64, 75, 83, 95, 104, 114, 125, 137, 142, 149, 152, 155, 163, 171, 177, 184, 190, 195, 202, 206, 215, 221, 224, 227, 229, 241, 245}

func (i TokenType) String() string {
	if i >= TokenType(len(tti)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return ttn[tti[i]:tti[i+1]]
}
