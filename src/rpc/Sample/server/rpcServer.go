package main

import (
	"fmt"
	"net"
	"net/rpc"
)

/*
* rpc:远程过程调用协议
* 本地通过传入函数名、参数，达到想访问本地函数一样去访问远程的函数并得到结果
 */
//  S定义结构体
type S struct {
}

// Say 绑定结构体方法
func (s *S) Say(str string, r *string) error {
	*r = fmt.Sprintf("say:%s", str)
	return nil
}
func main() {
	err := rpc.RegisterName("frist_rpc", new(S))
	if err != nil {
		fmt.Println(err)
		return
	}
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("=====>等待连接")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	rpc.ServeConn(conn)
}
