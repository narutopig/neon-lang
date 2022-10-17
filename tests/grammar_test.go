package tests

import (
	"testing"

	"github.com/narutopig/neon-lang/lib"
	"github.com/narutopig/neon-lang/lib/grammar"
)

func TestIsSequence(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []lib.TokenType
		sequence grammar.Sequence
		want     bool
	}{
		{
			"Arithmetic operation 1",
			grammar.T(lib.NUMVALUE, lib.ADD, lib.NUMVALUE),
			grammar.EXPRESSION, true,
		},
		{
			"Arithmetic operation 2",
			grammar.T(lib.IDENTIFIER, lib.SUBTRACT, lib.NUMVALUE),
			grammar.EXPRESSION, true,
		},
		{
			"Arithmetic operation 3",
			grammar.T(lib.NUMVALUE, lib.MULTIPLY, lib.IDENTIFIER),
			grammar.EXPRESSION, true,
		},
		{
			"Arithmetic operation 4",
			grammar.T(lib.IDENTIFIER, lib.DIVIDE, lib.IDENTIFIER),
			grammar.EXPRESSION, true,
		},
		{
			"Arithmetic operation 5",
			grammar.T(lib.IDENTIFIER, lib.MODULUS, lib.IDENTIFIER),
			grammar.EXPRESSION, true,
		},
		{
			"Assigning",
			grammar.T(lib.INTTYPE, lib.IDENTIFIER, lib.ASSIGN, lib.NUMVALUE, lib.ADD, lib.NUMVALUE, lib.SEMI),
			grammar.DECLARATION, true,
		},
		/*
			{
				"Assigning to Expression",
				grammar.T(lib.INTTYPE, lib.IDENTIFIER, lib.ASSIGN, lib.NUMVALUE),
				grammar.DECLARATION, true,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grammar.IsSequence(tt.tokens, tt.sequence); got != tt.want {
				t.Errorf("IsSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
