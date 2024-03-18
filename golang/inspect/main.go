package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../data/main.go", nil, parser.ParseComments)
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
