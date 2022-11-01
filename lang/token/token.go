package token

import "fmt"

type Token struct {
	Type  TokenType
	Value string
}

func New(Type TokenType, Value string) Token {
	return Token{Type, Value}
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %s, Value: %s}", t.Type, t.Value)
}
