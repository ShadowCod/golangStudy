package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test01(arr *[3]int) {
	(*arr)[0] = 11
	fmt.Println(*arr)
}

// 数组：存放多个同一类型数据，数组也是一种数据类型，Go中是值类型
func main() {
	// num1 := 2.0
	// num2 := 4.0
	// num3 := 6.0
	// num4 := 8.0
	// total := num1 + num2 + num3 + num4
	// // 将total/4的结果四舍五入保留到小数点两位
	// avg := fmt.Sprintf("%.2f", total/4)
	// fmt.Println(avg)

	// 数组的地址是数组的首地址（第一个元素的地址）,各个地址的间隔取决于存储的类型（int64->8字节）
	// var intArr [3]int
	// 定义完数组后存在默认类型的零值
	// fmt.Println(intArr)
	// fmt.Printf("地址:%p\n", &intArr)

	// var score [5]float64
	// for i := 1; i < len(score); i++ {
	// 	fmt.Printf("请输入第%d个值", i+1)
	// 	fmt.Scanln(&score[i])
	// }
	// fmt.Println(score)

	//使用函数改变数组本身的值
	// var arr = [3]int{1, 2, 3}
	// fmt.Println(arr)
	// test01(&arr)
	// fmt.Println(arr)

	// 数组应用案例
	// var byteArr [26]byte
	// for i := 0; i < len(byteArr); i++ {
	// 	byteArr[i] = 'A' + byte(i) //计算要统一类型
	// }
	// fmt.Println(byteArr)
	// for _, v := range byteArr {
	// 	fmt.Printf("%c", v)
	// }

	var intArr2 [5]int
	len := len(intArr2)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr2[i] = rand.Intn(100) //随机时如果seed是固定的，结果也是固定的
	}
	fmt.Println("反转前:", intArr2)

	// 第一个和最后一个，第二个和倒数第二个(转换的次数是len/2,如果转换len次则又被转回来了)
	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr2[len-1-i]
		intArr2[len-1-i] = intArr2[i]
		intArr2[i] = temp
	}
	fmt.Println("反转后:", intArr2)
}

// 上面代码优化：多次使用len()内置函数，可以将其值赋值给一个变量进行处理
