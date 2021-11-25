package main

import "fmt"

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

//递归练习
func main() {
	// 给出一个整数，求斐波那契数
	fmt.Println(fbn(1))
	fmt.Println(fbn(3))
	fmt.Println(fbn(5))
}
