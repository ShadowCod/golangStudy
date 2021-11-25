package main

import "fmt"

/*
快速排序
*/
// 说明：left表示左边的下标，right表示右边的下标
func QuickSort(left int, right int, arr *[11]int) {
	l := left
	r := right
	// pivot：中轴，支点
	pivot := arr[(left+right)/2]
	//比pivot大的数放到右边
	for l < r {
		// 从pivot的左边找到大于pivot的值
		for arr[l] < pivot {
			l++
		}
		// 从piv的右边找到小于pivot的值
		for arr[r] > pivot {
			r--
		}
		// l>=r表明本次分解任务完成
		if l >= r {
			break
		}
		// 交换
		arr[l], arr[r] = arr[r], arr[l]
		// 优化
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	// fmt.Println(arr)
	// 这一步必须加，不然会死循环
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuickSort(left, r, arr)
	}
	// 向又递归
	if right > l {
		QuickSort(l, right, arr)
	}
}

func quickSort(left, right int, arr *[11]int) {
	// 分别将left、right赋值给一个用来进行移动的变量
	l := left
	r := right
	// 通过left和right取一个支点的元素
	pivot := arr[(l+r)/2]
	// 通过for循环将小于和大于支点的值分到不同的一边
	for l < r {
		// 将小于支点的和大于支点的索引的值进行交换
		// 使用for循环找到不满足条件的下标
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		// 当两个下标移动到同一个位置则本次分解完成
		if l >= r {
			break
		}
		// 交换两个下标索引的值
		arr[l], arr[r] = arr[r], arr[l]
		// 优化
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	fmt.Println(arr)
	// 这一步必须加，不然会死循环
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		quickSort(left, r, arr)
	}
	// 向右递归
	if right > l {
		quickSort(l, right, arr)
	}
}

func main() {
	arr := [11]int{-9, 78, 0, 23, -567, 70, 100, 1, -20, 40, -1}
	fmt.Println(arr)
	// QuickSort(0, len(arr)-1, &arr)
	quickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}
