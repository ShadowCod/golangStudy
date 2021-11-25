package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 基础数据类型转String
	var num1 int = 99
	var num2 float64 = 23.12
	var b bool = true
	// var myChar byte = 'h'
	var str string

	//使用第一种方式来转换 fmt.Sprintf方法
	// str = fmt.Sprintf("%d", num1)
	// fmt.Printf("类型:%T	类型:%T	值:%v\n", num1, str, str)
	// fmt.Println(str)
	// str = fmt.Sprintf("%f", num2)
	// fmt.Printf("类型:%T	类型:%T  值:%q\n", num2, str, str)

	//使用方式二 使用strconv 函数
	str = strconv.FormatInt(int64(num1), 10)
	str = strconv.FormatFloat(num2, 'f', 10, 64)
	str = strconv.FormatBool(b)
	fmt.Printf("类型:%T	类型:%T  值:%q\n", num2, str, str)

	//使用方式三 使用strconv包中的itoa函数
	str = strconv.Itoa(num1)
	fmt.Printf("类型:%T	类型:%T  值:%q\n", num1, str, str)

}
