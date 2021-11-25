package main

import "fmt"

func main() {
	// append,对切片进行追加
	var slice = []int{100, 200, 300}
	fmt.Println("slice", slice)
	// 通过append给slice追加具体元素（元素类型要一致）
	slice1 := append(slice, 400)
	fmt.Println("slice1", slice1) //原数组的值不会改变
	fmt.Printf("slice的值:%v\n", slice)
	// 通过append追加切片,后面的值只能说切片，不能是数组且...是固定写法
	slice2 := append(slice, slice...)
	fmt.Println("slice2", slice2)

	// 切片的拷贝,两个数据空间是独立的，修改其中一个不会影响另一个
	slice3 := make([]int, 5)
	copy(slice3, slice)
	fmt.Println(slice3)

}
