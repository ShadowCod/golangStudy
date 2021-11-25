package main

import (
	"errors"
	"fmt"
	"os"
)

// 环形队列的容量是最大值减一，一个空的作为标志
// 使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int    //队列容量
	array   [5]int //数组==>环形队列
	head    int    //指向队首
	tail    int    //指向队尾
}

// 添加方法
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		err = errors.New("queue full")
		return
	}
	// 这样写tail是队尾，但是不包含值
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// 出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		err = errors.New("queue empty")
		return
	}
	// head指向队首，并且包含队首元素
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

// 判断队列已满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// 判断队列为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 查看环形队列中有多少个元素
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

// 显示队列
func (this *CircleQueue) Show() {
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	// 设计一个辅助变量，指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d] = %d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}
func main() {
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add向队列中添加值")
		fmt.Println("2.输入get向队列中取出值")
		fmt.Println("3.输入show展示队列中的值")
		fmt.Println("4.输入exit退出程序")
		fmt.Println("请输入:")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入要添加的值:")
			fmt.Scanln(&val)
			// 使用向队列添加值的方法
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err)
			}
		case "get":
			// 使用取出队列值的方法
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("取出的值:%d\n", val)
			}
		case "show":
			queue.Show()
		case "exit":
			os.Exit(0)
		}
	}
}
