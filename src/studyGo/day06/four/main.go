package main

import "fmt"

//匿名函数
// var f1 = func(x, y int) {
// 	fmt.Println(x + y)
// }

func main() {
	//函数内部没法声明带名字的函数
	// var f1 = func(x, y int) {
	// 	fmt.Println(x + y)
	// }
	// f1(1, 2)

	//如果只是调用一次，可以写出立即调用的函数
	func(a, b int) {
		fmt.Println(a + b)
		fmt.Println("hello")
	}(100, 200)
}
