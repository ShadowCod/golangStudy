package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// redis操作
// func main() {
// 	//链接redis
// 	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
// 	if err != nil {
// 		fmt.Println("连接失败", err)
// 		return
// 	}
// 	fmt.Println("------连接成功------")
// 	defer conn.Close() //关闭连接
// 	// 通过连接写入
// 	// _, err = conn.Do("set", "name", "tom")
// 	// if err != nil {
// 	// 	fmt.Println("写入失败", err)
// 	// 	return
// 	// }
// 	fmt.Println("------写入成功------")
// 	// 通过连接读取
// 	r, err := redis.String(conn.Do("get", "name"))
// 	if err != nil {
// 		fmt.Println("读取失败", err)
// 		return
// 	}
// 	// 因为返回的r是interface{}类型，需要使用断言,redis提供了方法转换
// 	// nameStr := r.(string)
// 	fmt.Println("读取的值:", r)
// }

func main() {
	// 连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()
	// 写入
	_, err = conn.Do("hmset", "stu", "name", "tom", "age", 18)
	if err != nil {
		fmt.Println("写入失败:", err)
		return
	}
	// 读取	注意加S
	r, err := redis.Strings(conn.Do("hmget", "stu", "name", "age"))
	if err != nil {
		fmt.Println("读取失败:", err)
		return
	}
	// 返回的值是个切片
	for _, v := range r {
		fmt.Println(v)
	}
}
