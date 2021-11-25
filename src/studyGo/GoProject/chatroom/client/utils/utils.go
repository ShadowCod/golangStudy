package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"studyGo/GoProject/chatroom/common/message"
)

// 这里将这些方法关联到结构体中
type Transfer struct {
	// 分析应该有哪些字段
	Conn net.Conn
	Buf  [1024]byte //这是传输时使用的缓冲
}

// 接收数据的封装函数
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 8096)
	// conn.read在conn没有被关闭的情况下，才会阻塞
	// 如果客户端关闭了不会阻塞
	// fmt.Println("conn:", this.Conn)
	_, err = this.Conn.Read(this.Buf[0:4])
	if err != nil {
		fmt.Println("client conn.Read err", err)
		// 自定义错误
		// err = errors.New("read pkg header err")
		return
	}
	// 根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	// 根据pkglen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// fmt.Println("conn.Read err", err)
		// 自定义错误
		// err = errors.New("read pkg body err")
		return
	}
	// 把pkgLen反序列化
	err = json.Unmarshal(this.Buf[:pkgLen], &mes) //注意：要加&
	if err != nil {
		// fmt.Println("json.Unmarshal err", err)
		return
	}
	return
}

// 返回数据的封装函数
func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		return
	}
	// 发送消息本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		return
	}
	return
}
