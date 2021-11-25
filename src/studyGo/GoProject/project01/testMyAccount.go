package main

import "fmt"

func main() {
	key := "" //如果放在for循环中，每次都要创建和销毁
	money := 0.0
	note := ""
	balance := 0.0
	detail := "类型\t余额\t具体\t说 明"
	for { //无限循环
		fmt.Println("\n-----------界面----------")
		fmt.Println("       1 收支明细")
		fmt.Println("       2 登记收入")
		fmt.Println("       3 登记支出")
		fmt.Println("       4 退出界面")
		fmt.Println()
		fmt.Println("请选择1-4:")
		fmt.Scanln(&key)
		// 如果使用switch则需要定义一个标志，根据标志判断是否退出
		if key == "1" {
			fmt.Println("----------收支详情----------")
			fmt.Println(detail)
		} else if key == "2" {
			fmt.Println("收入金额:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("收入说明:")
			fmt.Scanln(&note)
			// 将记录拼接到详情中去
			detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		} else if key == "3" {
			fmt.Println("支出金额:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
			} else {
				balance -= money
				fmt.Println("支出说明:")
				fmt.Scanln(&note)
				// 将记录拼接到详情中去
				detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			}
		} else if key == "4" {
			fmt.Println("即将退出...")
			break
		} else {
			fmt.Println("选项错误")
		}
	}
	fmt.Println("退出成功")
}
