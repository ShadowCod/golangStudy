package main

import "fmt"

// map:无序的key-value数据结构		slice|map|function不能作为key
func main() {
	// var map 变量名 map[keyType]valueType		声明是不会分配内存的，初始化需要使用make，分配内存后才能赋值和使用
	// 方式一：
	// var a map[string]string
	// // a["1"] = "1" //未分配数据空间不能赋值
	// // fmt.Println(a)

	// // 在使用map前需要先make分配数据空间
	// a = make(map[string]string, 10)
	// a["no1"] = "北京"
	// fmt.Println(a)

	// 方式二:
	// b := make(map[string]string)
	// b["no1"] = "北京"
	// fmt.Println(b)

	// 方式三
	// c := map[string]string{
	// 	"no1": "北京",
	// 	"no2": "上海", //最后一个也必须又","
	// }
	// fmt.Println(c)

	// 练习
	studentMap := make(map[int]map[string]string)
	studentMap[1] = make(map[string]string) //此处make不能少
	studentMap[1]["name"] = "tom"
	fmt.Println(studentMap[1]["name"])
}
