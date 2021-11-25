package main

import "fmt"

//关于append
func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s1 := a1[:]

	//删除索引为1的那个3
	s1 = append(s1[:1], s1[2:]...)
	fmt.Printf("切片值:%v\n", s1)
	fmt.Printf("数组值:%v\n", a1)
}
