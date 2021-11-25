package main

import "fmt"

//数组:数组的长度是数组类型的一部分（必须指定存放元素的长度|容量和类型）
func main() {

	//定义：var 数组变量名 [元素数量]存储类型
	var a [3]int
	fmt.Println(a)

	//初始化：如果不初始化，默认元素都是零值（bool:false,整型和浮点型:0,字符串:""）
	//方法一（自定义个数）
	var firstArray [3]int //只是声明，默认初始化
	fmt.Println(firstArray)
	var twoArray = [3]int{1, 2} //注意：声明+赋值
	fmt.Println(twoArray)

	//方法二（编译器推断个数）
	var threeArray = [...]int{1, 2} //必须进行赋值
	fmt.Println(threeArray)

	//方法三（指定索引值）
	var fourArray = [...]int{1: 1, 3: 2}
	fmt.Println(fourArray)

	//遍历数组
	for _, v := range fourArray {
		fmt.Println(v)
	}

	//二维数组定义
	b := [3][2]int{{1, 2}, {3, 4}}
	fmt.Println(b)

	//编译器推断（只能有一层推断，一般是外层）
	c := [...][2]string{{"1", "2"}, {"3", "4"}}
	fmt.Println(c)

	//二维数组遍历
	for _, v1 := range c {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
		fmt.Println("------")
	}

	//数组是值类型，赋值和传参是复制整个数组
	modifyArray(twoArray)
	fmt.Printf("twoArray:%d\n", twoArray)

	practiceArray := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range practiceArray {
		sum += v
	}
	fmt.Printf("sum:%d\n", sum)

	practiceArray2 := [...]int{1, 3, 5, 7, 8}

	// for k, v := range practiceArray2 {
	// 	for k1, v1 := range practiceArray2 {
	// 		if v+v1 == 8 && k != k1 {
	// 			fmt.Printf("value1:%d value2:%d\n", k, k1)
	// 		}
	// 	}
	// }

	for i := 0; i < len(practiceArray2); i++ {
		for j := i + 1; j < len(practiceArray2); j++ {
			if practiceArray2[i]+practiceArray2[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func modifyArray(x [3]int) {
	x[0] = 100
}
