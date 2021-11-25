package main

import "fmt"

// 定义HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode //表示指向下一个节点
}

//给链表插入一个节点
// 第一种插入方式：在单链表的最后加入
func InsetHeroNode(head *HeroNode, val *HeroNode) {
	// 思路：
	// 1.先找到该链表最后的节点
	// 2.创建一个辅助节点
	temp := head
	for {
		if temp.next == nil { //表示最后的节点
			break
		}
		temp = temp.next //让temp不断的指向下一个节点
	}
	// for循环退出一定是指向了最后一个节点
	// 3.将val加入到最后
	temp.next = val

}

// 显示
func show(head *HeroNode) {
	// 1.创建一个辅助节点
	temp := head
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[%d,%s,%s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

// 根据no的编号从小到大插入
func InsetHeroNode2(head *HeroNode, val *HeroNode) {
	// 思路：
	// 1.找到适当的节点
	// 2.创建一个辅助节点
	temp := head
	// 使用一个标识符来判断是否加入
	flag := true
	// 让插入的节点no和temp的下一个节点的no比较
	for {
		if temp.next == nil {
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
		// 顺序不能颠倒，颠倒了后面的数据就丢失了
		val.next = temp.next
		temp.next = val
	}

}

// 删除一个节点
func delete(head *HeroNode, val int) {
	temp := head
	flag := false
	for {
		if temp.next.no == val {
			flag = true
			break
		} else if temp.next == nil {
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("不存在要删除的元素")
	}
}

func main() {
	// 创建一个头节点
	head := &HeroNode{}
	// 创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	InsetHeroNode2(head, hero2)
	InsetHeroNode2(head, hero1)

	show(head)
}
