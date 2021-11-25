package main

import "fmt"

//GO语言中不存在指针操作，只需要记住两个符号：
// 1.&：取地址符号
// 2.*: 根据地址取值
func main() {
	n := 18
	fmt.Println("内存地址:", &n)
	p := &n
	fmt.Printf("%T\n", p) //*int是int类型的指针
	fmt.Println(*p)

	//new函数申请一个内存地址
	// var a *int //nil  错误操作
	var b = new(int)
	fmt.Println(b)
	*b = 100
	fmt.Println(*b)

	//make和new 都是用来申请内存的，new很少用，一般用于给基本数据类型申请内存，string|int返回对应的指针
	// make用来给slice|map|chan申请内存的，返回对应的三个类型指针
}
