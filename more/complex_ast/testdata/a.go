// This file is used to test the AST parser.
package main

import (
	"fmt"
	"strings"
)

type A struct {
	B
	ID int
}

type B struct {
	Name string
}

// NewA is a constructor for A
func NewA(id int, name string) *A {
	return &A{
		B: B{
			Name: name,
		},
		ID: id,
	}
}

// Print prints the A
func (a *A) Print() {
	fmt.Println(a.Name, a.ID)
}

const (
	GlobalID   = 1
	GlobalName = "john"
)

func main() {
	a := NewA(GlobalID, GlobalName)
	a.Print()

	// modify the name
	a.B.Name = strings.ToUpper(a.B.Name)
	a.Print()

	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			fmt.Printf("%d%%3==0\n", i)
		} else if i%3 == 1 {
			fmt.Printf("%d%%3==1\n", i)
		} else {
			fmt.Printf("%d%%3==2\n", i)
		}
	}
}
