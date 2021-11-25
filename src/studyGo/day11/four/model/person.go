package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

// 工厂模式函数 ----- 创建person结构体
func CreatePerson(n string) *person {
	return &person{
		Name: n,
	}
}

func (p *person) SetAge(n int) {
	if n > 0 && n < 150 {
		p.age = n
	} else {
		fmt.Println("输入年龄错误")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(n float64) {
	if n < 3000 || n > 15000 {
		fmt.Println("输入薪资错误")
	} else {
		p.sal = n
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
