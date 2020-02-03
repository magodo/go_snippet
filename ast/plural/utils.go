package main

import (
	"errors"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ast/astutil"
)

func LookupTypeobjectDefFile(pass *analysis.Pass, obj types.Object) *ast.File {
	var file *ast.File
	for _, f := range pass.Files {
		if pass.Fset.Position(f.Pos()).Filename == pass.Fset.Position(obj.Pos()).Filename {
			file = f
			break
		}
	}
	return file
}

func Typeobject2Astnode4IdentDecl(pass *analysis.Pass, obj types.Object) (ast.Node, error) {
	f := LookupTypeobjectDefFile(pass, obj)
	if f == nil {
		return nil, errors.New("no file defines this Object")
	}

	path, _ := astutil.PathEnclosingInterval(f, obj.Pos(), obj.Pos())
	// the 1st node is the object's declaring identifier,
	// by walking up one level, we get the enclosing declaration
	if len(path) < 2 {
		return nil, errors.New("enclosing interval has less than 2 levels of ast.Node(s)")
	}
	declNode := path[1]
	return declNode, nil
}

func IsParentNodeOf(parent ast.Node, child ast.Node) bool {
	return parent.Pos() <= child.Pos() && parent.End() >= child.End() && parent != child
}
