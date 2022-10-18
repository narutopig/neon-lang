package tests

import (
	"fmt"
	"testing"

	l "github.com/narutopig/neon-lang/lexer"
)

func TestLex(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    []l.Token
		wantErr bool
	}{
		{"Simple hello world", args{"print(\"Hello World!\")"}, []l.Token{
			l.NewToken(l.IDENTIFIER, "print"),
			l.NewToken(l.LEFTPAREN, ""),
			l.NewToken(l.STRINGVALUE, "Hello World!"),
			l.NewToken(l.RIGHTPAREN, ""),
		}, false},
		{
			"Main function",
			args{"int main() {\nprint(\"Hello World!\");\n}"},
			[]l.Token{
				l.NewToken(l.INTTYPE, ""),
				l.NewToken(l.IDENTIFIER, "main"),
				l.NewToken(l.LEFTPAREN, ""),
				l.NewToken(l.RIGHTPAREN, ""),
				l.NewToken(l.LEFTCURLY, ""),
				l.NewToken(l.IDENTIFIER, "print"),
				l.NewToken(l.LEFTPAREN, ""),
				l.NewToken(l.STRINGVALUE, "Hello World!"),
				l.NewToken(l.RIGHTPAREN, ""),
				l.NewToken(l.SEMI, ""),
				l.NewToken(l.RIGHTCURLY, ""),
			},
			false,
		},
	}

	fmt.Println(tests)
}
