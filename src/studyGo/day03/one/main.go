package main

import "fmt"

func main() {
	//流程控制，跳出for循环
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		fmt.Println("break for")
	// 		break //跳出for循环
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("for over")
	//流程控制，跳过for循环
	// for j := 0; j < 10; j++ {
	// 	if j == 3 {
	// 		fmt.Println("jump")
	// 		continue //继续下一次循环
	// 	}
	// 	fmt.Printf("j:%d\n", j)
	// }
	// fmt.Println("for_over")

	//switch 简化大量判断
	// var i = 3
	// switch i {
	// case 1:
	// 	fmt.Println("大拇指")
	// case 2:
	// 	fmt.Println("食指")
	// case 3:
	// 	fmt.Println("中指")
	// case 4:
	// 	fmt.Println("无名指")
	// case 5:
	// 	fmt.Println("小拇指")
	// default:
	// 	fmt.Println("无效参数")
	// }

	//switch 变种一
	// switch n := 1; n {
	// case 1:
	// 	fmt.Println("小拇指")
	// default:
	// 	fmt.Println("无效参数")
	// }

	//switch 变种二
	// switch n := 1; n {
	// case 1, 2, 3:
	// 	fmt.Println("小拇指")
	// default:
	// 	fmt.Println("无效参数")
	// }

	//switch 变种三
	// var n = 1
	// switch {
	// case n > 1:
	// 	fmt.Println("小拇指")
	// default:
	// 	fmt.Println("无效参数")
	// }

	//跳出多重循环
	// for i := 0; i < 5; i++ {
	// 	var flag = false
	// 	for j := 0; j < 3; j++ {
	// 		if j == 2 {
	// 			flag = true
	// 			break
	// 		}
	// 		fmt.Printf("%d-%d\n", i, j)
	// 	}
	// 	if flag {
	// 		break
	// 	}
	// }

	//使用goto+labal跳出
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				goto XX
			}
			fmt.Printf("%d-%d\n", i, j)
		}
	}
XX:
	fmt.Println("here")
}
