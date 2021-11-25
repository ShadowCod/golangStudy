package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义变量，用于接收参数值
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机,默认为空")
	flag.IntVar(&port, "port", 3306, "端口,默认为3306")

	// 这里有一个非常重要的操作，必须使用该方法
	flag.Parse()

	// 输出
	fmt.Println(user, pwd, host, port)
}
