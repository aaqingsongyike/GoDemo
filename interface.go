package main

import (
	"fmt"
)

type USB interface { //接口
	Name() string //方法1
	Connect()     //方法2
}
type PhoneConnecter struct { //实现接口的结构
	name string
}

func (pc PhoneConnecter) Name() string { //实现方法1
	return pc.name
}
func (pc PhoneConnecter) Connect() { //实现方法2
	fmt.Println("Connect：", pc.name)
}

func Disconnect(usb USB) {
	/*
	if pc, ok := usb.(PhoneConnecter1); ok { //判断是不是传进来PhoneConnection(类型判断)
		fmt.Println("Disconnect.", pc.name)
		return
	}
	fmt.Println("Unknow decive")
	*/
	switch v:= usb.(type){	//作用同上
		case PhoneConnecter:
			fmt.Println("Disconnect:",v.name)
		default:
			fmt.Println("Unknow decive")
	}
}
func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}
