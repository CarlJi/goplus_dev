package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to inspect the AST.
	src := `
package p
const c = 1.0
var X = f(3.14)*2 + c

func f(a float64) int {
	return int(a)
}
`
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, f.Decls[2])

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
			fmt.Printf("ast.BasicLit, %s:\t%s\n", s, fset.Position(n.Pos()))
		case *ast.Ident:
			s = x.Name
			fmt.Printf("ast.Ident %s:\t%s\n", s, fset.Position(n.Pos()))
		case *ast.FuncDecl:
			s = x.Name.Name
			fmt.Printf("ast.FuncDecl %s:\t%s\n", s, fset.Position(n.Pos()))
		}

		return true
	})
}
