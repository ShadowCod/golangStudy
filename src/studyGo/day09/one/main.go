package main

import (
	"fmt"
	"time"
)

// 时间和日期相关
func main() {
	// 1) 获取当前时间
	now := time.Now()
	fmt.Printf("类型:%T,值:%v\n", now, now)

	// 2) 获取其他的日期信息（通过now获取年月日时分秒）
	fmt.Printf("year:%v\n", now.Year())
	fmt.Printf("month:%v\n", int(now.Month())) //获取的是英文需要用int()强转
	fmt.Printf("day:%v\n", now.Day())
	fmt.Printf("Hour:%v\n", now.Hour())
	fmt.Printf("min:%v\n", now.Minute())
	fmt.Printf("Second:%v\n", now.Second())

	// 3) 格式化日期时间
	fmt.Printf(now.Format("2006/01/02 15:04:05")) //数字不能改
	fmt.Println()
	fmt.Printf(now.Format("2006/01/02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	// 4) 时间常量引用	在程序中可用于获取指定时间单位的时间
	n := 100 * time.Millisecond
	fmt.Println(n)

	// 需求，每隔一秒钟打印一个数字
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// 	// time.Sleep(time.Millisecond * 100) //休眠0.1秒
	// 	time.Sleep(time.Second) //时间常量的转换不能为乘以小数
	// }

	// 5) 获取unix时间戳和unixnano时间戳（作用是可以获取随机数字）
	time := now.Unix()
	time1 := now.UnixNano()
	fmt.Println(time, time1)
}
