package main

import (
	"fmt"
)

func main() {
	//slice切片的声明方法1
	s1 := []int{}
	fmt.Println(s1)
	//slice切片的声明方法2
	s2 := make([]int, 3, 10)      //make(类型,值的数量,初始容量)
	fmt.Println(len(s2), cap(s2)) //len()获取元素个数，cap()获取元素容量
	//--------------------------------
	//slice获取数组中的元素
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //数组
	s3 := a[4:8]                               //通过切片slice获取元素
	fmt.Println(s3)
	//-------------------
	//append追加函数
	s4 := make([]int, 4, 8)
	fmt.Printf("%p\n", s4)   //s4的内存地址
	s4 = append(s4, 1, 2, 3) //append(被追加的函数,值)
	fmt.Printf("%v %p\n", s4, s4)
}
