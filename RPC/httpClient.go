//HTTP的Rpc
//调用服务
package main

import(
	"fmt"
	"log"
	"net/rpc"
	"os"
)

//传入的类型
type Args struct {
	A, B int
}

type Quotient struct {	//商的类型
	Quo, Rem int	//商  余数
}

func main() {
	//通过命令行参数取得远程参数的地址
	if len(os.Args) != 2 {	//命令行参数不为2个
		fmt.Println("Usage:", os.Args[0], "server")
		os.Exit(1)
	}
	serverAdd := os.Args[1]
	client, err := rpc.DialHTTP("tcp", serverAdd + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//调用
	args := Args{17, 8}

	//乘法
	var reply int
	err =  client.Call("Math.Multiply", args, &reply)
	if err != nil {
		log.Fatal("math error:", err)
	}

	fmt.Printf("Math %d*%d=%d\n", args.A, args.B ,reply)

	//除法
	var quo Quotient
	err = client.Call("Math.Divide", args, &quo)
	if err != nil {
		log.Fatal("math error:", err)
	}
	fmt.Printf("Math %d/%d=%d remainder %d\n", args.A, args.B , quo.Quo, quo.Rem)
}