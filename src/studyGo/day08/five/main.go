package main

//常用字符串函数
func main() {
	//1) 计算长度 内置函数，根据传入的值返回
	// str := "hello"        //ascii中的占一个字节，中文占三个字节
	// fmt.Println(len(str)) //按字节返回

	// 2) 字符串遍历，同时处理中文的问题
	// str2 := "hello 北京"
	// //使用传统的遍历方式打印
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("字节：%c\n", str2[i]) //按字节输出，会出现乱码
	// }
	// // 先将字符串转换成[]rune切片
	// rune1 := []rune(str2)
	// for i := 0; i < len(rune1); i++ {
	// 	fmt.Printf("字符：%c\n", rune1[i])
	// }
	// // for range 方式
	// for _, v := range str2 {
	// 	fmt.Printf("字符：%c\n", v)
	// }

	// 3) 字符串转整数
	// n, err := strconv.Atoi("abc")
	// if err != nil {
	// 	fmt.Println("转换错误")
	// } else {
	// 	fmt.Println(n)
	// }

	// 4) 整数转字符串
	// str := strconv.Itoa(123)
	// fmt.Printf("类型:%T,值:%v\n", str, str)

	// 5) 字符串转换成baty切片
	// var bytes = []byte("hello go")
	// fmt.Println(bytes)

	// 6) []byte转换成字符串
	// var str = string([]byte{98, 99})
	// fmt.Println(str)

	// 7) 10进制转2,8,16进制，返回的是字符串
	// str := strconv.FormatInt(123, 2)
	// fmt.Println(str)
	// str = strconv.FormatInt(123, 16)
	// fmt.Println(str)

	// 8) 查找字符串中有无子串
	// str := "hello"
	// b := strings.Contains(str, "ho")
	// fmt.Println(b)

	// 9) 统计一个字符串有几个指定的子串
	// n := strings.Count("hello", "ll")
	// fmt.Println(n)

	// 10) 不区分大小写的字符串比较（==是区分大小写的）
	// fmt.Println(strings.EqualFold("abc", "ABC"))
	// fmt.Println("abc" == "aBc")

	// 11) 返回子串第一次出现在字符串中的index,没有则返回-1 (也可以用于判断字符串中是否有子串)
	// n := strings.Index("abcrgt", "cg")
	// fmt.Println(n)

	// 12) 返回子串最后出现中字符串中的index
	// n := strings.LastIndex("sdjfoiaosabcajdfabc", "abc")
	// fmt.Println(n)

	// 13) 将指定的子串替换strings.Replace("go go hello","go","go语言",n)，全部替换n=-1
	// str1 := "go go hello"
	// str := strings.Replace(str1, "go", "go语言", 1) //返回的是新的字符串，本身不变
	// fmt.Println(str, str1)

	// 14) 按照指定的某个字符进去分割,将字符串拆分成字符串切片
	// str := "hello,word,golang"
	// strArr := strings.Split(str, ",")
	// fmt.Printf("类型:%T", strArr)
	// fmt.Println(strArr)

	// 15) 将字符串的字母进行大小写的转换
	// str := strings.ToUpper("golang")
	// fmt.Println(str)
	// str1 := strings.ToLower("GoLang")
	// fmt.Println(str1)

	// 16) 将字符串左右两边的空格去掉(只会去除两边的)
	// str := strings.TrimSpace(" the a long ")
	// fmt.Println(str)

	// 17) 将字符串左右两边指定的字符去掉(只能去掉左右两边的) 返回新串
	// str := strings.Trim("! hello !", "! ")
	// fmt.Println(str)

	// 18) 只去掉字符串左边（右边）的字符
	// str := strings.TrimLeft("! hello !", "!")
	// str1 := strings.TrimRight("! hello !", "!")
	// fmt.Println(str)
	// fmt.Println(str1)

	// 19) 判断字符串是否以某个字符串开头或者结尾
	// b := strings.HasPrefix("ftp://192.168.0.1", "ftp")
	// b1 := strings.HasSuffix("ftp://192.168.0.1", "0.1")
	// fmt.Println(b)
	// fmt.Println(b1)

}
