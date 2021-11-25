package main

import (
	"fmt"
	"os"
)

/*
hashTable|散列
*/
// 定义Emp结构体
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 定义EmpLink结构体
// 这里使用的EmpLink不带表头，即第一个节点就存放雇员
type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) Insert(emp *Emp) {
	// 判断该链表是否为空
	if this.Head == nil {
		this.Head = emp
		return
	}
	// 定义两个辅助指针
	cur := this.Head
	var pre *Emp = nil
	// 不是空链表的情况，通过循环cur中的ID和要插入的ID进行比较找到对应的位置后退出循环[注意:pre指针要一直保持在cur前一个]
	for {
		if cur == nil { //说明已经比较到最后一个元素了
			break
		} else {
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		}
	}
	// 将emp插入到pre和cur之间
	// 注意：加入现在列表中有一个元素，要插入一个ID小于该元素的值会出问题
	if pre == nil {
		this.Head = emp
		emp.Next = cur
	} else {
		pre.Next = emp
		emp.Next = cur
	}
}

// 显示链表中的内容
func (this *EmpLink) ShowLink(num int) {
	if this.Head == nil {
		fmt.Printf("第%d条链表没有数据\n", num)
		return
	}
	cur := this.Head
	for {
		fmt.Printf("链表%d,编号%d,名称%s", num, cur.Id, cur.Name)
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	fmt.Println()
}

// 定义hashTable,含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 给HashTable编写Insert方法  添加雇员
func (this *HashTable) Insert(emp *Emp) {
	// 使用散列函数，确定该雇员添加到那个链表
	linkNo := this.HashFunc(emp.Id)
	this.LinkArr[linkNo].Insert(emp)
}

// 编写一个散列方法
func (this *HashTable) HashFunc(id int) int {
	return id % 7 //得到链表的下标
}

// 编写方法显示hashTable中的所有元素
func (this *HashTable) Show() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

func main() {
	var chioce string
	var id int
	var name string
	hashTable := &HashTable{}
	for {
		fmt.Println("------menu------")
		fmt.Println("---1.input添加")
		fmt.Println("---2.show 展示")
		fmt.Println("---3.exit 退出")
		fmt.Println("请输入:")
		fmt.Scanln(&chioce)
		switch chioce {
		case "input":
			fmt.Println("请输入ID:")
			fmt.Scanln(&id)
			fmt.Println("请输入名称:")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.Show()
		case "exit":
			os.Exit(0)
		}
	}
}
