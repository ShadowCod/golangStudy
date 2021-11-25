package main

import "fmt"

// 断言
type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point = Point{1, 2}
	a = point
	var b Point
	// b = a 不可以		b=a.(Point)类型断言
	b = a.(Point)
	fmt.Println(b)

	// 类型断言的案例
	// var x interface{}
	// var c float32 = 11.1
	// x = c
	// y := x.(float32) //必须知道x原先的类型
	// fmt.Printf("类型%T,值%v", y, y)

	// 使断言带上检查机制
	var x interface{}
	var c float32 = 11.1
	x = c
	y, ok := x.(float64)
	if ok {
		fmt.Printf("类型%T,值%v", y, y)
	} else {
		fmt.Println("转换失败")
	}

	fmt.Println("继续执行...")
}
