package main

import (
	"fmt"
	"go/token"

	"github.com/visualfc/goplus_dev/load"
)

func main() {
	fset := token.NewFileSet()
	pkg, info, err := load.LoadPackage(fset, "main", "../testdata/main.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(pkg)
	fmt.Println(pkg.Scope())
	fmt.Println(info)
}
