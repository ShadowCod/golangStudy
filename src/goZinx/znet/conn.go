package znet

import (
	"fmt"
	"goZinx/ziface"
	"net"
)

//Connection 连接模块属性
type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	IsClosed  bool
	HandleAPI ziface.HandleFunc
	ExitChan  chan bool
}

//CreateConnection 创建连接模块的方法
func CreateConnection(conn *net.TCPConn, connID uint32, callBackAPI ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		IsClosed:  false,
		HandleAPI: callBackAPI,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

//StartRead start方法中需要的读取信息方法
func (c *Connection) StartRead() {
	fmt.Println("start read")
	defer fmt.Println("connID=",c.ConnID,"Reader is exit ,remote addr is ",c.RemoteAddr().String())
	defer c.Stop()
	for {
		//读取客户端的数据到buf中
		buf:=make([]byte,512)
		cnt,err:=c.Conn.Read(buf)
		if err !=nil{
			fmt.Println("recv buf err",err)
			continue
		}

		//调用当前连接所绑定的HandleAPI
		if err = c.HandleAPI(c.Conn,buf,cnt);err !=nil{
			fmt.Println("ConnID",c.ConnID,"handle is err",err)
			break
		}
	}
}

//Start 启动连接的方法
func (c *Connection) Start() {
	fmt.Println("conn start()...,connID:", c.ConnID)
	// 启动从当前的连接中读取数据
	go c.StartRead()
	// TODO
}

//Stop 停止连接的方法
func (c *Connection) Stop() {
	fmt.Println("conn stop()...,connID:", c.ConnID)
	//如果已经关闭就不必继续执行了
	if c.IsClosed {
		return
	}
	// 关闭连接
	c.IsClosed = true
	// 关闭socket连接
	c.Conn.Close()
	// 回收channel
	close(c.ExitChan)
}

//GetTCPConn 获取连接的方法
func (c *Connection) GetTCPConn() *net.TCPConn {
	return c.Conn
}

//GetConnID 获取连接ID的方法
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//RemoteAddr 获取客户端的地址、端口方法
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//Send 给客户端发送消息的方法
func (c *Connection) Send(data []byte) error {
	return nil
}
