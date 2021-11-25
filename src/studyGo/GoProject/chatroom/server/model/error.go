package model

import "errors"

// 根据业务逻辑需要，自定义的错误信息
var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS    = errors.New("用户已存在")
	ERROR_USER_PWD       = errors.New("密码不正确")
	ERROR_UNMARSHAL      = errors.New("反序列化失败")
)
