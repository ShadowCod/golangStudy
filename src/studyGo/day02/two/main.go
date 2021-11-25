package main

import (
	"fmt"
	"strings"
)

//字符串(双引号,单引号是字符)
func main() {
	s := "hello"
	//转义符号\  单引号不需要转义
	path := "\"D:\\Go\\bin\""
	fmt.Printf("%v\n", s)
	fmt.Println(path)

	//多行的字符串| ``(数字1旁边的)
	s2 := `
		世情薄
		人情恶
		雨送黄昏花易落
	`
	fmt.Println(s2)

	s3 := `D:\Go\bin`
	fmt.Println(s3)
	fmt.Printf("%T\n", s3)
	//字符串长度
	fmt.Println(len(s3))

	//拼接字符串
	s4 := "hehe"
	s5 := "haha"
	fmt.Println(s4 + s5)
	ss := fmt.Sprintf("%s%s", s4, s5)
	fmt.Println(ss)

	//分割字符串
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	fmt.Printf("%T\n", ret)

	//包含
	fmt.Println(strings.Contains(ss, "hahe"))

	//前缀
	fmt.Println(strings.HasPrefix(ss, "he"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "he"))

	//字符所在位置
	s6 := "abcdefga"
	fmt.Println(strings.Index(s6, "a"))
	fmt.Println(strings.LastIndex(s6, "a"))

	//拼接
	fmt.Println(strings.Join(ret, "_"))
}
