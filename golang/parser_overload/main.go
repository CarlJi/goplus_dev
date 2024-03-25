package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
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
	f, err := parser.ParseFile(fset, "../testdata/overload.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	fmt.Println(f.Name)
	for _, i := range f.Imports {
		fmt.Printf("%v %v %v\n", i.Name, i.Path, fset.Position(i.Pos()).Offset)
	}
	//ast.FuncDecl{}
	for _, d := range f.Decls {
		fmt.Printf("%T %#v\n", d, d)
		switch d := d.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch spec := spec.(type) {
				case *ast.ValueSpec: // var/const
					for _, id := range spec.Names {
						if d.Tok == token.CONST && strings.HasPrefix(id.Name, "Gopo_") {
							if lit, ok := spec.Values[0].(*ast.BasicLit); ok {
								if s, err := strconv.Unquote(lit.Value); err == nil {
									for _, v := range strings.Split(s, ",") {
										if obj := f.Scope.Lookup(v); obj != nil {
											fmt.Println(obj)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
