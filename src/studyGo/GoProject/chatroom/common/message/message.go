package message

// 确定一些消息类型
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotufyUserStatusMesType = "NotufyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// 定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
)

type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"` //消息的内容
}

//定义需要的类型
type LoginMes struct {
	UserID  int    `json:"userID"`  //用户id
	UserPwd string `json:"userPwd"` //用户密码
}

type LoginResMes struct {
	Code    int    `json:"code"`    //返回状态码
	UsersID []int  `json:"usersID"` // 增加一个保存在线人员的字段
	Error   string `json:"error"`   //返回的错误信息
}

type RegisterMes struct {
	User User `json:"user"` //类型是User结构体
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //返回状态码
	Error string `json:"error"` //返回的错误信息
}

// 为了配合服务器推送用户状态变化的消息
type NotufyUserStatusMes struct {
	UserID int `json:"userID"`
	Status int `json:"status"`
}

// 增加一个SmsMes，发送的消息
type SmsMes struct {
	Content string `json:"content"` //发送的内容
	User           //使用匿名结构体，继承
}
