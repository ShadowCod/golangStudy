package main

// 使用全局锁解决资源争夺的问题
import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int) //因为需要多个协程操作，所以需要定义成全局的
	lock  sync.Mutex          //为了解决争夺资源的问题，可以使用一个全局互斥锁
)

func test(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}
	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}

func main() {

	for i := 1; i < 20; i++ {
		go test(i) //启动多个协程同时进行
	}
	// 问题：如果不等待，主程序执行完毕后就关闭了进程
	time.Sleep(time.Second * 5)
	// 问题：同时操作map会出现资源争夺出现错误
	lock.Lock()
	for i, v := range myMap {
		fmt.Println(i, v)
	}
	lock.Unlock()
}
