package main

import (
	"aoc/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {

	is := antlr.NewInputStream("1+2+3")
	lexer := parser.NewCalcLexer(is)
	for {
		t := lexer.NextToken()
		fmt.Println(t)
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
	}
}
