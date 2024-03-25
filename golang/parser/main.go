package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

// .gop
// .gox

// index.gmx
// main.spx // project and class
// Car.spx  // class

// demo_yap.gox // class !normal => demo.yap
// demo.gox // class .gox normal
func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../testdata/main.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	fmt.Println(f.Name)
	for _, i := range f.Imports {
		fmt.Printf("%v %v %v\n", i.Name, i.Path, fset.Position(i.Pos()).Offset)
	}
	for _, d := range f.Decls {
		fmt.Printf("%T %#v\n", d, d)
	}
}
