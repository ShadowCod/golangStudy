package model

// Session 结构体定义
type Session struct {
	SessionID string //使用UUID生成
	UserName  string //账号名称
	UserID    int    //账号的ID
}
