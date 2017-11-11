//设置channal的缓存
//channal无缓存是可以先取再存（同步阻塞）
//channal有缓存是异步的
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10) //设置channal的缓存为10
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println("index =", index, "a = ", a)
	c <- true

}
