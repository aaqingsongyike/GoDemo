package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}
type person1 struct {
	Name string
	Age  int
}
type human struct { //组合（嵌入）
	Sex  int
	Name string
	Age  int
}
type teacher struct {
	human
}
type students struct {
	human
	Hobby string
}

func main() {
	a := person{}
	a.Name = "joe"
	a.Age = 22
	fmt.Println(a)
	b := &person1{ //推荐写&
		Name: "aa",
		Age:  23,
	}
	fmt.Println(b)
	tea := teacher{
		human: human{
			Sex:  0,
			Name: "Mary",
			Age:  29,
		},
	}
	stu := students{
		human: human{
			Sex:  0,
			Name: "jone",
			Age:  22,
		},
		Hobby: "篮球",
	}
	fmt.Println("teacher", tea, "students", stu)
}
