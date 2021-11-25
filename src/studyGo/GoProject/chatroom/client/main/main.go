package main

import (
	"fmt"
	"os"
	"studyGo/GoProject/chatroom/client/process"
)

// 定义全局变量用户|密码
var (
	userID int
	pwd    string
	name   string
)

// 注意：编译时要退到src下，使用绝对路径进行编译
func main() {
	// 定义一个变量来存储用户选项
	var choess int64
	// 定义一个变量来存储是否关闭主界面的标识
	// var flag = true
	// 通过循环来控制让用户选择的界面
	// 通过flag让界面一直循环，选择了其中一个正确选项就不再进行循环
	for {
		fmt.Println("----------界面------------")
		fmt.Println("		1 登陆用户")
		fmt.Println("		2 注册用户")
		fmt.Println("		3 退	出")
		fmt.Print("请输入选项(1-3):")
		fmt.Scanf("%d\n", &choess)
		switch choess {
		case 1:
			fmt.Println("请输入用户")
			fmt.Scanln(&userID)  //必须带'\n'，否则会出问题
			fmt.Println("请输入密码") //必须带'\n'，否则会出问题
			fmt.Scanln(&pwd)
			// flag = false
			// 完成登陆功能
			// 1.创建一个UserProcess实例
			up := &process.UserProcess{}
			up.Login(userID, pwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &pwd)
			fmt.Println("输入用户的昵称:")
			fmt.Scanf("%s\n", &name)
			// 1.创建一个UserProcess实例,完成注册
			up := &process.UserProcess{}
			up.Register(userID, pwd, name)
			// flag = false
		case 3:
			fmt.Println("退	出")
			os.Exit(0)
			// flag = false
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
	// 根据用户的选项显示新的
	// if choess == 1 {
	// 说明用户要登陆
	// fmt.Println("请输入用户")
	// fmt.Scanln(&userID)  //必须带'\n'，否则会出问题
	// fmt.Println("请输入密码") //必须带'\n'，否则会出问题
	// fmt.Scanln(&pwd)
	// 将功能写到server文件
	// 因为login和main都在main包中，所以可以直接调用
	// fmt.Println("client main pwd:", pwd)
	// err := login(userID, pwd)
	// if err != nil {
	// 	fmt.Println("login err", err)
	// } else {
	// 	fmt.Println("login success")
	// }
	// } else if choess == 2 {
	// 	fmt.Println("user registration")
	// }
}
