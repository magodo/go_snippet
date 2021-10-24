package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/tools/go/callgraph"

	"path"

	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	// pkgs, err := loadPackage("/home/magodo/projects/terraform-provider-azurerm", []string{"./internal/services/web"})
	pwd, _ := os.Getwd()
	pkgs, err := loadPackage(path.Join(path.Dir(pwd), "analyzee"), []string{".", "./lib"})
	if err != nil {
		log.Fatal(err)
	}

	// pkg := pkgs[0]

	// ssapkg, _, err := ssautil.BuildPackage(
	// 	&types.Config{Importer: importer.Default()}, pkg.Fset, pkg.Types, pkg.Syntax, ssa.PrintFunctions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _ = ssapkg

	// Create SSA packages for all well-typed packages.
	prog, ssapkgs := ssautil.Packages(pkgs, ssa.SanityCheckFunctions)
	_ = prog

	// Build SSA code for the well-typed initial packages.
	for _, p := range ssapkgs {
		if p != nil {
			p.Build()
		}
	}

	graph := cha.CallGraph(prog)

	ssamainpkgs := ssautil.MainPackages(ssapkgs)
	rootFunc := ssamainpkgs[0].Members["main"].(*ssa.Function)
	rootNode := graph.Nodes[rootFunc]
	fmt.Println(AllCalleesOf(rootNode))
}

func loadPackage(dir string, args []string) ([]*packages.Package, error) {
	cfg := packages.Config{Dir: dir, Mode: packages.LoadAllSyntax}
	pkgs, err := packages.Load(&cfg, args...)
	if err != nil {
		return nil, err
	}

	if packages.PrintErrors(pkgs) > 0 {
		return nil, errors.New("packages contain errors")
	}

	return pkgs, nil
}

func AllCalleesOf(root *callgraph.Node) map[*callgraph.Node]bool {
	s := map[*callgraph.Node]bool{}
	wl := map[*callgraph.Node]bool{root: true}
	for len(wl) != 0 {
		callees := map[*callgraph.Node]bool{}
		for node := range wl {
			directCallees := callgraph.CalleesOf(node)
			for k := range directCallees {
				callees[k] = true
			}
		}
		for k := range wl {
			s[k] = true
		}

		wl = map[*callgraph.Node]bool{}
		for k := range callees {
			if !s[k] {
				wl[k] = true
			}
		}
	}
	delete(s, root)
	return s
}
