package main

import "fmt"

//闭包是什么  闭包= 函数 + 外部变量的引用
// 闭包是一个函数，这个函数包含了它外部作用域的变量
// 底层原理：
// 1.函数可以作为返回值
// 2.函数内部查找变了的顺序，先在自己内部找，找不到往外层找
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定义一个函数对f2进行包装
func f3(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}

func main() {
	//将f2作为参数传给f1
	// f1(f3(f2, 2, 3))
	// ret := adder()
	// fmt.Println(ret(100))
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x = x + y
		return x
	}
}
