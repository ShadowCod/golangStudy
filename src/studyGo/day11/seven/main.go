package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// interface实践
// 声明一个结构体
type Hero struct {
	Name string
	Age  int
}

// 声明一个结构体切片类型
type HeroSlice []Hero

// 实现sort.Sort接口的方法
func (h HeroSlice) Len() int {
	return len(h)
}

// less这个方法就是决定使用什么标准排序
// 按age从大到小
func (h HeroSlice) Less(i, j int) bool {
	// return h[i].Age > h[j].Age
	// 修改成对Name排序
	return h[i].Name > h[j].Name
}

func (h HeroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

type Student struct {
	Name  string
	Age   int
	Score int
}

type StuSlice []Student

func (s StuSlice) Len() int {
	return len(s)
}

func (s StuSlice) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

func (s StuSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	// 测试
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	fmt.Println(heros)

	sort.Sort(heros)
	fmt.Println(heros)
}
