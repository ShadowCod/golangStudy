package utils

import (
	"encoding/json"
	"goZinx/ziface"
	"io/ioutil"
)

type GlobalObj struct {
	//server
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	//zinx
	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:           "ZinxServer",
		Version:        "v4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        3,
		MaxPackageSize: 512,
	}
	GlobalObject.Reload()
}
