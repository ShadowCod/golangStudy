package main

import "fmt"

type A struct {
	Num int
}

// 方法：func (recevier type) methodName (参数列表) (返回值列表) { 方法体 return 返回值}
func (a A) test() {
	fmt.Println("test", a.Num)
}

type Person struct {
	Name string
}

func (person Person) speak() {
	fmt.Println(person.Name + "是一个好人")
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// 为了提高效率，选择传递指针
func (c *Circle) area2() float64 {
	// return 3.14 * (*c).Radius * (*c).Radius
	return 3.14 * c.Radius * c.Radius
}

type Stu struct {
	Name string
	Age  int
}

func (s *Stu) String() string {
	res := fmt.Sprintf("name:%v age:%v\n", (*s).Name, (*s).Age)
	return res
}

func main() {
	// var a A
	// a.Num = 10
	// a.test()
	// person := Person{"tom"}
	// person.speak()
	var c = new(Circle)
	(*c).Radius = 5
	// res := (*c).area()
	res := c.area2()
	fmt.Println(res)

	stu := Stu{"tom", 20}
	fmt.Println(&stu)
}
