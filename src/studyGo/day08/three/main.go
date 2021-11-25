package main

import "fmt"

var func1 = func(n1, n2 int) int {
	return n1 + n2
}

// 函数练习
func sum(n1, n2 float32) float32 {
	return n1 + n2
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}

func main() {
	// var n1 float32 = 1
	// var n2 float32 = 2
	// fmt.Println(sum(n1, n2))

	// n1 := 10
	// n2 := 20
	// swap(&n1, &n2)
	// fmt.Println(n1, n2)

	// 匿名函数
	// 在定义你们函数时就直接调用，这种方式只能调用一次
	res := func(n1, n2 int) int {
		return n1 + n2
	}(1, 2)
	fmt.Println(res)

	// 将匿名函数赋值给一个变量，通过变量进去调用
	// a的数据类型就是函数类型
	a := func(n1, n2 int) int { return n1 + n2 }
	fmt.Println(a(1, 4))

	// 全局匿名函数，就是将一个匿名函数交给一个全局变量
	fmt.Println(func1(1, 5))
}
