package main

import (
	"fmt"
	"strings"
)

//闭包
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// jpgFunc := makeSuffixFunc(".jpg")
	// fmt.Println(jpgFunc("test"))
	// txtFunc := makeSuffixFunc(".txt")
	// fmt.Println(txtFunc("test"))
	f1, f2 := calc(10)

	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4)) //注意：前面调用了以后base不再是10，而是9
}
