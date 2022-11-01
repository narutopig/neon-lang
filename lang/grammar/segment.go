package grammar

import (
	"fmt"
	"reflect"

	t "github.com/narutopig/neon-lang/lang/token"
)

// Segment represents parts of a statement
type Segment struct {
	Static    bool // are there multiple valid sequences that can be placed here (length of Sequences is 1)
	Sequences []Sequence
}

// NewSegment is the default constructor
func NewSegment(static bool, sequences []Sequence) Segment {
	return Segment{Static: static, Sequences: sequences}
}

// NewSegmentS returns a static segment
func NewSegmentS(sequence Sequence) Segment {
	return Segment{true, []Sequence{sequence}}
}

// SegmentSeq returns stuff as the thing
func SegmentSeq(args ...interface{}) []Segment {
	stuffs := []Segment{}
	currTokenTypeList := [][]t.TokenType{}

	for _, arg := range args {
		switch reflect.TypeOf(arg).String() {
		case "lexer.TokenType":
			currTokenTypeList = append(currTokenTypeList, T(arg.(t.TokenType)))
		case "[]lexer.TokenType":
			currTokenTypeList = append(currTokenTypeList, arg.([]t.TokenType))
		case "grammar.Sequence":
			currTokenTypeList = append(currTokenTypeList, arg.(Sequence)...)
		case "grammar.Segment":
			stuffs = append(stuffs, NewSegmentS(currTokenTypeList))
			stuffs = append(stuffs, arg.(Segment))
			currTokenTypeList = [][]t.TokenType{}
		}
	}

	if len(currTokenTypeList) > 0 {
		stuffs = append(stuffs, NewSegmentS(currTokenTypeList))
	}

	return stuffs
}

// Validate checks if a series of tokens fits a []Segment
func Validate(tokens []t.TokenType, segseq []Segment) {
	possible := []Sequence{}
	recurse([][]t.TokenType{}, 0, &possible, segseq)

	for _, val := range possible {
		fmt.Println(IsSequence(tokens, val))
	}
}

// Maps out all possible Sequences
func recurse(currArr Sequence, index int, sequences *[]Sequence, segments []Segment) {
	if index == len(segments) {
		t := *sequences
		*sequences = append(t, currArr)
		return
	}

	for _, val := range segments[index].Sequences {
		recurse(append(currArr, val...), index+1, sequences, segments)
	}
}
