package znet

import "goZinx/ziface"

type Request struct{
	//已经和客户端建立好的连接
	conn ziface.ConnInterface
	//客户端请求的数据
	data []byte
}

func (r *Request) GetConnection() ziface.ConnInterface{
	return r.conn
}

func (r *Request) GetData() []byte{
	return r.data
}