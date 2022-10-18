package main

import (
	"fmt"
	"log"
	"os"

	"github.com/narutopig/neon-lang/expressions"
	"github.com/narutopig/neon-lang/lexer"
	t "github.com/narutopig/neon-lang/token"
)

func main() {
	ts, ne := expressions.Shunt([]t.Token{
		t.New(t.NUMVALUE, "4"),
		t.New(t.ADD, ""),
		t.New(t.NUMVALUE, "18"),
		t.New(t.DIVIDE, ""),
		t.New(t.LEFTPAREN, ""),
		t.New(t.NUMVALUE, "9"),
		t.New(t.SUBTRACT, ""),
		t.New(t.NUMVALUE, "3"),
		t.New(t.RIGHTPAREN, ""),
	},
		1)

	val, ne := expressions.Eval(ts)

	fmt.Println(val)

	// cli entry point
	if len(os.Args) < 2 {
		log.Panic("No input file found")
		os.Exit(1)
	}

	var fileName string = os.Args[1]

	c, err := os.ReadFile(fileName)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	content := string(c)

	l := lexer.New(content)
	tokens, ne := l.Tokenize()

	if !ne.IsNull() {
		fmt.Println(ne)
	} else {
		fmt.Println(len(tokens))
		for _, t := range tokens {
			fmt.Println(t)
		}
	}
}
