package grammar

import (
	"github.com/narutopig/neon-lang/lexer"
	l "github.com/narutopig/neon-lang/lexer"
)

// Statements
var (
	ASSIGNMENT []Segment = []Segment{
		NewSegmentS(S(TYPES, T(l.IDENTIFIER), T(l.ASSIGN))),
		NewSegment(false, []Sequence{S(T(l.UNKNOWNVALUE))}),
		NewSegmentS(S(T(l.SEMI))),
	}
	DECLARATION []Segment = []Segment{
		NewSegmentS(S(TYPES, T(l.IDENTIFIER), T(l.ASSIGN))),
		NewSegment(false, []Sequence{S(T(l.UNKNOWNVALUE))}),
		NewSegmentS(S(T(l.SEMI))),
	}
	ASSIGNMENTS Sequence = S(
		TYPES, T(l.IDENTIFIER), T(l.ASSIGN),
		T(l.UNKNOWNVALUE),
		T(l.SEMI),
	)
)

// Match checks if a []TokenType matches a Sequence
func Match(tokens []l.TokenType, segment Sequence) bool {
	split := -1

	// is it a two part
	for i := 0; i < len(segment); i++ {
		if segment[i][0] == l.UNKNOWNVALUE {
			split = i
			break
		}
	}

	if split == -1 {
		//  one part thing
		l := len(tokens)

		if l != len(segment) {
			return false
		}

		for i := 0; i < l; i++ {
			if !lexer.ContainsTT(tokens[i], segment[i]) {
				return false
			}
		}

		return true
	}

	for i := 0; i < split; i++ {
		if !lexer.ContainsTT(tokens[i], segment[i]) {
			return false
		}
	}

	// painful backwards loop
	for i := len(tokens) - 1; i > split; i-- {
		if !lexer.ContainsTT(tokens[i], segment[i]) {
			return false
		}
	}

	return true
}
