package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	monster := Monster{"红孩儿", 10, "喷火"}
	// 将monster变量序列化为json格式
	jsonStr, _ := json.Marshal(monster)
	fmt.Println(string(jsonStr))
}
