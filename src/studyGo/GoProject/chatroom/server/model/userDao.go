package model

import (
	"encoding/json"
	"fmt"
	"studyGo/GoProject/chatroom/common/message"

	"github.com/garyburd/redigo/redis"
)

// 在服务器启动后就初始化一个UserDao实例
// 做成全局变量，需要用时直接使用
var (
	MyUserDao *UserDao
)

// 定义一个UserDao结构体
// 完成对User结构体的各种操作
type UserDao struct {
	Pool *redis.Pool
}

// 使用工厂模式创建一个UserDao的实例
func CreateUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool: pool,
	}
	return
}

// 根据用户ID,返回一个User实例 + err
func (u *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// 通过给的ID，去redis查询这个用户
	res, err := redis.String(conn.Do("hget", "users", id))
	// fmt.Println("res:", res)
	if err != nil {
		// 出现错误
		if err == redis.ErrNil {
			// 表示在users哈希中，没有找到
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	// 声明一个user
	user = &User{}
	// 需要吧res反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		err = ERROR_UNMARSHAL
		return
	}
	return
}

// 完成登陆的校验
func (u *UserDao) Login(userID int, userPwd string) (user *User, err error) {
	// 从useDao的连接池中取出一个连接
	conn := u.Pool.Get()
	defer conn.Close()
	user, err = u.getUserById(conn, userID)
	// fmt.Println("server login:", user)
	if err != nil {
		// fmt.Println("login err", err)
		return
	}
	// 证明用户存在，但是还需要判断密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 注册功能
func (u *UserDao) Register(user *message.User) (err error) {
	// 取得数据库连接
	conn := u.Pool.Get()
	defer conn.Close()
	// 查询是否存在该userID
	_, err = u.getUserById(conn, user.UserID)
	// fmt.Println("server login:", user)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	// 没有则可以完成注册
	userByte, err := json.Marshal(user)
	_, err = conn.Do("hset", "users", user.UserID, string(userByte))
	if err != nil {
		fmt.Println("server userDao Register conn.Do err:", err)
		return
	}
	return
}
