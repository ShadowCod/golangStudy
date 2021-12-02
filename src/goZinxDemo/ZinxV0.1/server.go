package main

import (
	"fmt"
	"goZinx/ziface"
	"goZinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("PreHandle")
	_, err := request.GetConnection().GetTCPConn().Write([]byte("PreHandle"))
	if err != nil {
		fmt.Println("PreHandle_err")
	}
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Handle")
	_, err := request.GetConnection().GetTCPConn().Write([]byte("Handle"))
	if err != nil {
		fmt.Println("Handle_err")
	}
}

func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("PostHandle")
	_, err := request.GetConnection().GetTCPConn().Write([]byte("PostHandle"))
	if err != nil {
		fmt.Println("PostHandle_err")
	}
}

func main() {
	// 创建一个启动句柄
	s := znet.CreateServer("[zinxv0.1]")
	r := &PingRouter{}
	s.AddRouter(r)
	// 启动
	s.Serve()
}
