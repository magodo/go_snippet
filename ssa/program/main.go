package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	cfg := packages.Config{Mode: packages.LoadAllSyntax}
	pkgs, err := packages.Load(&cfg, os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}

	// Stop if any package had errors.
	// This step is optional; without it, the previous step
	// will create SSA for only a subset of packages.
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	// Build SSA for the specified "pkgs" and their dependencies.
	// The returned ssapkgs is the corresponding SSA Package of the specified "pkgs".
	prog, ssapkgs := ssautil.AllPackages(pkgs, 0)
	prog.Build()
	fmt.Println(ssapkgs)
}
