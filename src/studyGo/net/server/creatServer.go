package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Handle 用于生成处理器的函数参数必须为http.ResponseWriter、*http.Request
func Handle(w http.ResponseWriter, r *http.Request) {
	// 可以从r中取出访问请求中的值、w是服务器想要对该请求给予的响应
	// _, err := fmt.Fprintln(w, "hello golang")
	// if err != nil {
	// 	fmt.Println("写入响应失败:", err)
	// 	return
	// }

	/*
		给响应中写入内容也可以使用文件的格式：创建一个文件专门写响应的内容，在处理函数中读取内容并写入响应中
	*/

	// 读取文件中的内容
	bytes, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("读取文件内容失败:", err)
		return
	}
	// fmt.Println(string(bytes))
	_, err = fmt.Fprintln(w, string(bytes))
	if err != nil {
		fmt.Println("写入响应失败:", err)
		return
	}
}

func main() {
	// 将处理器函数转换为正真的处理器,并绑定地址
	http.HandleFunc("/hello", Handle)
	// 启动服务器监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("监听端口失败", err)
		return
	}
}
