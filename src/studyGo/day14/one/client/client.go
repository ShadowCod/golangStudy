package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("client ...")
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("dial err", err)
		return
	}
	//获取终端输入
	reader := bufio.NewReader(os.Stdin) //标准输入[终端]
	for {
		//从终端读数据
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reader.ReadString err", err)
		}
		// 判断是否要退出
		str1 := strings.Trim(str, " \r\n")
		if str1 == "exit" {
			break
		}
		// 将读取的数据发给服务器
		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Println("conn.write err", err)
		}
	}
	fmt.Println("client close")
}
