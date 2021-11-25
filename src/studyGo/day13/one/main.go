package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// goroutine(协程)
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello,world", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {

	go test() //开启协程

	for i := 1; i <= 10; i++ {
		fmt.Println("hello,gonglang", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	// 逻辑cpu数量
	num := runtime.NumCPU()
	fmt.Println(num)
	// 设置go使用cpu的数量
	runtime.GOMAXPROCS(4)
}
