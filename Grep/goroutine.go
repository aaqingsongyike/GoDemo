/*
并发编程的一种常见的方式就是有很多工作需要处理，且每个工作都可以独立的完成。
Go的标准库里的 net/http 包的HTTP服务器利用这种模式来处理并发，每个请求都在
独立的goroutine里处理，和其他的goroutine之间没有任何通信
*/
package main

import(
	//"net/http"
	"os"
	"fmt"
	"runtime"
	"regexp"
	"log"
	"path/filepath"
)

type Job struct {
	filename string		//需要被处理的文件
	results chan <- Result	//results是一个通道，所有处理完的文件都会被发送到这来
							/*将results定义为一个chan Result类型，但指往通道里发送数据，不会从里面读取数据
							，所以指定这里是一个单向的只允许发送的通道。
							*/
}

type Result struct {
	filename string 
	lino int
	line string 
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())	//使用所有的机器核心
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <regexp> <file>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatalf("invalid regexp: %s\n", err)
	} else {
		grep(lineRx, commandLineFile(os.Args[2:]))
	}
}

var workers = runtime.NumCPU()

func grep(linePx *regexp.Regexp, filenames []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, minimum(1000, len(filenames)))
	done := make(chan struct{}, workers)

	go addJobs(jobs, filenames, results)	//在自己的goroutine中执行
	for i := 0; i < workers; i++ {
		go doJobs(done, linePx, jobs)	//每一个都在自己的goroutine中执行
	}
	go awaitCompletion(done, results)	//在自己的gooutine中执行
	processResult(results)		//阻塞，直到工作完成
}