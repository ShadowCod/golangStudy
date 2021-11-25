package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [4]int //数组=>模拟队列
	front   int    //表示指向队列的队首
	rear    int    //表示指向队列的队尾
}

// 添加数据到队列的方法
func (this *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if this.rear == this.maxSize-1 { //重要提示：rear是队列尾部（含最后的元素）
		return errors.New("queue full")
	}
	this.rear++                 //rear后移
	this.array[this.rear] = val //数组中对应索引赋值
	return
}

// 从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if this.rear == this.front { //因为每取一个值的时候都是front+=1，因此当front==rear时就说明rear已经取了
		err = errors.New("queue empty")
		return
	}
	this.front++
	val = this.array[this.front]
	return
}

// 显示队列,找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	// 注意：此处的队首不包含值
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d] = %d\t", i, this.array[i])
	}
	fmt.Println()
}

// 使用主函数测试队列功能
func main() {
	// 先创建一个队列结构体
	queue := &Queue{
		maxSize: 4,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err)
			}
		case "get":
			// 使用取出队列值的方法
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("取出的值:%d\n", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
