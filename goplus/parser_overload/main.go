package main

import (
	"fmt"

	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
)

// .spx
// _yap.gox
func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../testdata/main.gop", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	fmt.Println("======>")
	ast.Print(fset, f)
	fmt.Println("======>")
	fmt.Println(f)
	fmt.Println(f.Name)
	for _, i := range f.Imports {
		fmt.Printf("%v %v\n", i.Name, i.Path)
	}
	for _, d := range f.Decls {
		fmt.Printf("%T %v\n", d, d)
		switch d := d.(type) {
		case *ast.OverloadFuncDecl:
			switch v := d.Funcs[0].(type) {
			case *ast.Ident:
				if obj := f.Scope.Lookup(v.Name); obj != nil {
					if decl, ok := obj.Decl.(*ast.FuncDecl); ok {
						fmt.Println(decl.Type)
					}
				}
			case *ast.FuncLit:
				//v.Type
				fmt.Println(v.Type)
			}
		}
	}
}
