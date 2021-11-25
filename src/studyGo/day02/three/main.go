package main

import "fmt"

//字符串修改(本身不能修改,需要转换成其他类型进行修改)
func main() {
	s1 := "白萝卜"
	fmt.Println(s1)
	s2 := []rune(s1) //把字符串强制转换成rune切片 [白 萝 卜]
	fmt.Println(s2)
	s2[0] = '红'             //应该是字符类型
	fmt.Println(string(s2)) //把rune切片强制转换成字符串
	c1 := "红"               //string
	c2 := '红'               //rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	//类型转换
	n := 10
	var f float64
	f = float64(n)
	fmt.Printf("%T\n", f)

	s8 := "hello中国"
	fmt.Println(len(s8))
}
