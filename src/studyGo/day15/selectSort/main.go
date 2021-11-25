package main

import (
	"fmt"
	"time"
)

// 选择排序：选出最小的值和第0索引位交换，接着选出剩下中最小的值和第2索引位交换.....

// 编写函数完成排序，数组是值传递，要改变数组中的值需要使用指针
func SelectSort(arr *[5]int) {
	// 先完成将第一个最大值和arr[0]交换=>先易后难
	// 1.假设arr[0]是最大值
	// max := arr[0]
	// maxIndex := 0
	// // 2.遍历后面1->[len(arr)-1]的值
	// for i := 0 + 1; i < len(arr); i++ {
	// 	if max < arr[i] {
	// 		max = arr[i]
	// 		maxIndex = i
	// 	}
	// }
	// // 通过循环找到了最大的值
	// if maxIndex != 0 {
	// 	arr[0], arr[maxIndex] = arr[maxIndex], arr[0]
	// }
	// 因为每一步都是循环上面的代码
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
}

func main() {
	// 定义一个数组,从大到小排序
	arr := [5]int{10, 34, 100, 19, 80}
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Println(end - start)
	fmt.Println(arr)
	// 这种是平行赋值
	for a, b, c := 1, 2, 3; a < 5; a++ {
		fmt.Println(a, b, c)
	}
}
