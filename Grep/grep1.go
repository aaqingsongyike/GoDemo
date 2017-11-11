//Channal是goroutine沟通的桥梁，大都是阻塞同步
//goroutine中简单的channal操作
package main

import (
	"fmt"
)

func main() {
	//创建一个channal
	//直接make是双向通道，能存能取
	c := make(chan bool)
	//goroutine
	go func() {
		fmt.Println("Goroutine1!!!")
		//对channal的读取操作
		c <- true //存操作
	}()

	//阻塞goroutine
	<-c //取操作
}
