package main

import "fmt"

type Person struct {
	Name string
}

// 函数
func test(p Person) {
	fmt.Println("函数中的~~~", p.Name)
}

// 方法
func (p Person) test01() { //实参的形式看方法中声明的样式
	fmt.Println("方法中的~~~", p.Name)
}

func main() {
	// 调用方式不一样
	// 函数调用方式：函数名（实参列表）
	// 方法调用方式：变量.方法名（实参列表）

	// 对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
	var p Person
	p.Name = "tom"
	test(p)
	// test(&p) //函数只能传入函数实参类型一致的形参

	// 对于方法，接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样
	p.test01()
	(&p).test01() //使用这个方法调用方法，方法中的实参也是拷贝的，并不是地址

	// 总结：不管调用形式如何，真正决定是值拷贝还是地址拷贝，看这个方式是和什么类型绑定的，如果是和值类型绑定，则是值拷贝
}
