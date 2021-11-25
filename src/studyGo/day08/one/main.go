package main

import (
	"fmt"
)

//接收输入
func main() {
	var name string
	var age byte
	var sal float64
	var isPass bool

	fmt.Println("请输入名称")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪资")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过")
	fmt.Scanln(&isPass)
	fmt.Printf("名称%v 年龄%v 薪资%v 是否通过:%v", name, age, sal, isPass)

	//方法二，使用scanf
	fmt.Scanf("%v %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名称%v 年龄%v 薪资%v 是否通过:%v", name, age, sal, isPass)
}
