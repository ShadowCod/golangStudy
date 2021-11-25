package main

import (
	"errors"
	"fmt"
)

/*
栈：先进后出，后进先出
*/
// 使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int //栈最大能存多少
	Top    int //栈顶，因为栈底是固定的
	Arr    [5]int
}

//入栈方法
func (s *Stack) Push(val int) (err error) {
	// 先判断栈是否已经满了
	if s.Top == s.MaxTop-1 {
		err = errors.New("stack full")
		return
	}
	// 加入
	s.Top++
	s.Arr[s.Top] = val
	return
}

// 展示栈里面的内容
func (s *Stack) Show() (err error) {
	// 判断是否为空栈
	if s.Top == -1 {
		err = errors.New("empty stack")
		return
	}
	// 遍历栈中的内容
	// 注意：需要从栈顶开始向下遍历
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("第%d个值为:%d\n", i, s.Arr[i])
	}
	return
}

// 出栈
func (s *Stack) Pop() (val int, err error) {
	// 判断栈是否为空
	if s.Top == -1 {
		err = errors.New("empty stack")
		return
	}
	// 取出栈顶的值
	val = s.Arr[s.Top]
	s.Top--
	return
}

func main() {
	// 定义一个栈
	stack := &Stack{
		MaxTop: 5,
		Top:    -1, //当栈为-1是，表示栈为空
	}
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)
	stack.Show()
	val, _ := stack.Pop()
	fmt.Println(val)
	stack.Show()
}
