package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
)

func main() {
	filename, src, err := getInput()
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, parser.Mode(0))
	if err != nil {
		fmt.Println(err)
	}

	ast.Print(fset, f)
}

func getInput() (string, []byte, error) {
	filename := "input.go"
	file := os.Stdin

	if !hasStdin() {
		if len(os.Args) != 2 {
			return "", nil, errors.New("needs 1 argument: file to process")
		}
		filename = os.Args[1]

		var err error
		file, err = os.Open(filename)
		if err != nil {
			return "", nil, err
		}
	}

	buf, err := io.ReadAll(file)
	if err != nil {
		return "", nil, err
	}
	return filename, buf, nil
}

func hasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fi.Size() > 0
}
