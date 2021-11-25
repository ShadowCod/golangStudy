package dao

import (
	"database/sql"
	"studyGo/bookstore/model"
	"studyGo/bookstore/utils"
)

//CheckUserNameAndPassWord 登陆验证方法
func CheckUserNameAndPassWord(name string, pwd string) (*model.User, error) {
	// 写sql语句
	sqlStr := "select id,username,password  from users where username =? and password = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, name, pwd)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.PassWord)
	return user, err
}

//CheckUserName 验证用户名
func CheckUserName(name string) (bool, error) {
	// 写sql语句
	sqlStr := "select id from users where username = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, name)
	var userNo int
	err := row.Scan(&userNo)
	if err == sql.ErrNoRows {
		return false, err
	}
	return true, nil
}

// SaveUser 向数据库中插入信息
func SaveUser(name, pwd string) error {
	// 写sql语句
	sqlStr := "insert into users(username,password) values(?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, name, pwd)
	if err != nil {
		return err
	}
	return nil
}
