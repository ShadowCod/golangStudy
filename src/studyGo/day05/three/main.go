package main

import "fmt"

//map和slice组合
func main() {
	// 元素类型为map的切片
	// var s1 = make([]map[int]int, 2, 10)
	// s1[0][1] = 1 //make([]map[int]int, 0, 10)会抛错，因为申请内存时长度为0

	// s1[0][1] = 1//没有对内部的map做初始化
	// s1[0] = make(map[int]int, 1)
	// s1[0][1] = 1

	// fmt.Println(s1)

	//值为切片类型的map
	var m1 = make(map[int][]string, 5)
	fmt.Println(m1)
	m1[1] = []string{"北京", "上海"}
	fmt.Println(m1)
}
