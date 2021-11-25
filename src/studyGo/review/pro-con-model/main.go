package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
	生产者每天生产一件产品
	每生产一件消费者就消费一件
	消费10轮后，主协程退出
*/

// 思路：生产者每生产一件物品就往管道中添加，管道中存在物品消费者就消费

// 产品结构体
type Product struct {
	Name string
}

func main() {
	// 生成【商店管道】【计数管道】
	shopChan := make(chan *Product, 5)
	countChan := make(chan int, 10)
	// 生产者源源不断的向【商店管道】写入产品
	go produce(shopChan)
	// 消费者阻塞等待从【商店管道】读出数据
	// 消费者每读出一次数据（消费一次）就向【计数管道】写入一个数据
	go consume(shopChan, countChan)
	// 主协程阻塞从【计数管道】读出10个数据就结束
	for i := 0; i < 10; i++ {
		<-countChan
	}
	// 管道使用完毕后需要关闭（channel没有被任何协程用到后最终会被GC回收）
	close(shopChan)
	close(countChan)
	fmt.Println("main over")
}

// 生产者协程
func produce(c chan<- *Product) {
	//一直生产商品
	for {
		pro := &Product{
			Name: "商品" + strconv.Itoa(int(time.Now().Unix())),
		}
		c <- pro
		time.Sleep(1 * time.Second)
	}
}

// 消费者协程
func consume(s <-chan *Product, c chan<- int) {
	//一直等待，直到【商品管道】中有商品
	for {
		product := <-s
		fmt.Println("消费", product)
		// 往【计数管道】中插入次数
		c <- 1
	}
}
