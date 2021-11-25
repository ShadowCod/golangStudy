package main

import (
	"fmt"
	"reflect"
)

/* 反射入门*/

// 基本数据类型反射
func reflectTest(i interface{}) {
	// 获取reflect.Type类型
	rType := reflect.TypeOf(i)
	fmt.Printf("rType类型:%T\n", rType)
	vName := rType.Name()
	fmt.Println(vName)
	vKind := rType.Kind()
	fmt.Println(vKind)
	// 获取reflect.Value类型
	rValue := reflect.ValueOf(i)
	fmt.Printf("rValue类型:%T\n", rValue)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic err", err)
		}
	}()
	vAddr := rValue.Addr()
	fmt.Println(vAddr)
	// 将rValue类型转换为interface{}类型
	iType := rValue.Interface()
	// 使用断言将interface{}类型转换为想要的类型
	num := iType.(int64)
	fmt.Printf("num类型:%T\n", num)
}

// 结构体数据类型反射
func reflectTest02(i interface{}) {
	// 获取reflect.Type类型
	rType := reflect.TypeOf(i)
	fmt.Println(rType)
	// 获取reflect.Value类型
	rValue := reflect.ValueOf(i)
	fmt.Println(rValue)
	// 转换为空接口类型
	vType := rValue.Interface()
	// 空接口断言成对应的类型，可以用switch做得很灵活
	vValue, ok := vType.(Stu)
	if ok {
		fmt.Println("name:", vValue.Name)
	}
}

// 定义一个结构体
type Stu struct {
	Name string
	Age  int
}

func main() {
	// 基本数据类型使用反射
	// var num int64
	// num = 100
	// reflectTest(num)
	student := Stu{
		Name: "tom",
		Age:  20,
	}

	reflectTest02(student)
}
