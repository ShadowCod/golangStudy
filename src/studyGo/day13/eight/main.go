package main

import (
	"fmt"
	"time"
)

// channel细节
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello")
	}
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("test,err", err)
		}
	}()
	var myM map[int]int
	myM[1] = 1
}

func main() {
	// 管道可以声明为只读或者只写	默认情况下是双向的
	// var chan2 chan<- int //只写
	// chan2 = make(chan int, 3)
	// chan2 <- 20
	// num := <-chan2//只能写不能读

	// var chan3 <-chan int
	// chan3 = make(chan int, 3)
	// <-chan3

	// 使用select可以解决从管道去数据的阻塞问题
	// intChan := make(chan int, 5)
	// for i := 0; i < 5; i++ {
	// 	intChan <- i
	// }

	// strChan := make(chan string, 5)
	// for i := 0; i < 5; i++ {
	// 	strChan <- "str" + fmt.Sprintf("%d", i)
	// }
	// // label:
	// for {
	// 	// 注意：如果intChan一直没有关闭，不会一直阻塞而deadlock
	// 	// 会自动到下一个case匹配
	// 	select {
	// 	case v := <-intChan:
	// 		fmt.Println(v)
	// 	case v := <-strChan:
	// 		fmt.Println(v)
	// 	default:
	// 		fmt.Println("没有了")
	// 		// break label
	// 		return
	// 	}
	// }

	go sayHello()
	go test()
	for i := 0; i < 11; i++ {
		time.Sleep(time.Second)
		fmt.Println("worl")
	}
}
