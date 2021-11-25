package main

import (
	"fmt"
	"net"
	"studyGo/GoProject/chatroom/server/model"
	"time"
)

// // 接收数据的封装函数
// func readPkg(conn net.Conn) (mes message.Message, err error) {
// 	buf := make([]byte, 8096)
// 	// conn.read在conn没有被关闭的情况下，才会阻塞
// 	// 如果客户端关闭了不会阻塞
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		// fmt.Println("conn.Read err", err)
// 		// 自定义错误
// 		// err = errors.New("read pkg header err")
// 		return
// 	}
// 	// 根据buf[:4]转成一个uint32类型
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[0:4])
// 	// 根据pkglen读取消息内容
// 	n, err := conn.Read(buf[:pkgLen])
// 	if uint32(n) != pkgLen || err != nil {
// 		// fmt.Println("conn.Read err", err)
// 		// 自定义错误
// 		// err = errors.New("read pkg body err")
// 		return
// 	}
// 	// 把pkgLen反序列化
// 	err = json.Unmarshal(buf[:pkgLen], &mes) //注意：要加&
// 	if err != nil {
// 		// fmt.Println("json.Unmarshal err", err)
// 		return
// 	}
// 	return
// }

// // 返回数据的封装函数
// func writePkg(conn net.Conn, data []byte) (err error) {
// 	// 先发送长度给对方
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[:4], pkgLen)
// 	n, err := conn.Write(buf[:4])
// 	if n != 4 || err != nil {
// 		return
// 	}
// 	// 发送消息本身
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		return
// 	}
// 	return
// }

// 编写一个函数专门处理登陆请求
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	// 核心处理代码...
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes) //要写&
// 	if err != nil {
// 		fmt.Println("json.Unmarshal err", err)
// 		return
// 	}
// 	// 声明一个resMes
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType
// 	// 声明一个loginResMes
// 	var loginResMes message.LoginResMes
// 	//	判断是否合法
// 	if loginMes.UserID == 100 && loginMes.UserPwd == "123456" {
// 		fmt.Println("用户合法")
// 		loginResMes.Code = 200
// 	} else {
// 		fmt.Println("用户不合法")
// 		loginResMes.Code = 500
// 		loginResMes.Error = "用户不存在"
// 	}
// 	// 系列化loginResMes
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal err", err)
// 		return
// 	}
// 	resMes.Data = string(data)
// 	// 序列化resMes
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal err", err)
// 		return
// 	}
// 	// 返回数据准备完毕
// 	// 发送data的功能封装到一个函数中
// 	err = writePkg(conn, data)
// 	return
// }

// 编写一个serverProcessMes函数
// 功能：根据客户端发送消息种类不同，决定调用那个函数来 处理
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		// 处理登陆
// 		err = serverProcessLogin(conn, mes)
// 	default:
// 		fmt.Println("消息类型不存在，无法处理...")
// 	}
// 	return
// }

func goProcess(conn net.Conn) {
	// 延时关闭conn
	defer conn.Close()
	// 调用总控
	p := &Processor{
		Conn: conn,
	}
	err := p.Process2()
	if err != nil {
		return
	}
}

// 完成UserDao初始化函数
func initUserDao() {
	// 这里的pool本身就是一个全局变量
	model.MyUserDao = model.CreateUserDao(pool)
}

// 初始化函数
func init() {
	// 当服务器启动时，就去初始化
	initPool("127.0.0.1:6379", 16, 0, 300*time.Second)
	//初始化连接池
	initUserDao()
}

func main() {
	// 提示信息
	fmt.Println("服务器启动,等待连接...")
	// 监听
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listen.Close()
	// 监听成功就等待连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accpet err", err)
		} else {
			//连接成功则启动协程和客户端保持通讯
			go goProcess(conn)
		}
	}
}
