package grammar

import (
	l "github.com/narutopig/neon-lang/lexer"
)

// Sequence describes the valid tokens for each index defined in an expression
type Sequence [][]l.TokenType

// grouping of tokens
var (
	TYPES    = []l.TokenType{l.INTTYPE, l.FLOATTYPE, l.STRINGTYPE, l.BOOLEANTYPE}            // variable types
	VALUES   = []l.TokenType{l.IDENTIFIER, l.NUMVALUE, l.STRINGVALUE, l.BOOLEANVALUE}        // value
	ARITHOPS = []l.TokenType{l.ADD, l.SUBTRACT, l.MULTIPLY, l.DIVIDE, l.MODULUS}             // arithmetic operators
	COMPOPS  = []l.TokenType{l.EQUAL, l.GREATER, l.GREATEREQ, l.LESS, l.LESSEQ, l.AND, l.OR} // boolean operators
)

// grouping sequences of tokens into grammar? idk ngl
var (
	EXPRESSION Sequence = [][]l.TokenType{
		append(T(l.IDENTIFIER), VALUES...),
		append(ARITHOPS, COMPOPS...),
		append(T(l.IDENTIFIER), VALUES...),
	}
	VALUE       Sequence = S(VALUES)
	DECLARATION          = join(S(TYPES, T(l.IDENTIFIER), T(l.ASSIGN)), EXPRESSION, S(T(l.SEMI)))
	ASSIGNMENT           = join(S(T(l.IDENTIFIER), T(l.ASSIGN)), EXPRESSION, S(T(l.SEMI)))
)

// T returns l.TokenType's in an array
func T(types ...l.TokenType) []l.TokenType {
	return types
}

// S returns []l.TokenType's in an array (Sequence)
func S(t ...[]l.TokenType) Sequence {
	first := [][]l.TokenType{}

	for i := 0; i < len(t); i++ {
		first = append(first, t[i])
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
func IsSequence(tokens []l.TokenType, sequence Sequence) bool {
	length := len(tokens)
	if length != len(sequence) {
		return false
	}

	for i := 0; i < length; i++ {
		if !l.ContainsTT(tokens[i], sequence[i]) {
			return false
		}
	}

	return true
}
