package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

func LookupSchemaObject(pass *analysis.Pass, t string) (types.Object, error) {
	const schemaPkgPath = "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	for _, importpkg := range pass.Pkg.Imports() {
		if importpkg.Path() == schemaPkgPath {
			obj := importpkg.Scope().Lookup(t)
			if obj == nil {
				return nil, fmt.Errorf(`Object "%s" not defined in %s`, t, schemaPkgPath)
			}
			return obj, nil
		}
	}
	return nil, fmt.Errorf("source file not imported: %s", schemaPkgPath)
}

// LookupSchemaResourceNode look up for node (*ast.CompositeLit) representing schema.Resource composite literal
func LookupSchemaResourceNode(pass *analysis.Pass) (*ast.CompositeLit, error) {
	schemaResourceObj, err := LookupSchemaObject(pass, "Resource")
	if err != nil {
		return nil, err
	}
	schemaResourceType := schemaResourceObj.Type()

	// Look up for package level function which match schema def function signature.
	// We assume each source file has up to 1 such function.
	pkgscope := pass.Pkg.Scope()
	for _, name := range pkgscope.Names() {
		obj := pkgscope.Lookup(name)

		// object is a function
		funcobj, ok := obj.(*types.Func)
		if !ok {
			continue
		}

		// type is a signature
		functype, ok := funcobj.Type().(*types.Signature)
		if !ok {
			continue
		}

		// signature match target signature
		if !(functype.Params().Len() == 0 && functype.Results().Len() == 1 && types.Identical(functype.Results().At(0).Type(), types.NewPointer(schemaResourceType))) {
			continue
		}

		// types.Object -> ast.Node
		declNode, err := Typeobject2Astnode4IdentDecl(pass, obj)
		if err != nil {
			return nil, err
		}

		funcDeclNode := declNode.(*ast.FuncDecl)

		var retstmt *ast.ReturnStmt
		for _, stmt := range funcDeclNode.Body.List {
			var ok bool
			retstmt, ok = stmt.(*ast.ReturnStmt)
			if !ok {
				continue
			}
		}
		if retstmt == nil {
			return nil, errors.New("failed to find return statement in schema def function")
		}

		resNode, ok := retstmt.Results[0].(*ast.UnaryExpr).X.(*ast.CompositeLit)
		if !ok {
			return nil, fmt.Errorf("returned node is not of type *ast.CompositeLit, but %T",
				retstmt.Results[0].(*ast.UnaryExpr).X)
		}
		return resNode, nil
	}

	return nil, errors.New("package level function matching schema definition not found")
}

// LookupSchemaMapNode look up for node (*ast.CompositeLit) representing schema.Resource.Schema composite literal
func LookupSchemaMapNode(pass *analysis.Pass) (*ast.CompositeLit, error) {
	root, err := LookupSchemaResourceNode(pass)
	if err != nil {
		return nil, err
	}

	for _, element := range root.Elts {
		kvnode, ok := element.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		if kvnode.Key.(*ast.Ident).Name == "Schema" {
			return kvnode.Value.(*ast.CompositeLit), nil
		}
	}
	return nil, errors.New(`not found an "Schema" key`)
}
