package main

import "fmt"

/*
插入排序:将最前面的一个元素看成一个有序数组，后面的值依次和有序数组中的值比较找到正确的位置（每一次下标的移动都要随着数据的移动进行）
*/
func InsertSort(arr *[5]int) {
	// //完成第一次，给第二个元素找到合适的位置并插入
	// insertVal := arr[1]  //将第一个值当成一个有序数组，因此取第二个值
	// insertIndex := 1 - 1 //下标总是元素的前一个值
	// // 从大到小（下标往前移动）
	// for insertIndex >= 0 && arr[insertIndex] < insertVal {
	// 	arr[insertIndex+1] = arr[insertIndex] //数据后移
	// 	insertIndex--
	// }
	// // 插入
	// if insertIndex+1 != 1 {
	// 	arr[insertIndex+1] = insertVal
	// }
	// fmt.Println("第一次插入后:", arr)

	// //完成第二次，给第三个元素找到合适的位置并插入
	// insertVal = arr[2]  //将第一个值当成一个有序数组，因此取第二个值
	// insertIndex = 2 - 1 //下标总是元素的前一个值
	// // 从大到小（下标往前移动）
	// for insertIndex >= 0 && arr[insertIndex] < insertVal {
	// 	arr[insertIndex+1] = arr[insertIndex] //数据后移
	// 	insertIndex--
	// }
	// // 插入
	// if insertIndex+1 != 2 {
	// 	arr[insertIndex+1] = insertVal
	// }
	// fmt.Println("第二次插入后:", arr)

	// 将上面的重复代码使用for循环
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
	fmt.Println("插入排序后:", arr)
}

func main() {
	arr := [5]int{23, 0, 12, 56, 34}
	sort(&arr)
}

/*
插入排序：将数字第0索引的元素当成一个有序数组，从索引为1的元素开始比较；创建一个变量保存对应索引的值，创建一个变量表示插入的位置索引，该索引从1索引开始
*/
func sort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		// 注意：我们取值时是从索引1开始取值，但是比较的时候应该是从取值索引的前一个开始比较，因此需要i-1
		insertIndex := i - 1
		insertVal := arr[i]
		// 从大到小排列
		for {
			// 注意：判断属性是否后移，需要根据需求进行且insertInde必须>=0
			if insertIndex >= 0 && insertVal > arr[insertIndex] {
				arr[insertIndex+1] = arr[insertIndex]
				insertIndex--
			} else {
				break
			}
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Println("插入排序后:", arr)
	}

}
