package main

import "fmt"

/*
	二叉树的三种遍历
*/
// 定义结构体
type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

// 前序遍历：先输出根节点，再输出左子树，再输出右子树
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("No:%d,Name:%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历：先输出左子树，再输出根节点，再输出右子树
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("No:%d,Name:%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

// 后序遍历：先输出左子树，再输出右子树，再输出根节点
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("No:%d,Name:%s\n", node.No, node.Name)
	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "A",
	}
	left1 := &Hero{
		No:   2,
		Name: "B",
	}
	right1 := &Hero{
		No:   3,
		Name: "E",
	}
	root.Left = left1
	root.Right = right1
	right2 := &Hero{
		No:   4,
		Name: "F",
	}
	right1.Right = right2

	left2 := &Hero{
		No:   5,
		Name: "C",
	}

	left3 := &Hero{
		No:   6,
		Name: "D",
	}
	right3 := &Hero{
		No:   7,
		Name: "G",
	}
	right4 := &Hero{
		No:   8,
		Name: "H",
	}
	right2.Left = right3
	right2.Right = right4
	left1.Left = left2
	left1.Right = left3

	// PreOrder(root)
	// fmt.Println()
	// InfixOrder(root)
	// fmt.Println()
	PostOrder(root)
}
