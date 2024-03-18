package main

import (
	"fmt"

	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"

	"golang.org/x/tools/gop/ast/astutil"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../data/main.gop", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	fmt.Println(f.Name)
	paths, exact := astutil.PathEnclosingInterval(f, 146, 146)
	fmt.Println(paths, exact)
	for _, path := range paths {
		fmt.Printf("%#v %v\n", path, fset.Position(path.Pos()))
	}
}
