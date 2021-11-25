/**
*	socket编程
 */
package main

import (
	"fmt"
	"net"
	"sync"
)

// User 定义的结构体
type User struct {
	Name string      //名字
	ID   string      //唯一的ID
	Msg  chan string //管道
}

// 创建一个全局唯一保存用户的map  不允许同时读写，需要上锁
var allUser = make(map[string]*User)

// 声明一个全局的锁
var lock sync.Locker

// 定义一个全局message管道
var message = make(chan string, 10)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听网络地址失败")
		return
	}
	fmt.Println("服务器启动成功，监听中...")
	defer listener.Close()
	// 启动广播
	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接收连接失败")
			continue
		}
		// 处理成功的连接(使用多线程)
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	// defer conn.Close()
	// 连接进来后创建一个User
	user := &User{
		Name: conn.RemoteAddr().Network(),
		ID:   conn.RemoteAddr().String(),
		Msg:  make(chan string, 10),
	}
	// 将新的连接用户存放到map中
	lock.Lock()
	allUser[user.ID] = user
	lock.Unlock()
	// 将User上线的消息存放到message中
	str := fmt.Sprintf("%s上线", user.ID)
	message <- str
	go writeBackToClient(user, conn)
	// 具体业务
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		if n == 0 { //用户使用ctrl+c退出
			str1 := fmt.Sprintf("%s下线", user.ID)
			message <- str1
			delete(allUser, user.ID)
			conn.Close()
			user.Msg <- "下线"
			return
		}
		fmt.Println("传入的消息:", string(buf[:n]))
		// 获取客户端传入的消息
		message <- string(buf[:n])
	}
}

// 广播
func broadcast() {
	fmt.Println("====>广播go线程监听中")
	defer fmt.Println("广播线程====>退出")
	// 获取message中的信息
	for { //为了让其不退出
		info := <-message
		// fmt.Println("message中的消息:", info)
		// 将信息写入到每个user的msg中
		for _, user := range allUser {
			// 如果msg是非缓冲的（make(chan string)）,会在这里阻塞
			user.Msg <- info
		}
	}
}

func writeBackToClient(u *User, c net.Conn) {
	fmt.Println("====>writeBackToClient线程监听中")
	defer fmt.Println("writeBackToClient线程====>退出")
	for info := range u.Msg { //当msg中没有任何消息时会一直阻塞在这里
		if info == "下线" {
			return
		}
		_, err := c.Write([]byte(info))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("writeBackToClient写回的消息", info)
	}
}
