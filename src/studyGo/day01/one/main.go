package main

//导入语句
import (
	"fmt"
	"unsafe"
)

//函数外只能生成标识符(变量|常量)
//程序入口
func main() {
	// fmt.Println("study Golang!")
	//转义字符练习   常见的：\t	\\	\n	\r	\"
	fmt.Println("姓名\t年龄\n张三\t20")
	var a int64 = 10
	// 查看变量类型:%T	查看占用字节:unsafe.Sizeof()
	fmt.Printf("变量类型:%T,变量占用字节:%d", a, unsafe.Sizeof(a))
}
