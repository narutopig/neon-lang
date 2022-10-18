package tests

import (
	"fmt"
	"testing"

	t "github.com/narutopig/neon-lang/token"
)

func TestLex(test *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    []t.Token
		wantErr bool
	}{
		{"Simple hello world", args{"print(\"Hello World!\")"}, []t.Token{
			t.NewToken(t.IDENTIFIER, "print"),
			t.NewToken(t.LEFTPAREN, ""),
			t.NewToken(t.STRINGVALUE, "Hello World!"),
			t.NewToken(t.RIGHTPAREN, ""),
		}, false},
		{
			"Main function",
			args{"int main() {\nprint(\"Hello World!\");\n}"},
			[]t.Token{
				t.NewToken(t.INTTYPE, ""),
				t.NewToken(t.IDENTIFIER, "main"),
				t.NewToken(t.LEFTPAREN, ""),
				t.NewToken(t.RIGHTPAREN, ""),
				t.NewToken(t.LEFTCURLY, ""),
				t.NewToken(t.IDENTIFIER, "print"),
				t.NewToken(t.LEFTPAREN, ""),
				t.NewToken(t.STRINGVALUE, "Hello World!"),
				t.NewToken(t.RIGHTPAREN, ""),
				t.NewToken(t.SEMI, ""),
				t.NewToken(t.RIGHTCURLY, ""),
			},
			false,
		},
	}

	fmt.Println(tests)
}
