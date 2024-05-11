package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	code := `
	package main

func main() {
	x := 10
	y := 5
	result := x + y*2
}
	`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, decl := range file.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			for _, stmt := range funcDecl.Body.List {
				fmt.Printf("Statement: %s\n", stmt)
				if assignStmt, ok := stmt.(*ast.AssignStmt); ok {
					fmt.Printf("Expression: %v\n", assignStmt.Lhs)
					for _, expr := range assignStmt.Rhs {
						if binaryExpr, ok := expr.(*ast.BinaryExpr); ok {
							fmt.Printf("Binary Expression: %s\n", binaryExpr.Op)
							fmt.Printf("Left Operand: %s\n", binaryExpr.X)
							fmt.Printf("Right Operand: %s\n", binaryExpr.Y)
						}
					}
				}
			}
		}
	}
}
