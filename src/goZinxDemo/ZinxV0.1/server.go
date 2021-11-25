package main

import "goZinx/znet"

func main() {
	// 创建一个启动句柄
	s := znet.CreateServer("[zinxv0.1]")
	// 启动
	s.Serve()
}
