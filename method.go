package main

import (
	"fmt"
)

type A struct {
	Name string
}
type B struct {
	Name string
}

func main() {
	a := A{}
	a.aa()
	fmt.Println(a.Name)

	b := B{}
	b.aa()
	fmt.Println(b.Name)
}

func (a *A) aa() { //method方法
	a.Name = "AA"
	fmt.Println("A")
}
func (b B) aa() {
	b.Name = "BB"
	fmt.Println("B")
}
