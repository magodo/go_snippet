package addcheck

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"os"
)

func Addcheck() {
	var files []*ast.File
	fset := token.NewFileSet()
	for _, gofile := range os.Args[1:] {
		f, err := parser.ParseFile(fset, gofile, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}
		files = append(files, f)
	}

	typeConf := types.Config{}
	typeInfo := &types.Info{
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	if _, err := typeConf.Check("addcheck", fset, files, typeInfo); err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		ast.Inspect(file, func(node ast.Node) (ret bool) {
			ret = true
			be, ok := node.(*ast.BinaryExpr)
			if !ok {
				return
			}
			if be.Op != token.ADD {
				return
			}
			if _, ok := be.X.(*ast.BasicLit); !ok {
				return
			}
			if _, ok := be.Y.(*ast.BasicLit); !ok {
				return
			}
			if !isInteger(typeInfo, be.X) || !isInteger(typeInfo, be.Y) {
				return
			}
			fmt.Fprintf(os.Stderr, "%s: find addition expression: %s\n", fset.Position(be.Pos()), render(fset, be))
			return
		})
	}
}

func render(fset *token.FileSet, node ast.Node) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, node); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func isInteger(tinfo *types.Info, expr ast.Expr) bool {
	t := tinfo.TypeOf(expr)
	if t == nil {
		return false
	}

	bt, ok := t.Underlying().(*types.Basic)
	if !ok {
		return false
	}

	if (bt.Info() & types.IsInteger) == 0 {
		return false
	}
	return true
}
