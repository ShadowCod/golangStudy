package main

import "fmt"

// 内置函数
func main() {
	// 1) len:用来求长度	string、array、slice、map、channel
	fmt.Println(len("1234"))
	// 2) new:用来分配内存，主要用来分配值类型，比如int、float32、struct...返回指针
	var num = new(int)
	fmt.Printf("num类型:%T,num值:%v,num内存地址:%v\n", num, num, num)
	// num本身存储的是一个内存地址，改地址指向的值为0
	fmt.Println(*num)
	// 将0改为100
	*num = 100
	fmt.Println(*num)

	// 3) make:用来分配内存，主要是用来分配引用类型，比如channel、map、slice
}
