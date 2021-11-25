package main

import "fmt"

//函数是一段代码的封装，使用函数能使代码结构更清晰，更简洁
//函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

//有参数，没有返回值
func f1(x int, y int) {
	fmt.Println(x, y)
}

//没有参数，没有返回值
func f2() {
	fmt.Println("f2")
}

//没有参数，有返回值
func f3() int {
	return 3
}

//命名返回值，就相当于在函数中声明一个变量
func f4() (ret int) {
	ret = 1
	return //使用命名返回值可以return后省略
}

//多个返回值(接收也必须两个)
func f5() (int, string) {
	return 1, "2"
}

//参数类型简写
func f6(x, y int, a, b string) int {
	return x + y
}

//可变参数（y可以不传），可变长参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //类型是slice
}

//GO语言中函数没有默认参数这个概念

func main() {
	r := sum(1, 2)
	fmt.Println(r)
	f7("1", 2, 3, 4, 5)
}
