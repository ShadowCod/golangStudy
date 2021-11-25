package main

import "fmt"

// 使用管道解决资源争夺的问题
func main() {
	// 创建一个可以存放int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Println(intChan)
	// 向管道写入数据
	// 注意当给管道写入数据时不能超过容量
	intChan <- 10
	num := 100
	intChan <- num
	fmt.Println(len(intChan))
	// 从管道中读取数据
	// 在没有使用协程的情况下，如果管道数据已经全部取出，再取会报deadlock错误
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Println(num1, num2, num3, len(intChan))
}
