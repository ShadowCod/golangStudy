package main

import "fmt"

//make函数制造切片
func main() {
	s1 := make([]int, 5, 10) //第一个参数：存储类型   第二个参数：长度   第三个参数：容量
	fmt.Printf("s1:%v s1_len:%d s1_cap:%d\n", s1, len(s1), cap(s1))
	//切片的本质就是一个框，框一个连续的内存	属于引用类型，真正的数据存在底层数组中
	s2 := make([]int, 0, 10)
	fmt.Printf("s2:%v s2_len:%d s2_cap:%d\n", s2, len(s2), cap(s2))

	//切片的赋值,切片不能保存值
	s3 := []int{1, 3, 5}
	s4 := s3 //s3和s4都指向同一个底层数组
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	//切片的遍历
	//1.索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	//2.for range循环
	for _, v := range s3 {
		fmt.Println(v)
	}
}
