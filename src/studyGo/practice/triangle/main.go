package main

import "fmt"

/*
	打印正三角形
*/
// 分析得出:打印的字和行之间的关系是:2乘以行数-1，打印的空格和列之间的关系是:最大列数-打印的字的个数除以2
// Stu 结构体
type Stu struct {
	ID   int64
	Name string
}

func main() {
	// var row = 5 //行
	// var col = 9 //列
	// for i := 1; i <= row; i++ {
	// 	word := 2*i - 1
	// 	space := (col - word) / 2.0
	// 	for j := 0; j < space; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 0; k < word; k++ {
	// 		fmt.Print("1")
	// 	}
	// 	for l := 0; l < space; l++ {
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Println()
	// }
	/*
		注意：随机数字是使用的math/rand包，如果不给rand设定Seed那么随机出来的值永远一样,一般使用系统的毫秒作为种子
	*/
	// randNum := rand.Int()
	// fmt.Println(randNum)
	// rand.Seed(time.Now().Unix())
	// randNum2 := rand.Int()
	// fmt.Println(randNum2)
	// fmt.Println(time.Now().Unix())

	// 验证for range时v是否为复制的值
	stuSlice := []*Stu{
		&Stu{
			ID:   1,
			Name: "lili",
		},
		&Stu{
			ID:   2,
			Name: "jeck",
		},
	}
	fmt.Println("init_stuSlice", stuSlice)
	for k, v := range stuSlice {
		fmt.Println("K:", k)
		fmt.Println("v:", v)
		v.Name = "tom"
	}
	fmt.Println("stuSlice", stuSlice)
	stu1 := Stu{ID: 1, Name: "lishi"}
	stu2 := Stu{ID: 2, Name: "zhangshan"}
	fmt.Println(stu1.ID > stu2.ID)
}
