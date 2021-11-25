package main

import "fmt"

// 切片：是数组的一个引用，因此是引用类型，长度是可以改变的
func main() {
	// 方式一
	// var intArr = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(intArr)
	// slice := intArr[1:3]
	// fmt.Println(slice)
	// fmt.Printf("intArr[1]地址:%p,值:%v\n", &intArr[1], intArr[1])
	// fmt.Printf("slice[0]地址:%p,值:%v\n", &slice[0], slice[0])
	// fmt.Printf("slice地址:%p\n", &slice)

	// slice[1] = 7
	// fmt.Println("修改后的数组", intArr)
	// fmt.Println("修改后的切片", slice)

	// 方式二 使用make
	// var slice = make([]float64, 4)
	// slice[1] = 2
	// fmt.Println(slice)
	// fmt.Printf("len:%v cap:%v\n", len(slice), cap(slice))

	//通过make方式创建切片，可以指定切片的大小和容量
	// 没有赋值则是零值
	// 通过make方式创建的切片对应的数组由make底层维护对外不可见

	// 方式三
	var slice = []int{1, 2, 3}
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
}
