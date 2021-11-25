package main

import "fmt"

//切片Slice：（引用类型）与数组比较不需要规定长度
func main() {
	//切片的声明
	// var name []T   name:表示变量名   T:表示切片中的元素类型
	var str []string
	fmt.Println("声明不初始~~", str)
	var a = []int{}
	fmt.Println("声明并初始化为空:", a)
	var b = []bool{true}
	fmt.Println("声明并初始化:", b)
	//判断一个切片是否被初始就要nil进行比较
	fmt.Println(str == nil)
	fmt.Println(a == nil)
	fmt.Println(len(b), cap(b))
	//数组也可以用len(),cap()
	var c = [...]int{1, 2, 3}
	fmt.Println(len(c))

	//由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5}
	s1 := a1[0:4] //基于一个数组切割，左闭右开 a1[:4]   a1[:]   a1[2:]
	fmt.Println(s1)
	s2 := a1[2:]
	fmt.Println(s2)
	//切片的容量是值底层数组的容量（不完整）
	fmt.Printf("len:%d cap:%d\n", len(s1), cap(s1))
	//切片的长度就是它元素的个数，切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
	fmt.Printf("len:%d cap:%d\n", len(s2), cap(s2))

	//切片再切片   切片是一个引用类型，都指向底层的一个数组
	s3 := s1[1:3]
	fmt.Printf("len:%d cap:%d\n", len(s3), cap(s3))
}
