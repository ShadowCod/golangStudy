package message

// 定义一个用户的结构体
type User struct {
	// 确定字段信息
	UserID     int    `json:"userID"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
}
