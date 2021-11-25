package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1                               //①
	b := 2                               //②
	defer calc("1", a, calc("10", a, b)) //defer(⑧,③)
	a = 0                                //④
	defer calc("2", a, calc("20", a, b)) //(⑦,⑤)
	b = 1                                //⑥
}

//流程
/*
	1.a = 1
	2.b = 2
	3.calc("10", 1, 2)	defer calc("1",1,3) 此处a已经进入栈中且值为1
	4.a = 0
	5.calc("20", 0, 2)	defer calc("2",0,2) 此处的a被赋值为0
	6.b = 1
	7.最后一个defer
	8.前面一个defer
*/
