package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tchajed/go-unused-funcs/unused"
)

func Two() int {
	return 2
}

func main() {
	uff := unused.NewUnusedFunctionFinder()
	flag.BoolVar(&(uff.Verbose), "v", false,
		"prints extra information during execution to stderr")
	flag.BoolVar(&(uff.IncludeAll), "all", false,
		"includes all found packages in analysis, not just main packages")
	flag.StringVar(&(uff.Ignore), "ignore", "",
		"don't read files that match the given string (use to avoid /testdata, etc) ")
	flag.StringVar(&(uff.CallgraphJSON), "callgraphjson", "",
		"pass in a callgraph in json format instead of computing one")
	flag.Parse()

	unusedFuncs, err := uff.Run(flag.Args())
	if err != nil {
		os.Exit(1)
	}

	for _, f := range unusedFuncs {
		fmt.Printf("%v in '%v'\n", f.Name, f.File)
	}
}
