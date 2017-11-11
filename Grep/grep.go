//最简单的goroutine
package main

import (
	"fmt"
	"time"
)

/*
	main函数和Go函数同时运行
	若不暂停则不会显示Go函数中的fmt.Println()
*/
func main() {
	go Go()
	time.Sleep(2 * time.Second) //暂停2s否则不会输出
}

func Go() {
	fmt.Println("Go Go Go!!!")
}
