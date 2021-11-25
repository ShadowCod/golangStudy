package main

import (
	"fmt"
	"net/http"
)

// 创建cookie
func handler(w http.ResponseWriter, r *http.Request) {
	// 方式一:
	// 创建cookie
	cookie := http.Cookie{
		Name:  "user",
		Value: "admin",
	}
	cookie2 := http.Cookie{
		Name:  "user2",
		Value: "admin2",
	}
	// 将cookie发送给客户端
	w.Header().Set("Set-Cookie", cookie.String())
	w.Header().Add("Set-Cookie", cookie2.String())
	// 使用http的方法将cookie发送给客户端
	// 方式二:
	// http.SetCookie(w, &http.Cookie{
	// 	Name:  "user",
	// 	Value: "golang",
	// })
}

// Handler 获取cookie的方式
func Handler(w http.ResponseWriter, r *http.Request) {
	// 方法一:
	// 获取请求头中所有的cookie
	// cookie := r.Header["Cookie"]
	// 方法二:
	// 获得指定的cookie
	cookie, _ := r.Cookie("user")
	fmt.Println("得到的cookie", cookie)
}
func main() {
	http.HandleFunc("/setCookie", handler)
	http.HandleFunc("/getCookie", Handler)
	http.ListenAndServe(":8080", nil)
}
