package ziface

import "net"

/*
	连接需要实现的方法
*/
type ConnInterface interface {
	// 启动连接，开始工作
	Start()
	// 停止连接，关闭工作
	Close()
	// 获取当前连接的conn（套接字）
	GetTCPConn() *net.TCPConn
	// 获取连接ID
	GetConnID() uint32
	// 获取客户端连接的地址和端口
	RemoteAddr() net.Addr
	// 发送数据给客户端
	Send(data []byte) error
}

//HandleFunc 连接绑定处理业务的函数类型
type HandleFunc func(*net.TCPConn, []byte, int) error
