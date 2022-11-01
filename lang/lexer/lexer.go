package lexer

import (
	"fmt"
	"sort"

	"github.com/narutopig/neon-lang/lang/token"
)

type Lexer struct {
	content string
	cursor  int
	keys    []token.TokenType // sorted array of keys (maps get scrambled weirdly otherwise)
	tokens  []token.Token
}

func New(content string) Lexer {
	return Lexer{content: content, cursor: 0, tokens: []token.Token{}}
}

func (l *Lexer) Init() {
	l.setupKeys()

	for l.cursor < len(l.content) {
		l.Tokenize(l.content[l.cursor:])
	}

	for _, tok := range l.tokens {
		fmt.Println(tok)
	}
}

func (l *Lexer) setupKeys() {
	bk := make([]int, 0, len(token.Spec))

	for key := range token.Spec {
		bk = append(bk, int(key))
	}

	sort.Ints(bk)

	for val := range bk {
		l.keys = append(l.keys, token.TokenType(val))
	}
}
