package main

import (
	"errors"
	"fmt"
)

// 错误处理:加入预警代码，使代码更加健壮
// 在默认情况下，当发生错误后（panic），程序就会退出（崩溃）
// 如果希望，当发生错误后，可以捕获错误，并进行处理，保证程序可以继续执行。还可以在捕获到错误后，给管理员一个提示
// go语言中处理错误使用的是defer、panic、recover
// Go中可以抛出一个panic错误，然后在defer中通过recover捕获这个异常进行处理
func test() {
	// 使用defer和recover来捕获和处理异常
	defer func() {
		err := recover() //recover()内置函数，可以捕获到异常
		if err != nil {  //捕获到异常
			fmt.Println("err:", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

// 自定义错误
func readConf(name string) error {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取错误。")
	}
}

func test02() {
	err := readConf("config.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("后续代码...")
}
func main() {
	// test()
	// fmt.Println("后续代码...")
	// 自定义错误使用
	test02()
}
