package main

import (
	"fmt"
)

const Pi = 3.14

var version = "1.0"

type T struct {
	X int
	Y int
}

func (t *T) Info() string {
	return fmt.Sprintf("%v-%v", t.X, t.Y)
}

func main() {
	fmt.Println(version)
	a := 100
	_ = a
}
