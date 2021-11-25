package main

import "fmt"

// 冒泡排序
func BubbleSort(arr *[5]int64) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
	fmt.Println(*arr)
}

func main() {
	var arr = [5]int64{13, 3, 5, 20, 6}
	BubbleSort(&arr)
}

// 二分查找
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("没有")
		return
	}

	middle := (leftIndex + rightIndex) / 2
	if findVal < middle {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if findVal > middle {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Println("index", middle)
	}
}
