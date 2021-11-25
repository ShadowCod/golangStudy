package dao

import (
	"net/http"
	"studyGo/bookstore/model"
	"studyGo/bookstore/utils"
)

// AddSession 添加session
func AddSession(s *model.Session) (err error) {
	// 编写sql语句
	sqlStr := "insert into sessions (session_id,username,user_id) values (?,?,?)"
	// 执行sql语句
	_, err = utils.Db.Exec(sqlStr, s.SessionID, s.UserName, s.UserID)
	return
}

// DelSession 删除session
func DelSession(id string) (err error) {
	// 编写sql语句
	sqlStr := "delete from sessions where session_id = ?"
	// 执行sql语句
	_, err = utils.Db.Exec(sqlStr, id)
	return
}

// GetSession 查询session
func GetSession(id string) (s *model.Session, err error) {
	// 编写sql语句
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	// // 执行sql语句
	// row := utils.Db.QueryRow(sqlStr, id)
	// 使用预编译的方式
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	// 执行预编译
	row := inStmt.QueryRow(id)
	// 生成session实例
	s = &model.Session{}
	err = row.Scan(&s.SessionID, &s.UserName, &s.UserID)
	return
}

// IsLogin 判断是否已经登陆	不能写到utils包下面，会构成循环调用的错误
func IsLogin(r *http.Request) (bool, *model.Session) {
	// 获取r中的cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		// 获取cookie中的value
		cookieValue := cookie.Value
		// 查询session数据库
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			return true, session
		}
	}
	return false, nil
}
