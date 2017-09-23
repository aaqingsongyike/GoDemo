package main

import (
	"fmt"
)

func main() {
	a := 1
	var p *int = &a
	//&表示内存地址	*表示值
	fmt.Println(p)  //内存地址
	fmt.Println(*p) //值
	b := 2
	fmt.Println(b)
}
