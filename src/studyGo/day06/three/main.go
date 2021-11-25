package main

import "fmt"

//函数进阶
var x = 10 //全局变量x
//函数中查找变量的顺序，1.先在函数内部查找 2.在函数外面查找，一直找到全局
func f1() {
	//x := 10
	name := "xxx"
	//函数中查找变量的顺序
	//1.先在函数内部查找,找到就不找了
	//2.在函数外面查找，一直找到全局
	fmt.Println(name, x)
}

func f2() {
	fmt.Println("hello")
}

func f3() int {
	return 2
}

//函数也可以作为参数的类型
func f4(x func()) {
	x()
}

//函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	// f1()
	//fmt.Println(name) //函数内部定义的变量只能在函数内部使用
	// if i := 10; i < 19 {
	// 	fmt.Println("study")
	// }
	// 	fmt.Println(i)//i只能在该语句块中使用

	// for i := 0; i < 2; i++ {
	// 	fmt.Println(i)
	// }
	// 	fmt.Println(i)//i只能在该语句块中使用

	// a := f2
	// b := f3
	// fmt.Printf("aType:%T\n", a)
	// fmt.Printf("bType:%T\n", b)
	f4(f2)
}
