package main

import "fmt"

// 编写一个函数，完成找路功能
// myMap:地图，要保证是同一个地图，使用引用
// i,j：表示出发点
func FindWay(myMap *[8][7]int, i, j int) bool {
	// 分析出什么情况下就找到了出口
	// myMap[6][5] == 2
	if myMap[6][5] == 2 { //出口等于2则说明找到了出口，退出该函数
		return true
	} else {
		// 说明需要继续找
		// 注意：这里会出现两种情况，要么是没有走过的、要么是墙
		// 1、判断现在所在的位置是不是没有走过的
		if myMap[i][j] == 0 {
			// 假设这个点是通的，然后才能进行反推
			myMap[i][j] = 2
			if FindWay(myMap, i+1, j) { //向下探测
				return true
			} else if FindWay(myMap, i, j-1) { //向左探测
				return true
			} else if FindWay(myMap, i, j+1) { //向右探测
				return true
			} else if FindWay(myMap, i-1, j) { //向上探测
				return true
			} else { //死路
				myMap[i][j] = 3
				return false
			}

		} else { //说明这个点并不能探测，是墙
			return false
		}

	}
}

func main() {
	// 创建二维数组，模拟迷宫
	// 规则
	// 1.如果元素的值为1，就是墙
	// 2.如果元素的值为0，是没有走过的点
	// 3.如果元素的值为2，是一个通路
	// 4.如果元素的值为3，是走过的点，但是走不通

	var myMap [8][7]int
	// 生成迷宫墙壁(行)
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	// 生成迷宫墙壁(列)
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	// 打印迷宫样式
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
	FindWay(&myMap, 1, 1)
	fmt.Println("pass...")
	// 打印迷宫样式
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
