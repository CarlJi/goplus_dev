package complexast_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestXXX(t *testing.T) {
	// 创建 token.FileSet
	fset := token.NewFileSet()

	// 解析代码
	node, err := parser.ParseFile(fset, "/Users/jicarl/goplus/goplus_dev/more/complex_ast/testdata/a.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println("解析代码出错:", err)
		return
	}

	ast.Print(fset, node)
}
