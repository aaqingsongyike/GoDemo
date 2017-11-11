/*
同步解决并发问题
*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(10) //增加任务数为10

	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}

	wg.Wait() //等待
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println("index =", index, "a = ", a)

	wg.Done() //标记Done  表示完成一次消去一次

}
