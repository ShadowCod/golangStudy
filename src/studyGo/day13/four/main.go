package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Cat struct {
	Name string
	Age  int
}

type Person struct {
	Name string
	Age  int
}

func main() {
	// allChan := make(chan interface{}, 3)
	// allChan <- "tom"
	// allChan <- 10
	// cat := Cat{"小白", 3}
	// allChan <- cat
	// // 如果只想要第3个，则需要将前两个推出
	// <-allChan
	// <-allChan
	// newCat := <-allChan
	// fmt.Printf("newCat的类型:%T,newCat的值:%v\n", newCat, newCat)
	// // 类型断言		因为取出时是空接口类型
	// a := newCat.(Cat)
	// fmt.Println(a.Name)

	// 定义channel类型
	var personChan chan Person
	personChan = make(chan Person, 10)
	// 放入随机时需要的随机种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		// 生成person
		num := rand.Intn(10)
		person := Person{
			Name: "person" + strconv.Itoa(num),
			Age:  rand.Intn(30),
		}
		personChan <- person
	}
	fmt.Println(len(personChan))
	// // 变量channel	不能使用for，需要使用for-range且只有一个值，前提需要关闭管道
	// for i := 0; i < 10; i++ {
	// 	newP := <-personChan
	// 	fmt.Println(newP)
	// }
}
