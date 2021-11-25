package process

import (
	"encoding/json"
	"fmt"
	"net"
	"studyGo/GoProject/chatroom/common/message"
	"studyGo/GoProject/chatroom/server/utils"
)

type SmsProcess struct {
	// 暂时不需要字段
}

func (sp *SmsProcess) SendGroupMes(mes *message.Message) {
	// 取出发送的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("服务器反序列化群发消息失败", err)
		return
	}
	// 将整个要发送的信息再序列化
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("服务器序列化群发消息失败", err)
		return
	}
	// 遍历服务器端的onlineUsers
	for id, up := range userMgr.onlineUsers {
		// 过滤掉自己
		if id == smsMes.UserID {
			continue
		}
		// 将消息转发
		sp.SendMesToOnlineUser(data, up.Conn)
	}

}

func (sp *SmsProcess) SendMesToOnlineUser(data []byte, conn net.Conn) {
	// 创建一个Transfer实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发群消息失败", err)
	}
}
