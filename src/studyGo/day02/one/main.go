package main

import "fmt"

func main() {
	// math.MaxFloat32()
	f1 := 1.23456
	fmt.Printf("%T\n", f1)
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) //显示声明float32类型(不能将float32类型的值赋值给float64)

	b1 := true
	var b2 bool //默认是false
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value:%v\n", b2, b2)
	test()
}

//fmt占位符
func test() {
	n := 100
	fmt.Printf("%T\n", n) //类型
	fmt.Printf("%v\n", n) // 值
	fmt.Printf("%b\n", n) //二进制
	fmt.Printf("%d\n", n) //十进制
	fmt.Printf("%o\n", n) //八进制
	fmt.Printf("%x\n", n) //十六进制
	var s = "长江"
	fmt.Printf("%s\n", s) //字符串
}
