package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求地址是:", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求参数是:", r.URL.RawQuery)
	fmt.Fprintln(w, "head消息:", r.Header)
}

func main() {
	// 注册处理器
	http.HandleFunc("/", handler)
	// 设置监听
	http.ListenAndServe(":8080", nil)
}
