package process

import (
	"encoding/json"
	"fmt"
	"studyGo/GoProject/chatroom/client/utils"
	"studyGo/GoProject/chatroom/common/message"
)

type SmsProcess struct {
}

// 发送群聊的方法
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	// 1.创建一个message实例
	var mes message.Message
	mes.Type = message.SmsMesType
	// 2.创建一个SmsMes实例
	var sendCon message.SmsMes
	sendCon.Content = content
	sendCon.UserID = curUser.UserID
	sendCon.UserStatus = curUser.UserStatus
	// 3.将SmeMes序列化
	data, err := json.Marshal(sendCon)
	if err != nil {
		fmt.Println("client smsProcess json.Marshal(sendCon) err", err)
		return
	}
	// 4.将data转换成string赋值给mes.data
	mes.Data = string(data)
	// 5.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("client smsProcess json.Marshal(mes) err", err)
		return
	}
	// 6.发送信息，使用utils中的Transfer
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("客户端发送群聊消息失败", err)
		return
	}
	return
}
