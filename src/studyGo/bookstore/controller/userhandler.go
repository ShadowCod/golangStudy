package controller

import (
	"net/http"
	"studyGo/bookstore/dao"
	"studyGo/bookstore/model"
	"studyGo/bookstore/utils"
	"text/template"
)

// Login 处理用户登陆的函数
func Login(w http.ResponseWriter, r *http.Request) {
	// 判断是否已经登陆，否则会出现一直增加session的情况
	flag, _ := dao.IsLogin(r)
	if !flag { //未登陆情况下
		// 获取客户端传过来的用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		// 调用userdao中验证用户名|密码
		user, err := dao.CheckUserNameAndPassWord(username, password)
		if err == nil {
			// 正确
			// 生成UUID赋值给session.SessionID
			uuid := utils.CreateUUID()
			// 生成Session实例
			session := &model.Session{
				SessionID: uuid,
				UserName:  user.UserName,
				UserID:    user.ID,
			}
			// 向数据库中添加session
			dao.AddSession(session)
			// 生成cookie,让其与session关联
			http.SetCookie(w, &http.Cookie{
				Name:  "user",
				Value: uuid,
			})
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			// 不正确
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用户名或密码不正确")
		}
	} else {
		GetPageBooksByPrice(w, r)
	}
}

// Regist 处理用户注册的函数
func Regist(w http.ResponseWriter, r *http.Request) {
	// 获取客户端传入的数据
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	// 查询注册的用户名是否已经存在
	res, _ := dao.CheckUserName(username)
	if res { //true表示已经存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在")
	} else {
		// 调用userdao中存储用户的函数
		dao.SaveUser(username, password)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

//CheckName 处理判断用户名的ajax请求
func CheckName(w http.ResponseWriter, r *http.Request) {
	// 获取用户输入的用户名
	username := r.PostFormValue("name")
	res, _ := dao.CheckUserName(username)
	if res { //true表示已经存在
		w.Write([]byte("用户名已存在"))
	} else {
		w.Write([]byte("<font style='color:green'>用户名可用</font>"))
	}
}

// Logout 用户注销
func Logout(w http.ResponseWriter, r *http.Request) {
	// 获取用户cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		// 删除数据库中对应session
		dao.DelSession(cookieValue)
		// 设置cookie失效 MaxAge >0:多少秒后失效;MaxAge=0:会话结束失效;MaxAge<0:立即失效
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	// 直接去首页(调用函数)
	GetPageBooksByPrice(w, r)
}
