package tests

import (
	"testing"

	"github.com/narutopig/neon-lang/lib"
)

func TestIsSequence(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []lib.TokenType
		sequence lib.Sequence
		want     bool
	}{
		{
			"Arithmetic operation 1",
			lib.T(lib.NUMVALUE, lib.ADD, lib.NUMVALUE),
			lib.EXPRESSION, true,
		},
		{
			"Arithmetic operation 2",
			lib.T(lib.IDENTIFIER, lib.SUBTRACT, lib.NUMVALUE),
			lib.EXPRESSION, true,
		},
		{
			"Arithmetic operation 3",
			lib.T(lib.NUMVALUE, lib.MULTIPLY, lib.IDENTIFIER),
			lib.EXPRESSION, true,
		},
		{
			"Arithmetic operation 4",
			lib.T(lib.IDENTIFIER, lib.DIVIDE, lib.IDENTIFIER),
			lib.EXPRESSION, true,
		},
		{
			"Arithmetic operation 5",
			lib.T(lib.IDENTIFIER, lib.MODULUS, lib.IDENTIFIER),
			lib.EXPRESSION, true,
		},
		{
			"Assigning",
			lib.T(lib.INTTYPE, lib.IDENTIFIER, lib.ASSIGN, lib.NUMVALUE, lib.ADD, lib.NUMVALUE, lib.SEMI),
			lib.DECLARATION, true,
		},
		/*
			{
				"Assigning to Expression",
				lib.T(lib.INTTYPE, lib.IDENTIFIER, lib.ASSIGN, lib.NUMVALUE),
				lib.DECLARATION, true,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lib.IsSequence(tt.tokens, tt.sequence); got != tt.want {
				t.Errorf("IsSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
