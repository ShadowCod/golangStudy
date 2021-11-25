package main

import "fmt"

// 变量的声明
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string //""
	age  int    // 0
	isOk bool   //false
)

func main() {
	name = "张三"
	age = 12
	isOk = true
	// var s1 sting = "李四"
	// var s2 = "王五"
	// s3 := "呵呵"  只能在函数中使用
	// name = "王五"
	fmt.Printf("name:%s", name)
	fmt.Println(age)
	fmt.Print(isOk)
}

// 常量(定义后不能修改)
const (
	statusOk = 200
	notFound = 404
)

//批量声明常量,如果某一行声明后没有赋值,默认取上一行的值
const (
	n1 = 100
	n2
	n3
)

//iota 出现时就为0,每增一行声明就+1	(每个const中出现iota时都为0)
const (
	a1 = iota
	a2 = 0
	a3
)

func test1() {
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
}
