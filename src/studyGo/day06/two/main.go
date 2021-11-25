package main

import "fmt"

//defer语句:执行顺序先进后出（一个函数中可以有多个defer）
//使用情景：执行完需要释放资源时
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("hello") //defer把它后面的语句延迟到将返回时再执行
	defer fmt.Println("world") //一个函数中可以有多个defer
	defer fmt.Println("china") //执行顺序先进后出
	fmt.Println("end")
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

//defer经典案例
func f1() int { //赋值后return返回的那个值已经是5，再修改x变得只是x
	x := 5
	defer func() {
		x++
	}()
	return x //第一步：给返回值赋值  未命名返回值 = 5  第二步：执行defer，x=6  第三步：返回 未命名返回值
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //第一步：给返回值赋值  x = 5  第二步：执行defer，x=6  第三步：返回 x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //第一步：给返回值赋值  y = 5  第二步：执行defer，x=6  第三步：返回 y
}

func f4() (x int) {
	defer func(x int) { //x是值类型，传入的应该是副本
		x++
	}(x)
	return 5 //第一步：给返回值赋值  x = 5  第二步：执行defer，传入x的副本  第三步：返回x
}
