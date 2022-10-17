package tests

import (
	"testing"

	g "github.com/narutopig/neon-lang/grammar"
	l "github.com/narutopig/neon-lang/lexer"
)

func TestIsSequence(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []l.TokenType
		sequence g.Sequence
		want     bool
	}{
		{
			"Arithmetic operation 1",
			g.T(l.NUMVALUE, l.ADD, l.NUMVALUE),
			g.EXPRESSION, true,
		},
		{
			"Arithmetic operation 2",
			g.T(l.IDENTIFIER, l.SUBTRACT, l.NUMVALUE),
			g.EXPRESSION, true,
		},
		{
			"Arithmetic operation 3",
			g.T(l.NUMVALUE, l.MULTIPLY, l.IDENTIFIER),
			g.EXPRESSION, true,
		},
		{
			"Arithmetic operation 4",
			g.T(l.IDENTIFIER, l.DIVIDE, l.IDENTIFIER),
			g.EXPRESSION, true,
		},
		{
			"Arithmetic operation 5",
			g.T(l.IDENTIFIER, l.MODULUS, l.IDENTIFIER),
			g.EXPRESSION, true,
		},
		{
			"Assigning",
			g.T(l.INTTYPE, l.IDENTIFIER, l.ASSIGN, l.NUMVALUE, l.ADD, l.NUMVALUE, l.SEMI),
			g.DECLARATION, true,
		},
		/*
			{
				"Assigning to Expression",
				g.T(lib.INTTYPE, lib.IDENTIFIER, lib.ASSIGN, lib.NUMVALUE),
				g.DECLARATION, true,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.IsSequence(tt.tokens, tt.sequence); got != tt.want {
				t.Errorf("IsSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
