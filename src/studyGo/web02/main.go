package main

import (
	"fmt"
	"net/http"
)

/*
	使用http.Handle和自定义的多路复用器
*/
// 自定义的处理器结构体
type MyHandler struct{}

// 实现serveHTTP
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("实现接口方法")
}

func main() {
	// myHandler := &MyHandler{}
	// 映射处理器
	// http.Handle("/myHandler", myHandler)
	// 监听端口
	// http.ListenAndServe(":8080", nil)

	// 使用server来设置配置信息
	// server := &http.Server{
	// 	Addr:        ":8080",
	// 	Handler:     myHandler,
	// 	ReadTimeout: 2 * time.Second,
	// }
	// server.ListenAndServe()

	var num1 uint = 3
	var num2 uint = 2
	fmt.Println(num2 - num1)
}
