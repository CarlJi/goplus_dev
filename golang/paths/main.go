package main

import (
	"fmt"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../testdata/main.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	fmt.Println(f.Name)
	paths, exact := astutil.PathEnclosingInterval(f, 162, 162)
	fmt.Println(paths, exact)
	for _, path := range paths {
		fmt.Printf("%#v %v\n", path, fset.Position(path.Pos()))
	}
}
