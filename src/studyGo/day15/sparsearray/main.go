package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type varNote struct {
	row int //行
	col int //列
	val int //值
}

func main() {
	// 1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //蓝子
	// 2.查看原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d", v2)
		}
		fmt.Println()
	}
	// 3.转成稀疏数组
	// 思路:(1).遍历chessMap，发现一个不为0的就创建一个对应的结构体(2).将结构体放入对应的切片中
	var slice []varNote
	// 标准的稀疏数组应该还含有数组的大小（行和列，默认值）
	ori := varNote{
		row: 11,
		col: 11,
		val: 0,
	}
	slice = append(slice, ori)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				vn := varNote{
					row: i,
					col: j,
					val: v2,
				}
				slice = append(slice, vn)
			}
		}
	}
	// 输出稀疏数组
	fmt.Println(slice)
	// // 将稀疏数组存入文件|磁盘
	// Save(slice)
	// 从文件|磁盘中读取并复原
	Revert()
}

func Save(s []varNote) {
	fileStr := "f:/abc.txt"
	file, err := os.OpenFile(fileStr, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件出错", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range s {
		str := fmt.Sprintf("%d\t%d\t%d\r\n", v.row, v.col, v.val)
		_, err = writer.WriteString(str)
		if err != nil {
			fmt.Println("写入文件出错", err)
			return
		}
	}
	writer.Flush()
}

func Revert() {
	fileStr := "f:/abc.txt"
	file, err := os.OpenFile(fileStr, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("打开文件出错", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	// 将读取到的数据还原为节点切片
	var slice []varNote
	for {
		str, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		strS := strings.Split(strings.Replace(str, "\r\n", "", -1), "\t")
		vn := varNote{}
		for i, v := range strS {
			val, _ := strconv.Atoi(v)
			if i == 0 {
				vn.row = val
			} else if i == 1 {
				vn.col = val
			} else {
				vn.val = val
			}
		}
		// vn := varNote{
		// 	row: strS[0],
		// 	col: strconv.Atoi(strS[1], 10, 64),
		// 	val: strconv.Atoi(strS[2]),
		// }
		// 还原数组
		slice = append(slice, vn)
	}
	// fmt.Println(slice)
	// 遍历slice还原棋盘
	var chessMap [11][11]int //应该行和列是变量（动态的，因此不能这样定义）
	for i, v := range slice {
		if i != 0 {
			chessMap[v.row][v.col] = v.val
		}
	}
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d", v2)
		}
		fmt.Println()
	}
}
