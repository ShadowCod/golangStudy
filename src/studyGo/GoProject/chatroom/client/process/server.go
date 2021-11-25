package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"studyGo/GoProject/chatroom/client/utils"
	"studyGo/GoProject/chatroom/common/message"
)

// 显示登陆成功后的界面
func showMenu() {
	fmt.Printf("---------%d登陆成功--------\n", curUser.UserID)
	fmt.Println("		   1.在线用户列表")
	fmt.Println("		   2.发送信息")
	fmt.Println("		   3.信息列表")
	fmt.Println("		   4.退		出")
	fmt.Println("请选择(1-4):")
	var key int
	var content string
	// 因为总是会使用到SmsProcess实例，如果每次大消息都创建太浪费，因此在此处声明
	sp := &SmsProcess{}
	fmt.Scanln(&key)
	switch key {
	case 1:
		// fmt.Println("查看在线列表")
		outputOnlineUser()
	case 2:
		// 1.输入消息
		fmt.Println("请输入:")
		fmt.Scanf("%s\n", &content)
		// 2.调用发送消息的方法
		sp.SendGroupMes(content)
	case 3:
		fmt.Println("查看信息列表")
	case 4:
		fmt.Println("选择退出")
		os.Exit(0)
	default:
		fmt.Println("输入的选项不正确")
	}
}

// 和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	// 创建一个transfer实例，不停的读取服务器发送的信息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		// 如果读取到消息，进行处理
		// 通过switch来判断不同类型进行不同处理
		switch mes.Type {
		case message.NotufyUserStatusMesType:
			// 1.取出NotufyUserStatusMes
			var nusm message.NotufyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &nusm)
			if err != nil {
				fmt.Println("client server.go json.Unmarshal err", err)
				return
			}
			// 2.将用户信息，状态保存到客户端维护的map(map[int]User)中
			updataUserStatus(&nusm)
		case message.SmsMesType: //处理群发的消息
			outputGroupMes(&mes)
		default:
			fmt.Println("暂时未处理")
		}
		// fmt.Println(mes)
	}

}
