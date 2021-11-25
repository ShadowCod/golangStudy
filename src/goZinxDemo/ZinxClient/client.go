package main

import (
	"fmt"
	"net"
	"time"
)

/*
	模拟客户端
*/
func main() {
	// 获取连接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()
	// 发送和读写信息
	for {
		_, err := conn.Write([]byte("hello Zinx..."))
		if err != nil {
			fmt.Println("client write err:", err)
			return
		}
		readInfo := make([]byte, 512)
		_, err = conn.Read(readInfo)
		if err != nil {
			fmt.Println("client read err:", err)
			return
		}
		fmt.Println("read info:", readInfo)
		time.Sleep(1 * time.Second)
	}
}
