package main

import "fmt"

// 声明一个CatNode结构体
type CatNode struct {
	no   int
	name string
	next *CatNode
}

// 向环形单链表中添加节点
func Add(head *CatNode, val *CatNode) {
	// 判断是不是添加第一只
	if head.next == nil {
		head.no = val.no
		head.name = val.name
		head.next = head //构成一个环形
		fmt.Println(val, "加入到环形的链表")
		return
	}
	// 先定义一个临时的变量，找到环形的最后节点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 让加进来的指向head，temp指向新加进来的
	val.next = head
	temp.next = val

}

// 删除环形链表中的某个节点
func Del(head *CatNode, val int) *CatNode {
	// 让temp指向head
	temp := head
	// 让helper指向环形链表的最后
	helper := head
	// 空链表
	if temp.next == nil {
		fmt.Println("环形链表为空")
		return head
	}
	// 如果只有一个节点
	if temp.next == head {
		temp.next = nil
		return head
	}
	// 将helper定位到环形列表的最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	// 两个及以上
	flag := true
	for {
		if temp.next == head {
			// 说明来到了最后一个，还没有比较
			break
		}
		if temp.no == val {
			if temp == head { //删除的是头节点
				head = head.next
			}
			// 找到了
			helper.next = temp.next
			fmt.Println("删除：", temp)
			flag = false
			break
		}
		temp = temp.next     //移动【比较】
		helper = helper.next //移动【一旦找到要删除的节点】
	}
	if flag { //在上面没有删除
		if temp.no == val {
			helper.next = temp.next
			fmt.Println("删除：", temp)
		} else {
			fmt.Println("未找到")
		}
	}
	return head
}

// 显示环形链表中的元素
func Show(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("环形链表为空")
		return
	}
	for {
		fmt.Printf("[%d]=%s\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	// 初始化一个环形链表的头节点
	head := &CatNode{}
	// 创建一只猫实例
	cat1 := &CatNode{
		no:   1,
		name: "tom1",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}
	Add(head, cat1)
	Add(head, cat2)
	Add(head, cat3)
	Show(head)
	head = Del(head, 61)
	Show(head)
}
