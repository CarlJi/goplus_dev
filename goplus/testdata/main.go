package main

import (
	"fmt"
)

// github.com/goplus/gogen
func demo[T ~int](v T) {
	fmt.Printf("%v %T\n", v, v)
}
