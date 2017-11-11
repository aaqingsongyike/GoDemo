/*
select是针对channal创建的结构
可以处理多个channal的发送与接收
按随机顺序处理
*/

package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	//用于通信channal：o
	o := make(chan bool)

	go func() {
		//确保o传入两次
		a, b := false, false

		//该组合实现不断的信息接收与发送操作
		for {
			select {
			//判断value和ok，取出c1的值
			case v, ok := <-c1:
				if !ok { //c1关闭

					if !a {
						a = true

						//c1或c2当中有任意一个关闭的时候就给channal：o传进一个值
						o <- true
					}

					break
				}
				fmt.Println("c1:", v) //没有关闭
			case v, ok := <-c2:
				if !ok {
					if !b {
						b = true

						o <- true
					}
					break
				}
				fmt.Println("c2:", v)
			}
		}
	}()

	//将值传进去
	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	//关闭channal
	//也可以只关闭其中一个
	close(c1)
	close(c2)

	//读出channal：o
	<-o
}
