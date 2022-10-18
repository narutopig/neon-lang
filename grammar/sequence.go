package grammar

import (
	t "github.com/narutopig/neon-lang/token"
)

// Sequence describes the valid tokens for each index defined in an expression
type Sequence [][]t.TokenType

// Categories of tokens
var (
	TYPES    = []t.TokenType{t.INTTYPE, t.FLOATTYPE, t.STRINGTYPE, t.BOOLEANTYPE}            // variable types
	VALUES   = []t.TokenType{t.IDENTIFIER, t.NUMVALUE, t.STRINGVALUE, t.BOOLEANVALUE}        // value
	ARITHOPS = []t.TokenType{t.ADD, t.SUBTRACT, t.MULTIPLY, t.DIVIDE, t.MODULUS}             // arithmetic operators
	COMPOPS  = []t.TokenType{t.EQUAL, t.GREATER, t.GREATEREQ, t.LESS, t.LESSEQ, t.AND, t.OR} // boolean operators
)

// T returns t.TokenType's in an array
func T(types ...t.TokenType) []t.TokenType {
	return types
}

// S returns []t.TokenType's in an array (Sequence)
func S(tokens ...[]t.TokenType) Sequence {
	first := [][]t.TokenType{}

	for i := 0; i < len(tokens); i++ {
		first = append(first, tokens[i])
	}

	return first
}

func join(sequences ...Sequence) Sequence {
	first := sequences[0]

	for i := 1; i < len(sequences); i++ {
		first = append(first, sequences[i]...)
	}

	return first
}

// IsSequence finds if the array is a Sequence
func IsSequence(tokens []t.TokenType, sequence Sequence) bool {
	length := len(tokens)
	if length != len(sequence) {
		return false
	}

	for i := 0; i < length; i++ {
		if !t.ContainsTT(tokens[i], sequence[i]) {
			return false
		}
	}

	return true
}
