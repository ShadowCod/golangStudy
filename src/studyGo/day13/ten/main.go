package main

import (
	"fmt"
	"reflect"
)

//修改int类型的值
func reflect01(i interface{}) {
	// 获取reflect.Value类型
	rVal := reflect.ValueOf(i)
	fmt.Printf("rVal的kind:%T\n", rVal) //注意，因为int是值类型，改变其值需要传入指针
	// 使用SetInt()方法修改值
	// rVal.SetInt(20)//这里会出错，因为是指针，需要先使用rVal.Elem()方法取得指向的值才能进行修改
	rVal.Elem().SetInt(20)
}

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%v完成了减法运行,%v-%v=%v", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func reflect03(i interface{}) {
	rVal := reflect.ValueOf(i)
	var argSlice []reflect.Value
	argSlice = append(argSlice, reflect.ValueOf("tom"))
	rVal.Method(0).Call(argSlice)
}

func reflect02(i interface{}) {
	rVal := reflect.ValueOf(i)
	elem := rVal.Elem()
	elem.FieldByName("Name").SetString("tom")
}

type Monster struct {
	Name string
	Age  int
}

// 测试协程顺序
func inputFunc(c chan int) {
	fmt.Println("inputFunc")
	for i := 1; i < 100000; i++ {
		c <- i
	}
	close(c)
}

func getFunc(c chan int, r chan bool) {
	fmt.Println("getFunc")
	for {
		num, ok := <-c
		if !ok {
			break
		}
		fmt.Println(num)
	}
	r <- true
}

func main() {
	// num := 10
	// reflect01(&num)
	// fmt.Println(num)
	// monster := Monster{
	// 	Name: "jack",
	// 	Age:  20,
	// }
	// reflect02(&monster)
	// fmt.Println(monster.Name)
	// cal := Cal{20, 10}
	// reflect03(cal)

	var intChan chan int
	intChan = make(chan int, 2000)
	var resChan chan bool
	resChan = make(chan bool, 4)
	go inputFunc(intChan)
	go getFunc(intChan, resChan)
	// for i := 0; i < 4; i++ {

	// }
	go func() {
		for i := 0; i < 4; i++ {
			<-resChan
		}
		close(resChan)
	}()
}
