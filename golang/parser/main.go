package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../data/main.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	fmt.Println(f.Name)
	for _, i := range f.Imports {
		fmt.Printf("%v %v\n", i.Name, i.Path)
	}
	for _, d := range f.Decls {
		fmt.Printf("%T %v\n", d, d)
	}
}
