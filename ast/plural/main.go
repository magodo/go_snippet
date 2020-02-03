package main

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

var myAnalyzer = &analysis.Analyzer{
	Name:     "R1001",
	Doc:      "In schema definition, List/Set Property should reflects plurality by nameing",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func main() {
	singlechecker.Main(myAnalyzer)
}

func run(pass *analysis.Pass) (interface{}, error) {
	schemaMapNode, err := LookupSchemaMapNode(pass)
	if err != nil {
		return nil, err
	}
	//ast.Walk(visitor{pass: pass}, schemaMapNode)

	schemaResourceObj, err := LookupSchemaObject(pass, "Resource")
	if err != nil {
		return nil, err
	}
	schemaResourcePtrType := types.NewPointer(schemaResourceObj.Type())
	schemaTypeListObj, err := LookupSchemaObject(pass, "TypeList")
	if err != nil {
		return nil, err
	}
	schemaTypeSetObj, err := LookupSchemaObject(pass, "TypeSet")
	if err != nil {
		return nil, err
	}

	typefilter := []ast.Node{
		&ast.KeyValueExpr{},
	}
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder(typefilter, func(n ast.Node) {
		// ignore node which is contained in schema map node
		if !IsParentNodeOf(schemaMapNode, n) {
			return
		}

		// since we filter out KeyValueExpr only, so it's safe to cast
		kvnode := n.(*ast.KeyValueExpr)

		// key should be basic literature, string actually
		key, ok := kvnode.Key.(*ast.BasicLit)
		if !ok {
			return
		}
		// value should be composite literal, *schema.Schema actually
		value, ok := kvnode.Value.(*ast.CompositeLit)
		if !ok {
			return
		}

		// iterate each property's elements and gather information
		var isList bool
		var isOneItemLimit bool
		var isNestedBlock bool
		for _, ele := range value.Elts {
			kvnode, ok := ele.(*ast.KeyValueExpr)
			if !ok {
				continue
			}
			key, ok := kvnode.Key.(*ast.Ident)
			if !ok {
				continue
			}

			switch key.Name {
			case "Type":
				valueValue := pass.TypesInfo.Types[kvnode.Value].Value
				if constant.Compare(valueValue, token.EQL, schemaTypeSetObj.(*types.Const).Val()) ||
					constant.Compare(valueValue, token.EQL, schemaTypeListObj.(*types.Const).Val()) {
					isList = true
				}
			case "MaxItems":
				value, ok := kvnode.Value.(*ast.BasicLit)
				if !ok {
					continue
				}
				if value.Value == "1" {
					isOneItemLimit = true
				}
			case "Elem":
				valueType := pass.TypesInfo.Types[kvnode.Value].Type
				if types.Identical(schemaResourcePtrType, valueType) {
					isNestedBlock = true
				}
			}
		}

		// For list property that allows more than one items or nested block,
		// the property name should reflect plurality.
		if isList {
			if !isOneItemLimit || isNestedBlock {
				keyVal := strings.Trim(key.Value, `"`)
				if !isPlural(keyVal) {
					pass.Report(analysis.Diagnostic{
						Pos:     kvnode.Pos(),
						End:     kvnode.End(),
						Message: fmt.Sprintf("%s should be plural(R1001)", key.Value),
					})
				}
			}
		}
	})

	return nil, nil
}

// TODO: enhancement needed
func isPlural(s string) bool {
	return strings.HasSuffix(s, "s") || strings.HasSuffix(s, "list")
}

type visitor struct {
	pass  *analysis.Pass
	level int
}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		v.level--
	}
	output := fmt.Sprintf("%s%v(%T)", strings.Repeat("\t", v.level), n, n)
	if expr, ok := interface{}(n).(ast.Expr); ok {
		output += fmt.Sprintf("%v", v.pass.TypesInfo.Types[expr])
	}
	fmt.Println(output)
	v.level++
	return v
}
