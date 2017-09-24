/*
并发select  可以同时处理多个channel的发送与接收
*/
package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2) //用于通信
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					/*
						c1或c2关闭的时候就传一个o
					*/
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					/*
						c1或c2关闭的时候就传一个o
					*/
					o <- true
					break
				}
				fmt.Println("c2", v)

			}
		}
	}()
	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1) //关闭其中1个就可以了

	for i := 0; i < 2; i++ { //确保都关闭
		<-o //读出值(读通道)
	}
}
