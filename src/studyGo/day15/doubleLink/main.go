package main

import "fmt"

// 定义一个双向链表的结构体
type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //指向前一个节点
	next     *HeroNode //指向后一个节点
}

// 向双向链表尾部插入一个节点
func InsertHeroNode(head *HeroNode, val *HeroNode) {
	// 1.创建一个辅助节点
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = val
	val.pre = temp
}

// 向双向链表中插入一个节点
func InsertHeroNode2(head *HeroNode, val *HeroNode) {
	temp := head
	// 使用一个标识符来判断是否加入
	flag := true
	// 让插入的节点no和temp的节点的no比较
	for {
		if temp.next == nil { //temp是最后一个节点
			break
		} else if temp.next.no > val.no {
			// 说明下一个节点的编号大于要插入的
			break
		} else if temp.next.no == val.no {
			// 说明链表中已有这个no
			flag = false
			break
		}
		// 注意这里一定要指向下一个
		temp = temp.next
	}
	if !flag {
		fmt.Println("编号已存在")
		return
	} else {
		// 注意：这里要判断temp是否为第一个|最后一个节点
		val.next = temp.next
		val.pre = temp
		if temp.next != nil {
			temp.next.pre = val
		}
		temp.next = val
	}
}

// 删除双向链表中的元素
func Delete(head *HeroNode, val int) {
	// 声明一个辅助节点
	temp := head
	// 判断双向链表是否为空
	if temp.next == nil {
		fmt.Println("双向链表为空")
		return
	}
	// 声明一个标志
	flag := false
	// 使temp指向要删除的节点
	for {
		if temp.no == val {
			flag = true
			break
		} else if temp.next == nil {
			break
		}
		temp = temp.next
	}
	if flag {
		if temp.next != nil {
			temp.next.pre = temp.pre
		}
		if temp.pre != nil {
			temp.pre.next = temp.next
		}
	}
}

// 显示双向链表的内容
func Show(head *HeroNode) {
	// 1.声明一个辅助节点
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		if temp.next == nil {
			break
		}
		fmt.Printf("[%d,%s,%s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
	}
	fmt.Println()
}

// 逆序显示
func InverseShow(head *HeroNode) {
	// 1.创建一个辅助节点
	temp := head
	// 2.让temp指向最后一个节点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	// 遍历链表
	for {

		if temp.pre == nil {
			break
		}
		fmt.Printf("[%d,%s,%s]<===>", temp.no, temp.name, temp.nickname)
		temp = temp.pre
	}
	fmt.Println()
}

func main() {
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:       1,
		name:     "唐三藏",
		nickname: "唐僧",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "孙悟空",
		nickname: "齐天大圣",
	}

	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	Show(head)
	// InverseShow(head)
	Delete(head, 1)
	Show(head)
}
