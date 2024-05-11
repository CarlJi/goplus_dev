package main

import (
	"fmt"

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
	fmt.Println(f)
	fmt.Println(f.Name)
	fmt.Println("======>")
	for _, i := range f.Imports {
		fmt.Printf("%v %v\n", i.Name, i.Path)
	}
	fmt.Println("======>")
	for _, d := range f.Decls {
		fmt.Printf("%T %v\n", d, d)
	}
}
