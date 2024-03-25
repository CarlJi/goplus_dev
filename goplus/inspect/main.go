package main

import (
	"fmt"

	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../testdata/main.gop", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			if x.IsCommand() {
				fmt.Printf("%#v %v\n", n, fset.Position(n.Pos()))
			}
		}
		return true
	})
}
