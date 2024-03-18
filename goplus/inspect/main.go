package main

import (
	"fmt"

	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../data/main.gop", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	ast.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.CallExpr:
			fmt.Printf("%#v %v\n", n, fset.Position(n.Pos()))
		}
		return true
	})
}
