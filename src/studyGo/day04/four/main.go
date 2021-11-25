package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("值:%v 长度:%d 容量:%d\n", s1, len(s1), cap(s1))
	// s1[3] = "广州" // 索引越界
	// fmt.Println(s1)

	//为切片追加元素: 调用append必须用原来的变量接收返回值
	s1 = append(s1, "广州") //append追加元素，底层数组放不下，GO就会把底层数组换一个
	fmt.Printf("值:%v 长度:%d 容量:%d\n", s1, len(s1), cap(s1))
	s2 := []string{"杭州", "成都", "苏州"}
	s1 = append(s1, s2...) //...是拆开的意思
	fmt.Printf("值:%v 长度:%d 容量:%d\n", s1, len(s1), cap(s1))

	//拷贝切片
	a1 := []int{1, 2, 3}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 1000
	fmt.Println(a1, a2, a3)

	//将a1中的2删除
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)

	//1.切片不保存具体的值
	//2.切片对应一个底层的数组
	//3.底层数组都是占用一块连续的内存
	x1 := [...]int{1, 3, 5}
	x2 := x1[:]
	fmt.Println(x2, len(x2), cap(x2))
	x2 = append(x2[:1], x2[2:]...) //修改了底层数组
	fmt.Println(x2, len(x2), cap(x2))
	x2[0] = 100 //修改底层数组
	fmt.Println(x1)

	//练习
	var a = make([]int, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
}
