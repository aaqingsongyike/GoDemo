/*反射*/
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() { //方法
	fmt.Println("Hello world")
}
func main() {
	u := User{1, "OK", 12}
	Info(u) //值copy
}
func Info(o interface{}) { //方法 用于传进一个空接口，输出传进去的具体信息
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct { //判断是否是想要的类型（是否是地址copy而不是值copy）
		fmt.Println("XX")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ { //如何获取某个字段的信息以及类型的信息
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ { //取得方法信息
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
