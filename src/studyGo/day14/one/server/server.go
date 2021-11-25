package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 循环接收客户端接收的数据
	defer conn.Close() //关闭连接
	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		//等待客户端通过conn发送信息，如果客户端没有发送，就会阻塞
		fmt.Println("等待输入", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("server read err", err)
			return
		}
		// 显示
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("server listen~~~")
	// 1.tcp边上使用的网络的协议是tcp
	// 2.9999表示监听的端口
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listen.Close() //延时关闭
	// 循环等待客户端连接
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err", err)
		} else {
			fmt.Printf("client ip:%v\n", conn.RemoteAddr().String())
			go process(conn)
		}
	}
}
