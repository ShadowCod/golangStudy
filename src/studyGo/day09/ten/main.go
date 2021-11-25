package main

import "fmt"

// map的增删改查
func main() {
	// 增加和更新：map["key"] = "value",没有key就是增加，有就是修改
	cities := make(map[string]string)
	cities["no1"] = "北京"
	fmt.Println(cities)
	cities["no1"] = "上海" //修改
	fmt.Println(cities)
	cities["no2"] = "广州" //增加
	fmt.Println(cities)

	// 删除：delete(map,"key")	key不存在时不会操作也不会报错
	// delete(cities, "no1")
	// fmt.Println(cities)
	// delete(cities, "no3")
	// fmt.Println(cities)

	// 一次性删除所有的key，①遍历逐一删除 ②make一个新的空间
	// cities = make(map[string]string)
	// fmt.Println(cities)

	//查询：value,findRes := cities["no1"]	存在findRes为true
	_, ok := cities["no1"]
	if ok {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}
