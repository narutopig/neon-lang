package main

import (
	"fmt"
	"log"
	"os"

	"github.com/narutopig/neon-lang/lang/lexer"
)

func main() {
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
