package main

import "fmt"

func fbn(n int) []uint64 {
	var slice = make([]uint64, n)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}

func main() {
	// string底层是一个byte数组，因此可以进去切片处理
	// str := "hello@golang"
	// slice := str[6:]
	// fmt.Println(slice)
	// string是不可变的
	// str[0] = 'z' //抛错

	// 如果需要修改字符串，string->[]byte|[]rune->修改->string
	// slice2 := []rune(str)
	// slice2[0] = 'z'
	// str = string(slice2)
	// fmt.Println(str)
	// 细节：转成byte切片只能处理英文和数组，不能处理中文

	// 切片练习
	slice := fbn(10)
	fmt.Println(slice)
}
