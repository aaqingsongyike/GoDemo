package main

import (
	"fmt"
)

func main() {
	//数组的定义
	var a [2]int
	var b [4]string
	fmt.Println(a)
	fmt.Println(b)
	//数组的简写方式
	c := [3]int{1, 1, 1}
	fmt.Println(c)
	//----------------------
	//多维数组
	d := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}} //注意
	fmt.Println(d)
	//---------------
	//冒泡排序
	e := [6]int{5, 3, 2, 4, 10, 0}
	fmt.Println(e)
	num := len(e)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if e[i] < e[j] {
				temp := e[i]
				e[i] = e[j]
				e[j] = temp
			}
		}
	}
	fmt.Println(e)
}
