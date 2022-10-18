package grammar

import (
	t "github.com/narutopig/neon-lang/token"
)

// Statements
var (
	ASSIGNMENT []Segment = []Segment{
		NewSegmentS(S(TYPES, T(t.IDENTIFIER), T(t.ASSIGN))),
		NewSegment(false, []Sequence{S(T(t.UNKNOWNVALUE))}),
		NewSegmentS(S(T(t.SEMI))),
	}
	DECLARATION []Segment = []Segment{
		NewSegmentS(S(TYPES, T(t.IDENTIFIER), T(t.ASSIGN))),
		NewSegment(false, []Sequence{S(T(t.UNKNOWNVALUE))}),
		NewSegmentS(S(T(t.SEMI))),
	}
	ASSIGNMENTS Sequence = S(
		TYPES, T(t.IDENTIFIER), T(t.ASSIGN),
		T(t.UNKNOWNVALUE),
		T(t.SEMI),
	)
)

// Match checks if a []TokenType matches a Sequence
func Match(tokens []t.TokenType, segment Sequence) bool {
	split := -1

	// is it a two part
	for i := 0; i < len(segment); i++ {
		if segment[i][0] == t.UNKNOWNVALUE {
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
			if !t.ContainsTT(tokens[i], segment[i]) {
				return false
			}
		}

		return true
	}

	for i := 0; i < split; i++ {
		if !t.ContainsTT(tokens[i], segment[i]) {
			return false
		}
	}

	// painful backwards loop
	for i := len(tokens) - 1; i > split; i-- {
		if !t.ContainsTT(tokens[i], segment[i]) {
			return false
		}
	}

	return true
}
