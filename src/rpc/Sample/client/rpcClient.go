package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()
	var r string
	err = client.Call("frist_rpc.Say", "你好", &r)
	fmt.Println("rpc======>返回")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
