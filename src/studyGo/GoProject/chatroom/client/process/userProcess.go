package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"studyGo/GoProject/chatroom/client/utils"
	"studyGo/GoProject/chatroom/common/message"
)

type UserProcess struct {
	// 字段
}

func (u *UserProcess) Login(userID int, pwd string) (err error) {
	// 开始定协议
	// fmt.Printf("userID:%d,pwd:%s\n", userID, pwd)
	// return nil
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.dial err", err)
		return
	}
	// 延时关闭
	defer conn.Close()
	// 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	// 创建一个loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = pwd
	fmt.Println("login loginMes:", loginMes)
	// 将结构体系列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.marshal err", err)
		return
	}
	mes.Data = string(data)
	// 将message结构体系列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err", err)
		return
	}
	// 可以发送消息了
	// 先把data的长度发送给服务器
	// 先获取data的长度，转换成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err", err)
		return
	}
	// fmt.Println("client send length sucess", len(data))
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}
	// 处理返回的消息
	// 创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	// fmt.Println("login tf")
	mes, err = tf.ReadPkg()
	// fmt.Println("login before err:", err)
	if err != nil {
		fmt.Println("readPkg err", err)
		return
	}
	// fmt.Println("login before loginResMes")
	// 将mes的data部分反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}
	if loginResMes.Code == 200 {
		// 初始化CurUser
		curUser.Conn = conn
		curUser.UserID = userID
		curUser.UserStatus = message.UserOnline
		// fmt.Println("登陆成功")
		// 显示当前在线用户的列表
		for _, v := range loginResMes.UsersID {
			if v == userID {
				continue
			}
			// fmt.Printf("用户ID:%d\n", v)
			// 完成客户端的map初始化
			user := &message.User{
				UserID:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		// 这里需要在客户端启动一个协程，该协程保持和服务器端的通讯
		go serverProcessMes(conn)
		// 显示登陆成功的菜单（循环显示）
		for {
			showMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (u *UserProcess) Register(userID int, pwd, name string) (err error) {
	// 连接到后台服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("userProcess register net.Dial err:", err)
		return
	}
	defer conn.Close()

	// 组装需要发送到服务器的数据
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserID = userID
	registerMes.User.UserPwd = pwd
	registerMes.User.UserName = name
	// 将登陆信息序列化赋值给mes.data
	registerMesByte, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("userProcess register json.Marshal registerMes err:", err)
		return
	}
	mes.Data = string(registerMesByte)
	// 将mes序列化方便传输到服务器
	mesByte, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("userProcess register json.Marshal mes err:", err)
		return
	}
	// 发送长度给服务器检验是否丢包
	var pkgLen = uint32(len(mesByte))
	// 因为不能传输数字类型，只能先将数字转换为字节进行传输
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	// 将转换好的表示长度的字节发送给服务器端
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("userProcess register conn.Write buf[:4] err:", err)
		return
	}
	// 如果上面没有报错，则连接稳定。可以发送消息本身
	_, err = conn.Write(mesByte)
	if n != 4 || err != nil {
		fmt.Println("userProcess register conn.Write mesByte err:", err)
		return
	}

	// 发送完毕后等待服务器端返回消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	returnRes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("userProcess register tf.ReadPkg err:", err)
		return
	}
	// 得到返回的数据后将数据反序列化
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(returnRes.Data), &registerResMes)
	if err != nil {
		fmt.Println("userProcess register json.Unmarshal(returnRes) err:", err)
		return
	}
	// 处理返回的详情
	if registerResMes.Code == 200 {
		fmt.Println("注册成功")
	} else {
		fmt.Println(registerResMes.Error)
	}
	return
}
