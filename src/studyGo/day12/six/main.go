package main

import (
	"encoding/json"
	"fmt"
)

// json的序列化
type Monster struct {
	Name     string  `json:"name"` //反射机制
	Age      int     `json:"age"`
	Brithday string  `json:"brithday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func JsonStruct() {
	monster := Monster{
		Name:     "悟空",
		Age:      500,
		Brithday: "001-1-1",
		Sal:      10000.0,
		Skill:    "不坏金身",
	}
	// 将struct转换成Json
	data, err := json.Marshal(&monster) //返回的data是一个[]byte
	if err != nil {
		fmt.Println("序列化失败,", err)
		return
	}
	fmt.Println(string(data))
}

func JsonMap() {
	// 使用map要先make
	a := make(map[string]interface{})
	a["name"] = "悟空"
	a["age"] = 500
	a["address"] = "花果山"
	// 将map序列化成json
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败,", err)
		return
	}
	fmt.Println(string(data))
}

func JsonSlice() {
	var s []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "tom"
	m1["agr"] = 18
	s = append(s, m1)
	m2 := make(map[string]interface{})
	m2["name"] = "jack"
	m2["agr"] = 20
	s = append(s, m2)
	// 将slice序列化成json
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("序列化失败,", err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	// 将结构体、map、切片序列化
	// JsonStruct()
	JsonMap()
	// JsonSlice()
	// 对基本数据类型序列化成json意义不大
	// s := []int{1, 2, 3, 4, 5}
	// var s float64
	// s = 234.9
	// // 将slice序列化成json
	// data, err := json.Marshal(s)
	// if err != nil {
	// 	fmt.Println("序列化失败,", err)
	// 	return
	// }
	// fmt.Printf("序列化后类型:%T,值:%v", data, string(data))
}
