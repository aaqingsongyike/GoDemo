/*
并发
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	go Go() //goroutine1(简单)
	time.Sleep(2 * time.Second)
	//------------------------------
	//* goroutine2
	c := make(chan bool) //定义一个channel
	go func() {
		fmt.Println("goroutine2")
		c <- true //存操作
	}()
	<-c //取操作

	//使用for range来迭代不断操作channel
	c1 := make(chan bool) //定义一个channel
	go func() {
		fmt.Println("goroutine3")
		c1 <- true //存操作
		close(c1)  //*关闭
	}()
	for v := range c1 {
		fmt.Println(v)
	}

	//channel的缓存(有缓存是异步的，无缓存是同步阻塞的)
	c2 := make(chan bool, 1) //定义一个channel
	go func() {
		fmt.Println("goroutine4")
		<-c2 //取操作
	}()
	c2 <- true //存操作

	//有缓存eg1
	fmt.Println("有缓存1")
	runtime.GOMAXPROCS(runtime.NumCPU()) //返回CPU的核数
	c3 := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Eg(c3, i)
	}
	for i := 0; i < 10; i++ {
		<-c3
	}
	//有缓存eg2
	fmt.Println("有缓存2")
	runtime.GOMAXPROCS(runtime.NumCPU()) //返回CPU的核数
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Eg2(&wg, i)
	}
	wg.Wait()
}

func Go() {
	fmt.Println("Go Go Go!!!")
}

//有缓存eg1
func Eg(c3 chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c3 <- true
}

//有缓存eg2
func Eg2(wg *sync.WaitGroup, index int) {
	a1 := 1
	for i := 0; i < 100000000; i++ {
		a1 += i
	}
	fmt.Println(index, a1)
	wg.Done()
}
