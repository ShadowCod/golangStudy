package process

import (
	"fmt"
	"studyGo/GoProject/chatroom/client/model"
	"studyGo/GoProject/chatroom/common/message"
)

// 客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var curUser model.CurUser //在用户登陆成功后完成初始化

// 在客户端显示当前在线用户
func outputOnlineUser() {
	fmt.Println("当前在线列表:")
	for id, _ := range onlineUsers {
		fmt.Printf("用户id:%d\n", id)
	}
}

// 编写函数处理状态信息
func updataUserStatus(n *message.NotufyUserStatusMes) {
	// 判断原来是否有这个id
	user, ok := onlineUsers[n.UserID]
	if !ok {
		user = &message.User{
			UserID: n.UserID,
		}
	}
	user.UserStatus = n.Status
	// 处理完后调用展示方法
	// outputOnlineUser()
	fmt.Printf("用户%d上线\n", n.UserID)
}
