package main

import "fmt"

/**
* 选择排序：循环找到最小或者最大的值，然后交换位置
* 优化 每次找到最小和最大值，同时将两个交换，可以省去一半的循环
 */
//  方式一
// func main() {
// 	var array = [5]int{5, 3, 6, 1, 2}
// 	// 方式一：很费（优化点：因为J=i+1,所以i值最大可以为len-1）
// 	for i := 0; i < len(array); i++ {
// 		minIndex := i
// 		for j := i + 1; j < len(array); j++ {
// 			if array[minIndex] > array[j] {
// 				minIndex = j
// 			}
// 			fmt.Println(minIndex)
// 		}
// 		// 找到最小值位置后交换
// 		array[i], array[minIndex] = array[minIndex], array[i]
// 		fmt.Println(array)
// 	}
// 	fmt.Println(array)
// }
// 优化方式
func main() {
	var list = [10]int{1, 3, 4, 2, 7, 5, 8, 6, 9}
	// 找到最大值和最小值
	for i := 0; i < len(list)/2; i++ {
		min, max := i, i
		for j := i + 1; j < len(list)-i; j++ {
			// 遍历找到最小值的索引
			if list[min] > list[j] {
				min = j
			}
			// 遍历找到最大值的索引
			if list[max] < list[j] {
				max = j
			}
		}
		fmt.Printf("第%d次遍历，最小值索引:%d,最大值索引:%d\n", i, min, max)
		// 交换位置
		list[i], list[min] = list[min], list[i]
		fmt.Println("交换最小值后的数组:", list)
		// 存在特殊情况，在最后一次时只需要一方交换位置就可以了
		if i != len(list)/2-1 {
			list[len(list)-i-1], list[max] = list[max], list[len(list)-i-1]
			fmt.Println("交换最大值后的数组:", list)
		}
	}
	fmt.Println("排序后的数组：", list)
}
