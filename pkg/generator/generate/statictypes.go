package main

import (
	"os"

	"github.com/gobwas/glob/syntax/ast"
	"github.com/gobwas/glob/syntax/lexer"
)

func generateStaticTypes() {

	content, err := os.ReadFile("../types.go")
	if err != nil {
		panic(err)
	}

	lexer := lexer.NewLexer(string(content))

	dat, err := ast.Parse(lexer)
	if err != nil {
		panic(err)
	}

	// for all interfaces, generate the file that looks like statictypes-example.gen.go

	if err := os.WriteFile("statictypes.gen.go", []byte(dat), 0644); err != nil {
		panic(err)
	}

}
