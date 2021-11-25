package znet

import (
	"fmt"
	"net"
)

// Server 定义一个服务器结构体
type Server struct {
	// 服务器名称
	Name string
	// 服务器IP类型
	IPVersion string
	// 服务器ip地址
	IP string
	// 服务器使用端口
	Port int
}

// CreateServer 创建Server的函数
func CreateServer(name string) (s *Server) {
	s = &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      8999,
	}
	return
}

// Start 实现接口的start方法
func (s *Server) Start() {
	fmt.Printf("%s,starting\n", s.Name)
	// 需要使用一个goroutine来一直等待连接
	go func() {
		// 1.解析TCP地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}
		// 2.监听解析出来的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen addr err:", err)
			return
		}
		// 3.等待客户端连接（使用for循环等待）
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("listener accept err:", err)
				continue
			}
			// 做一个简单的最大512字节的回应(每一个连接开启一个协程专门负责回应)
			go func() {
				// 创建一个存放信息的
				info := make([]byte, 512)
				for {
					// 读取发送过来的信息
					cnt, err := conn.Read(info)
					fmt.Println("server read ", cnt, info)
					if err != nil {
						fmt.Println("read err,err:", err)
						continue
					}
					// 回写
					_, err = conn.Write(info[:cnt])
					if err != nil {
						fmt.Println("write err,err:", err)
						continue
					}
				}
			}()
		}
	}()

}

// Stop 实现接口的stop方法
func (s *Server) Stop() {

}

// Serve 实现接口的server方法
func (s *Server) Serve() {
	// 调用start方法
	s.Start()
	// TODO 后续增加功能

	//	等待
	select {}
}
