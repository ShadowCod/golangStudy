package main

import "fmt"

// 结构体

//定义cat结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	// 创建cat变量
	// 方式一
	// var cat1 Cat
	// fmt.Println(cat1)
	// cat1.Name = "小白"
	// fmt.Println(cat1)

	// 方式二
	// cat2 := Cat{"小白", 3, "白"}
	// cat2 := Cat{}
	// cat2.Name = "小白"
	// fmt.Println(cat2)

	// 方式三	var person *Person = new(Person)
	// cat3 := new(Cat)
	// cat3是一个指针
	// (*cat3).Name = "tom" 也可以这样写 cat3.Name = "tom"
	// 原因：go的设计者 为了使用方便，在底层对其进行处理，会给cat3加上取值运算
	// (*cat3).Name = "tom"
	// cat3.Age = 11
	// fmt.Println(*cat3)

	// 方式四	cat4是一个指针
	var cat4 = &Cat{}
	// var cat4 = &Cat{"may",11}
	(*cat4).Age = 11
	cat4.Name = "may"
	fmt.Println(*cat4)
}
