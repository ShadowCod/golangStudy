package main

import "fmt"

func main() {
	//运算符
	a := 10
	b := 3
	//算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能放在语句的右边赋值
	b-- //单独的语句，不能放在语句的右边赋值
	fmt.Println(a)
	fmt.Println(b)

	//关系运算符 GO语言是强类型的，只用相同类型才能比较
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a == b)
	fmt.Println(a != b)

	//逻辑运算符 （&&：AND  ||：OR  !:非）

	//位运算符:针对的是二进制数
	//5的二进制表示：101
	//2的二进制表示： 10
	//将二进制向左移动指定位数
	fmt.Println(5 << 1) //x乘以2的1次方  101=>1010
	//将二进制向右移动指定位数
	fmt.Println(5 >> 3) //x除以2的3次方	 101=>0
	//&：按位与(两位均为1才为1)
	fmt.Println(5 & 2) //000
	//|：按位或(两位有一个1就为1)
	fmt.Println(5 | 2) //111
	//^：异或（两位不一样才为1）
	fmt.Println(5 ^ 2) //111
}
