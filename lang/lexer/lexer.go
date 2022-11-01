package lexer

import (
	"fmt"

	"github.com/narutopig/neon-lang/lang/token"
)

type Lexer struct {
	content string
	cursor  int
	tokens  []token.Token
}

func New(content string) Lexer {
	return Lexer{content: content, cursor: 0, tokens: []token.Token{}}
}

func (l *Lexer) Init() {
	for l.cursor < len(l.content) {
		l.Tokenize(l.content[l.cursor:])
	}

	for _, token := range l.tokens {
		fmt.Println(token)
	}
}
