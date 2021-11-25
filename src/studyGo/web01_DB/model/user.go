package model

import (
	"fmt"
	"studyGo/web01_DB/utils"
)

//User 字段应该和数据表中的对应
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

//AddUser 添加User的方法一
func (user *User) AddUser() (err error) {
	//写sql语句
	sql := "insert into users (username,password,email) values(?,?,?)"
	// 预编译
	inStmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println("预编译出现问题:", err)
		return
	}
	// 执行
	_, err = inStmt.Exec("admin", "123456", "admin@123.com")
	if err != nil {
		fmt.Println("执行sql语句出现问题:", err)
	}
	return
}

// AddUser2 添加User的方法二
func (user *User) AddUser2() (err error) {
	// 1.写sql语句
	sql := "insert into user(username,password,email) values(?,?,?)"
	// 2.执行语句
	_, err = utils.Db.Exec(sql, "admin", "123456", "admin@123.com")
	if err != nil {
		fmt.Println("执行sql语句出现问题:", err)
		return
	}
	return
}

//GetUserByID 根据用户ID从数据库中查询一条记录
func (user *User) GetUserByID() (err error) {
	// 写sql语句
	sqlStr := "select id,username,password,email from users where id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, 1)
	// 扫描
	var id int
	var username string
	var password string
	var email string
	err = row.Scan(&id, &username, &password, &email)
	if err != nil {
		fmt.Println("扫描出现问题")
		return
	}
	fmt.Println(id, username, password, email)
	return
}

// GetUsers 获取数据库所有的记录
func (user *User) GetUsers() (err error) {
	// 写sql语句
	sqlStr := "select id,username,password,email from users "
	// 执行语句
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return
	}
	// 扫描rows
	for rows.Next() {
		var id int
		var username string
		var password string
		var email string
		err = rows.Scan(&id, &username, &password, &email)
		if err != nil {
			fmt.Println("扫描出现问题")
			return
		}
		fmt.Println(id, username, password, email)
	}
	return
}
