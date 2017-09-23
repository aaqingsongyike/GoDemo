package main

import (
	"fmt"
)

func main() {
	//map的声明
	var m map[int]string //map key的类型 value的类型
	m = map[int]string{}
	fmt.Println(m)
	//map的应用
	m1 := map[int]string{}
	m1[0] = "ok" //存
	a := m1[0]   //取
	fmt.Println(m1)
	fmt.Println(a)
	delete(m1, 0) //删除
	a = m1[0]
	fmt.Println("删除后" + a + "为空")
	//-----------------------------------
	//多返回值
	var m2 map[int]map[int]string
	m2 = make(map[int]map[int]string)
	b, ok := m2[2][1] //多返回值 第1个返回对应的value   第2个返回boolean类型
	fmt.Println(b, ok)
	if !ok {
		m2[2] = make(map[int]string)
	}
	m2[2][1] = "GOOD"
	b, ok = m2[2][1]
	fmt.Println(b, ok)
	fmt.Println(m2)
}
