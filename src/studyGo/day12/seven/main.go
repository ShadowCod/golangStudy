package main

import (
	"encoding/json"
	"fmt"
)

// json的反序列化

// 反序列化结构体
type Monster struct {
	Name     string  `json:"name"` //反射机制
	Age      int     `json:"age"`
	Brithday string  `json:"brithday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func UnjsonStruct() {
	str := "{\"name\":\"悟空\",\"age\":500,\"brithday\":\"001-1-1\",\"sal\":10000,\"skill\":\"不坏金身\"}"
	// 执行json中的反序列化函数
	var m1 Monster
	err := json.Unmarshal([]byte(str), &m1) //第一个参数应该是[]byte
	if err != nil {
		fmt.Println("反序列化错误,", err)
		return
	}
	fmt.Println(m1)
}

func UnjsonSlice() {
	str := "[{\"agr\":18,\"name\":\"tom\"},{\"agr\":20,\"name\":\"jack\"}]"
	var s []map[string]interface{}
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println("反序列化错误,", err)
		return
	}
	fmt.Println(s)
}

func UnjsonMap() {
	str := "{\"address\":\"花果山\",\"age\":500,\"name\":\"悟空\"}"
	var m map[string]interface{}
	//反序列化中不需要make，make操作被封装到了Unmarshal函数中
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("反序列化错误,", err)
		return
	}
	fmt.Println(m)
}

func main() {
	// UnjsonStruct()
	UnjsonSlice()
	// UnjsonMap()
}
