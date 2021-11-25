package main

import "fmt"

func main() {
	// 关闭channel	使用内置函数close
	intChan := make(chan int, 3)
	intChan <- 3
	intChan <- 5
	close(intChan) //关闭管道，不能在写入，但是可以读取，且返回两个值
	//intChan <- 6//panic: send on closed channel
	// num := <-intChan
	// fmt.Println(num)

	// channel的遍历使用for-range,如果未关闭则会出现死锁错误
	// 遍历channel不能使用for循环
	for v := range intChan { //管道只会返回一个值
		fmt.Println(v)
	}
}
