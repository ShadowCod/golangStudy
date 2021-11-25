package main

import "fmt"

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1) //替换为十进制
	// fmt.Printf("%b\n", i1) //替换为二进制
	// fmt.Printf("%o\n", i1) //替换为八进制
	// fmt.Printf("%x\n", i1) //替换为十六进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	i3 := 0xeee
	fmt.Printf("%d\n", i3)
	//查看变量的类型
	fmt.Printf("%T\n", i3)
	//声明指定类型变量
	i4 := int8(9)
	fmt.Printf("%T\n", i4)
}
