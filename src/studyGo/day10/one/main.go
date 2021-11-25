package main

import "fmt"

func main() {
	// map的遍历只能使用for range 形式，不能使用for
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "天津"
	fmt.Println(cities)

	// 遍历
	for k, v := range cities {
		fmt.Printf("key:%v value:%v\n", k, v)
	}

	// 获取map的长度
	fmt.Println(len(cities))
}
