package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	content := `
package foo

import (
	"fmt"
	"time"
)

// global lossy comment

func Foo() {
	fmt.Println("hello foo")
}

type Bar struct {
	i int

	// struct lossy comment

	j int
}
`
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", content, 0)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = format.Node(&buf, fset, node)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())
}
