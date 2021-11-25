package main

import (
	"fmt"
	"io"
	"net"
	"studyGo/GoProject/chatroom/common/message"
	"studyGo/GoProject/chatroom/server/process"
	"studyGo/GoProject/chatroom/server/utils"
)

// 创建一个processor结构体
type Processor struct {
	Conn net.Conn
}

// 编写一个serverProcessMes函数
// 功能：根据客户端发送消息种类不同，决定调用那个函数来 处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登陆
		// 创建一个UserProcess结构体
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		// server端的conn可能是不一样的所以不能提取出去
		sp := &process.SmsProcess{}
		sp.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) Process2() (err error) {
	for {
		// 将读取数据包功能封装成一个函数
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		// fmt.Println("server main mes:", mes)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端已退出")
				return err
			} else {
				fmt.Println("server readPkg err", err)
				return err
			}
		}
		// fmt.Println("mes:", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			// fmt.Println("serverProcessMes err", err)
			return err
		}
	}
}
