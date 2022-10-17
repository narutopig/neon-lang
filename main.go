package main

import (
	"fmt"
	"log"
	"os"

	l "github.com/narutopig/neon-lang/lexer"
)

func main() {
	/*
		g.Validate(
			[]l.TokenType{l.INTTYPE, l.IDENTIFIER, l.ASSIGN, l.NUMVALUE, l.SEMI},
			[]g.Segment{
				g.NewSegmentS(g.S(g.TYPES, g.T(l.IDENTIFIER), g.T(l.ASSIGN))),
				g.NewSegment(false, []g.Sequence{g.S(g.VALUES), g.S(append(g.T(l.IDENTIFIER), g.VALUES...), append(g.ARITHOPS, g.COMPOPS...), append(g.T(l.IDENTIFIER), g.VALUES...))}),
				g.NewSegmentS(g.S(g.T(l.SEMI))),
			},
		)
	*/

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

	tokens, err := l.Lex(content)
	if err != nil {
		log.Panic(err)
		return
	}

	for _, t := range tokens {
		fmt.Println(t)
	}
}
