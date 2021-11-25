package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 定义一个全局的连接池
var pool *redis.Pool

// 使用init方法进行初始化
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

// 使用连接池
func main() {
	// 先从pool中取出连接
	conn := pool.Get()
	defer conn.Close() //及时关闭放回连接池
	// 使用连接
	res, _ := redis.String(conn.Do("hget", "stu", "name"))
	fmt.Println(res)
	// 如果使用了pool.Close()关闭连接池，则无法在取出连接conn
}
