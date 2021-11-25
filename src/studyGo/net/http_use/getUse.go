package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	GetFunc("https://www.liwenzhou.com/")
}

// GetFunc 使用get访问url(不带参数)
func GetFunc(url string) {
	// 使用"net/http"包中的Get函数发送访问请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("访问失败:", err)
		// 退出函数
		return
	}
	// 接收到的响应在使用完毕后需要关闭（注意）
	defer resp.Body.Close()
	// 读取响应中的内容（返回的body为[]byte类型）
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		// 退出函数
		return
	}
	// 打印响应体中的内容
	fmt.Println("响应体中的内容:", string(body))
}

// GetParamFunc 使用get访问url(带参数)
func GetParamFunc() {
	apiURL := "127.0.0.1:9090/hello"
	// url参数
	data := url.Values{}
	data.Set("name", "小王子")
	data.Set("age", "18")
	// 解析url地址
	u, err := url.ParseRequestURI(apiURL)
	if err != nil {
		fmt.Println("解析url失败:", err)
		return
	}
	// 将url参数编码为url编码格式
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	// 向地址发起请求
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("访问失败:", err)
		// 退出函数
		return
	}
	// 接收到的响应在使用完毕后需要关闭（注意）
	defer resp.Body.Close()
	// 读取响应中的内容（返回的body为[]byte类型）
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		// 退出函数
		return
	}
	// 打印响应体中的内容
	fmt.Println("响应体中的内容:", string(body))
}
