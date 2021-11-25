package main

import "fmt"

// 使用环形链表实现约瑟夫问题

// 小朋友结构体
type Child struct {
	No   int
	Next *Child
}

// 编写函数构成单向的环形链表
// 函数最后返回该环形链表的头节点
func Create(val int) *Child {
	head := &Child{}     //空节点 指针
	curChild := &Child{} //辅助节点 指针
	if val < 1 {
		fmt.Println("数量不正确")
		return head
	}
	// 循环构建环形链表
	for i := 1; i <= val; i++ {
		// 声明结构体
		child := &Child{
			No: i,
		}
		// 第一个需要特殊处理
		// 分析构成环形链表需要一个辅助指针
		if i == 1 {
			head = child //头节点不能动
			curChild = child
			curChild.Next = head
		} else {
			curChild.Next = child
			curChild = child
			curChild.Next = head //构成环形链表
		}
	}
	return head
}

// 显示环形链表中的元素
func Show(head *Child) {
	temp := head
	if temp.Next == nil {
		fmt.Println("空环形链表")
		return
	}
	for {
		fmt.Printf("编号:%d->", temp.No)
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
	fmt.Println()
}

// 实现功能的函数
func PlayGame(head *Child, startNo int, countNum int) {
	// 1.处理空链表
	if head.Next == nil {
		fmt.Println("空链表")
		return
	}
	// 判断startNo要小于总数
	// 2.需要定义一个辅助指针
	tail := head
	// 让tail指向环形链表的最后要给节点
	// 在删除时需要使用到
	for {
		if tail.Next == head { //说明到了最后一个节点
			break
		}
		tail = tail.Next
	}
	// 让head移动到startNo[删除是以head为准]
	for i := 1; i <= startNo-1; i++ {
		head = head.Next
		tail = tail.Next
	}
	// 移动countNum下并删除
	for {
		// 开始数countNum-1下
		for i := 1; i <= countNum-1; i++ {
			head = head.Next
			tail = tail.Next
		}
		fmt.Printf("编号%d出列\n", head.No)
		// 删除head指向的节点
		head = head.Next
		tail.Next = head
		// 判断head和tail是否指向同一个节点
		if head == tail {
			fmt.Printf("最后一个为%d", head.No)
			break
		}
	}
}

func main() {
	// 创建环形链表
	head := Create(50)
	// fmt.Println(head)
	Show(head)
	PlayGame(head, 2, 10)
}
