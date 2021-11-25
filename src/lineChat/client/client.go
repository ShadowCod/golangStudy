package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	// 监听连接返回的消息
	go serverBackInfo(conn)
	var info string
	// 提示输入消息
	for {
		fmt.Println("请输入message:")
		fmt.Scanln(&info)
		// 向服务器发送消息
		buf := []byte(info)
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println("conn.Write err:", err)
			continue
		}
	}
}

func serverBackInfo(c net.Conn) {
	// 展示服务端返回的消息
	info := make([]byte, 512)
	n, err := c.Read(info)
	if err != nil {
		fmt.Println("返回信息读取错误")
		return
	}
	fmt.Println("服务端返回的消息:", string(info[:n]))
}
