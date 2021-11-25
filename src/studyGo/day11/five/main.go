package main

import "fmt"

// go语言中的继承-----嵌套一个匿名结构体
type A struct {
	Name string
	age  int
}

func (a *A) sayOk() {
	fmt.Println("sayOk~")
}

func (a *A) Hello() {
	fmt.Println("Hello~")
}

type B struct {
	A
}

func main() {
	var b B
	b.A.Name = "tom" //因为系统会帮忙找所以也可以写成b.Name = "tom"
	fmt.Println(b)
	b.Hello()
	b.sayOk()

}

// 总结：1）编译器会先看b对应的类型有没有Name，如果有就直接调用B类型的Name字段，如果没有就去看B类型中嵌套的匿名结构体A中有没有Name，没有还没有就去找A类型中的匿名结构体，没找到就报错（方法和属性都会就近原则）
