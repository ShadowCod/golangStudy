package process

import (
	"encoding/json"
	"fmt"
	"studyGo/GoProject/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	// 显示即刻
	// 1.反序列化mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("展示发送消息失败", err)
		return
	}
	// 显示信息
	info := fmt.Sprintf("用户%d\t:%s\n", smsMes.UserID, smsMes.Content)
	fmt.Println(info)
}
