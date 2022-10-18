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
			t.New(t.IDENTIFIER, "print"),
			t.New(t.LEFTPAREN, ""),
			t.New(t.STRINGVALUE, "Hello World!"),
			t.New(t.RIGHTPAREN, ""),
		}, false},
		{
			"Main function",
			args{"int main() {\nprint(\"Hello World!\");\n}"},
			[]t.Token{
				t.New(t.INTTYPE, ""),
				t.New(t.IDENTIFIER, "main"),
				t.New(t.LEFTPAREN, ""),
				t.New(t.RIGHTPAREN, ""),
				t.New(t.LEFTCURLY, ""),
				t.New(t.IDENTIFIER, "print"),
				t.New(t.LEFTPAREN, ""),
				t.New(t.STRINGVALUE, "Hello World!"),
				t.New(t.RIGHTPAREN, ""),
				t.New(t.SEMI, ""),
				t.New(t.RIGHTCURLY, ""),
			},
			false,
		},
	}

	fmt.Println(tests)
}
