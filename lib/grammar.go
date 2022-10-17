package lib

type Sequence [][]TokenType

// grouping of tokens
var (
	TYPES    = []TokenType{INTTYPE, FLOATTYPE, STRINGTYPE, BOOLEANTYPE}      // variable types
	VALUES   = []TokenType{IDENTIFIER, NUMVALUE, STRINGVALUE, BOOLEANVALUE}  // value
	ARITHOPS = []TokenType{ADD, SUBTRACT, MULTIPLY, DIVIDE, MODULUS}         // arithmetic operators
	COMPOPS  = []TokenType{EQUAL, GREATER, GREATEREQ, LESS, LESSEQ, AND, OR} // boolean operators
)

// grouping sequences of tokens into grammar? idk ngl
var (
	EXPRESSION Sequence = [][]TokenType{
		append(T(IDENTIFIER), VALUES...),
		append(ARITHOPS, COMPOPS...),
		append(T(IDENTIFIER), VALUES...),
	}
	VALUE       Sequence = S(VALUES)
	DECLARATION          = join(S(TYPES, T(IDENTIFIER), T(ASSIGN)), EXPRESSION, S(T(SEMI)))
)

// T returns TokenType's in an array
func T(types ...TokenType) []TokenType {
	return types
}

// S returns []TokenType's in an array (Sequence)
func S(t ...[]TokenType) Sequence {
	first := [][]TokenType{}

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

// IsThing finds if the array is a Sequence
func IsSequence(tokens []TokenType, sequence Sequence) bool {
	length := len(tokens)
	if length != len(sequence) {
		return false
	}

	for i := 0; i < length; i++ {
		if !ContainsTT(tokens[i], sequence[i]) {
			return false
		}
	}

	return true
}
