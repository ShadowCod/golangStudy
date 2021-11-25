package process

import (
	"encoding/json"
	"fmt"
	"net"
	"studyGo/GoProject/chatroom/common/message"
	"studyGo/GoProject/chatroom/server/model"
	"studyGo/GoProject/chatroom/server/utils"
)

type UserProcess struct {
	// 分析字段
	Conn net.Conn
	// 增加一个字段表示该conn是谁的
	UserID int
}

// 编写通知所有在线用户的方法
// 注意：是参数这个用户通知其他在线用户
func (this *UserProcess) NotifyOthersOnlineUser(userID int) {
	// 遍历所有在线，然后一个一个发送
	for id, u := range userMgr.onlineUsers {
		if id == userID {
			continue
		}
		// 开始通知,u是
		// fmt.Println(u)
		u.NotifyStatus(userID)
	}
}

// 通知函数
func (this *UserProcess) NotifyStatus(userID int) {
	//组装消息
	var mes message.Message
	mes.Type = message.NotufyUserStatusMesType
	var sendMes message.NotufyUserStatusMes
	sendMes.UserID = userID
	sendMes.Status = message.UserOnline
	// 序列化sendMes
	data, err := json.Marshal(sendMes)
	if err != nil {
		return
	}
	// 序列化mes
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}
	// 使用conn发送信息
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		return
	}
}

// 编写一个函数专门处理登陆请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 核心处理代码...
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes) //要写&
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}
	// 声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	// 声明一个loginResMes
	var loginResMes message.LoginResMes
	// 到redis数据库查询
	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPwd)
	fmt.Println("user:", user)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误(redis未启动)"
		}
	} else {
		loginResMes.Code = 200
		this.UserID = loginMes.UserID
		// 这里表示后台验证通过，即用户登陆成功
		userMgr.AddOnlineUser(this)
		// 登陆成功向在线发送上线信息
		this.NotifyOthersOnlineUser(loginMes.UserID)
		// 将当前在线用户的id放入到loginResMes.UsersID中
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersID = append(loginResMes.UsersID, id)
		}
	}
	// //	判断是否合法
	// if loginMes.UserID == 100 && loginMes.UserPwd == "123456" {
	// 	fmt.Println("用户合法")
	// 	loginResMes.Code = 200
	// } else {
	// 	fmt.Println("用户不合法")
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "用户不存在"
	// }
	// 序列化loginResMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	resMes.Data = string(data)
	// 序列化resMes
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	// 返回数据准备完毕
	// 发送data的功能封装到一个函数中
	// 因为使用了分成模式，要先创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 将mes中的data部分反序列化
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("server userProcess ServerProcessRegister json.Unmarshal err:", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	// 操作数据库完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	// 将组装好的返回值序列化
	regRes, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("server userProcess ServerProcessRegister json.Marshal err:", err)
		return
	}
	resMes.Data = string(regRes)
	resData, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("server userProcess ServerProcessRegister json.Marshal(resMes) err:", err)
		return
	}
	// 将结果发送给客户端
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(resData)
	return
}
