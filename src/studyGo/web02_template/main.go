package main

import (
	"net/http"
	"text/template"
)

// handler 生成处理器的函数
func handler(w http.ResponseWriter, r *http.Request) {
	// 解析模版
	// t, _ := template.ParseFiles("index.html")
	// 通过must函数自动处理异常
	t := template.Must(template.ParseFiles("index.html"))
	// 执行
	t.Execute(w, "Hello")
	// 将响应数据放在其他文件中显示
	t.ExecuteTemplate(w, "index2.html", "响应在index2中的数据")
}
func main() {
	http.HandleFunc("/Template", handler)
	http.ListenAndServe(":8080", nil)
}
