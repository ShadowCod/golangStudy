package main

import (
	"fmt"
	"strings"
)

func addUpper() func(int) int {
	var n = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

// 闭包相关
func main() {
	f := addUpper()
	// res := addUpper()(1) //这种调用方式n不会被改变
	// res2 := addUpper()(2)
	fmt.Println(f(1))
	fmt.Println(f(2)) //这种调用方式n的值会被改变
	/*
		说明和总结
		1)addUpper是一个函数，返回的数据类型是一个func(int) int
		2)闭包的说明：返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数和n就形成一个整体，构成闭包
		3）f这个闭包包含了n和匿名函数，n只初始化了一次
	*/

	// 闭包的练习

}

// 闭包的练习
func makeSuffix(suffix string) func(string) string {
	return func(file string) string {
		if strings.HasSuffix(file, suffix) {
			return file
		}
		return file + suffix
	}
}
