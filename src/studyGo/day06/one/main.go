package main

import (
	"fmt"
)

func main() {
	//判断字符串中汉字的数量
	//判断一个字符是汉字
	// s1 := "Hello中国"
	// fmt.Println(s1)
	// count := 0
	// for _, v := range s1 {
	// 	if unicode.Is(unicode.Han, v) {
	// 		count++
	// 	}
	// }
	// fmt.Println("count:", count)

	//判断一段话中每个单词出现的次数
	// s2 := "how do you do"
	// s3 := strings.Split(s2, " ")
	// fmt.Println(s3)
	// m1 := make(map[string]int, 5)
	// for _, w := range s3 {
	// 	i, ok := m1[w]
	// 	if !ok {
	// 		m1[w] = 1
	// 	} else {
	// 		m1[w] = i + 1
	// 	}
	// }
	// fmt.Println(m1)

	//回文判断（字符串从左往右和从右往左一样）
	ss := "上海自来水来自海上"
	sr := make([]rune, 0, len(ss))
	for _, v := range ss {
		sr = append(sr, v)
	}
	fmt.Println("rune_slice:", sr)
	for i := 0; i < len(sr)/2; i++ {
		if sr[i] != sr[len(sr)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
}
