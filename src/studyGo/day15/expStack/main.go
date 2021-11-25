package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
使用栈实现计算表达式
*/

// 定义一个栈的结构体
type Stack struct {
	MaxTop int
	Top    int
	Arr    [10]int
}

// 入栈
func (s *Stack) Push(val int) (err error) {
	// 判断栈是否已满
	if s.Top == s.MaxTop-1 {
		err = errors.New("stack full")
		return
	}
	// 入栈操作
	s.Top++
	s.Arr[s.Top] = val
	return
}

// 出栈
func (s *Stack) Pop() (val int, err error) {
	// 判断栈是否为空
	if s.Top == -1 {
		err = errors.New("stack empty")
		return
	}
	// 出栈操作
	val = s.Arr[s.Top]
	s.Top--
	return
}

// 展示栈中记录的值
func (s *Stack) Show() (err error) {
	//判断栈是否为空
	if s.Top == -1 {
		err = errors.New("stack empty")
		return
	}
	// 遍历栈中的元素并打印（需要从上往下遍历）
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d", i, s.Arr[i])
	}
	return
}

// 判断运算是否为合法运算符
func IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 运算的发放
func Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

// 优先级判断
func Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	// 声明数栈
	numStack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	// 声明符号栈
	operStack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	// 表达式
	// 注意：表达式需要做多位数的处理
	exp := "30+3*6-4"
	// 定义一个index，帮助扫描exp
	index := 0
	// 定义需要用到的变量
	num1, num2, oper, res := 0, 0, 0, 0
	keepNum := ""
	for {
		// 注意：表达式需要做多位数的处理
		ch := exp[index : index+1] //"3"
		// 将"+"换行成ascii
		temp := int([]byte(ch)[0])
		// 判断是数字还是符号
		if IsOper(temp) {
			// 符号部分有两个逻辑
			// 1.空栈情况
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				/*
					2.1如果栈顶的运算符的优先级大于等于当前要入栈的运算符的优先级，就从栈中pop出运算符，同时从数栈中pop两个值进行运算，运算结果重新入栈到数栈，再将要入栈的运算符入符号栈
				*/
				if Priority(operStack.Arr[operStack.Top]) >= Priority(temp) {
					oper, _ = operStack.Pop()
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					res = Cal(num1, num2, oper)
					numStack.Push(res)
					operStack.Push(temp)
				} else {
					/*
						2.2直接入符号栈
					*/
					operStack.Push(temp)
				}
			}

		} else { //说明是数字
			// numStack.Push(temp) 此处temp是ascii码值需要转换成数字
			// 处理多位数的思路
			// 1.定义一个变量 keepNum string,做拼接
			// 2.每次要想index后面的字符测试一下，看是不是运算符
			keepNum += ch
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}
		// 继续扫描
		if index+1 == len(exp) {
			break
		} else {
			index++
		}
	}
	// 将符号栈的所有元素取出进行计算
	for {
		if operStack.Top == -1 {
			break
		}
		oper, _ = operStack.Pop()
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		res = Cal(num1, num2, oper)
		numStack.Push(res)
	}
	result, _ := numStack.Pop()
	fmt.Println(result)
}
