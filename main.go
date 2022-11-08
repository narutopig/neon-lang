package main

import (
	"fmt"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/narutopig/neon-lang/listener"
	"github.com/narutopig/neon-lang/parser"
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

	is := antlr.NewInputStream(content)
	lexer := parser.NewNeonLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewNeonParser(stream) // Create the Parser
	i := listener.NewInterpreter()
	antlr.ParseTreeWalkerDefault.Walk(i, p.Program())
	i.Print()
	fmt.Println(i.Memory)
}
