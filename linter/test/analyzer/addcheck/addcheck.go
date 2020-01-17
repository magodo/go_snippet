package addcheck

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"os"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "addlint",
	Doc:      "reports integer additions",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(node ast.Node) {
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
		if !isInteger(pass.TypesInfo, be.X) || !isInteger(pass.TypesInfo, be.Y) {
			return
		}
		fmt.Fprintf(os.Stderr, "%s: find addition expression: %s\n", pass.Fset.Position(be.Pos()), render(pass.Fset, be))
		return
	})

	return nil, nil
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
