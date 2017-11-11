//可以使用for range 来迭代不断操作channal
package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Goroutine 2 !!!")
		c <- true

		close(c) //***
	}()

	//对channal进行迭代操作的时候 必须在某个地方明确关闭，否则会死锁
	for v := range c { //***
		fmt.Println(v)
	}
}
