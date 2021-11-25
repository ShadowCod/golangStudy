package model

import (
	"net"
	"studyGo/GoProject/chatroom/common/message"
)

// 维护自身,因为在客户端在很多地方会使用到,因此做成全局的
type CurUser struct {
	Conn net.Conn
	message.User
}
