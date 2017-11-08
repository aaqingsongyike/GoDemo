//HTTP的Rpc

//提供服务
package main

import(
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
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
	rpc.HandleHTTP()

	//启动服务
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}