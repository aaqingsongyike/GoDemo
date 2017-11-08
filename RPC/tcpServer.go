//TCP的Rpc

//提供服务
package main

import(
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

//传入的类型
type Args struct {
	A, B int
}

type Math int

func (m *Math) Multiply (args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

type Quotient struct {	//商的类型
	Quo, Rem int	//商  余数
}
func (m *Math) Divide (args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	math := new(Math)
	//注册
	rpc.Register(math)

	tcpAdd, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		fmt.Println("Fatal error", err)
		os.Exit(2)
	}

	//监听地址
	listener, err := net.ListenTCP("tcp", tcpAdd)
	if err != nil {
		fmt.Println("Fatal error", err)
		os.Exit(2)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
			continue
		}
		rpc.ServeConn(conn)
	}
}