package main

import "fmt"

func inputData(c chan int) {
	for i := 1; i <= 8000; i++ {
		c <- i
	}
	close(c)
}

func primeFunc(c, r chan int, e chan bool) {
	var flag bool
	for {
		num, ok := <-c
		flag = true
		if !ok {
			e <- true
			break
		}
		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //不是素数
				flag = false
				break
			}
		}
		if flag {
			r <- num
		}
	}
}

func main() {
	var intChan chan int
	var resChan chan int
	var exitChan chan bool

	intChan = make(chan int, 1000)
	resChan = make(chan int, 2000)
	exitChan = make(chan bool, 4)

	go inputData(intChan)

	// 启动4个判断线程
	for i := 0; i < 4; i++ {
		go primeFunc(intChan, resChan, exitChan)
	}

	// 判断是否所有线程都运行完毕
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(resChan)
		close(exitChan)
	}()
	// 遍历素数
	// for v := range resChan {
	// 	fmt.Println(v)
	// }
	for {
		num, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println(num)
	}
}
