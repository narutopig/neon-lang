package grammar

import (
	l "github.com/narutopig/neon-lang/lexer"
)

// Statements
var (
	ASSIGNMENT []Segment = []Segment{
		NewSegmentS(S(TYPES, T(l.IDENTIFIER), T(l.ASSIGN))),
		NewSegment(false, []Sequence{S(VALUES), S(append(T(l.IDENTIFIER), VALUES...), append(ARITHOPS, COMPOPS...), append(T(l.IDENTIFIER), VALUES...))}),
		NewSegmentS(S(T(l.SEMI))),
	}
	DECLARATION []Segment = []Segment{
		NewSegmentS(S(TYPES, T(l.IDENTIFIER), T(l.ASSIGN))),
		NewSegment(false, []Sequence{S(VALUES), S(append(T(l.IDENTIFIER), VALUES...), append(ARITHOPS, COMPOPS...), append(T(l.IDENTIFIER), VALUES...))}),
		NewSegmentS(S(T(l.SEMI))),
	}
)
