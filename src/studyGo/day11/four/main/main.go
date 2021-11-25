package main

import (
	"fmt"
	"studyGo/day11/four/model"
)

// 封装案例
func main() {
	p := model.CreatePerson("tom")
	fmt.Println(p)
	p.SetAge(24)
	fmt.Println(p)
}
