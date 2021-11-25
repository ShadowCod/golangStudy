package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
* 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
* 你可以按任意顺序返回答案。
 */

func main() {
	//生成一个整数数组
	// array := createArray(5)
	// fmt.Println(array)
	var array = []int{2, 7, 11, 4, 5}
	target := 9
	// a(array, target)
	b(array, target)

}

func createArray(length int) (array []int) {
	rand.Seed(time.Now().Unix()) //设定种子时不能放在for循环内，放进去每个随机值就相同了
	for i := 0; i < length; i++ {
		array = append(array, rand.Intn(100000))
	}
	return
}

// 暴力枚举
func a(array []int, target int) {
	r := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i]+array[j] == target {
				if !sliceContainElement(r, i) {
					r = append(r, i)
				}
				if !sliceContainElement(r, j) {
					r = append(r, j)
				}
			}
		}
	}
	fmt.Println(r)
}

// 哈希表
func b(array []int, target int) {
	// 先声明一个map类型存储每个值的索引
	m := map[int]int{}
	for k, v := range array {
		if p, ok := m[target-v]; ok {
			fmt.Println(k, p)
			delete(m, target-v)
		}
		m[v] = k
	}
}

func sliceContainElement(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// 变式：求3数、4数之和
