package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("needs 1 argument: file to process")
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, file, parser.Mode(0))
	if err != nil {
		fmt.Println(err)
	}

	ast.Print(fset, f)
}
