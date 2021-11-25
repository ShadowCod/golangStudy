package main

import "fmt"

//流程控制
func main() {
	// var i1 = 10
	//if语句
	// if i1 < 10 {
	// 	fmt.Println("小于")
	// } else {
	// 	fmt.Println("大于")
	// }

	//if多个判断
	// if i1 < 10 {
	// 	fmt.Println("小于")
	// } else if i1 == 10 {
	// 	fmt.Println("等于")
	// } else {
	// 	fmt.Println("大于")
	// }

	//在if判断中声明变量,改变量的作用域只在if判断中
	if age := 10; age > 10 {
		fmt.Println("age > 10")
	} else {
		fmt.Println("age <= 10")
	}
	// fmt.Println(age) //age抛错,超出作用域

	//GO语言中只有for一种循环语句(但是有很多的变种),条件处加()要抛语法错误
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}

	//变种一,省略初始语句(必须要有分号)
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	//变种二,省略结束语句
	var j = 5
	for j < 10 {
		fmt.Printf("j:%v\n", j)
		j++
	}

	//无限循环(死循环)
	// for {
	// 	fmt.Println(1)
	// }

	//键值循环(遍历)
	for k, v := range "hello" {
		fmt.Printf("k:%v,v:%c\n", k, v)
	}

	//使用for循环打印
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d =%d\t", i, j, i*j)
		}
		fmt.Println()
	}
	for i := 9; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d =%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
