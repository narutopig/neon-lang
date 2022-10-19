package runner

import "github.com/narutopig/neon-lang/lexer"

// Runner runs the program
type Runner struct {
	Lexer  lexer.Lexer
	memory map[string]string
}

// New returns a new Runner
func New(lexer lexer.Lexer) Runner {
	return Runner{Lexer: lexer, memory: make(map[string]string)}
}
