package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	var files []*ast.File
	fset := token.NewFileSet()
	argIdx := 1
	if os.Args[1] == "--" {
		argIdx = 2
	}
	for _, gofile := range os.Args[argIdx:] {
		f, err := parser.ParseFile(fset, gofile, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}
		files = append(files, f)
	}

	for _, file := range files {
		var v visitor
		ast.Walk(v, file)
	}
}

type visitor int

func (v visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		v -= 1
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), node)
	v += 1
	return v
}
