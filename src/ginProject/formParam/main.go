package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	获取form表单提交的参数
*/
func fun1(c *gin.Context){
//渲染模版
	c.HTML(http.StatusOK,"login.html",nil)
}

func fun2(c *gin.Context){
//	方式一：
//	username:=c.PostForm("username")
//	password:=c.PostForm("password")
//	方式二：
//	username:=c.DefaultPostForm("username","somebody")
//	password:=c.DefaultPostForm("password","xxx")
//	方式三：
	username,ok:=c.GetPostForm("username")
	if !ok{
		username="somebody"
	}
	password,ok:=c.GetPostForm("password")
	if !ok{
		password = "xxx"
	}
	//fmt.Println("point")
	//测试能否用取query string的方法获取post表单	结果：取不到表单中的值
	//password:=c.Query("password")
	//fmt.Println("password:",password)
	//c.HTML(http.StatusOK,"home.html",gin.H{
	//		"username":username,
	//		"password":password,
	//	})
	c.JSON(http.StatusOK,gin.H{
				"username":username,
				"password":password,
	})
}

func main(){
//	生成默认路由
	r:=gin.Default()
//	解析模版
	r.LoadHTMLFiles("./login.html","./home.html")
//	指定处理函数处理请求
	r.GET("/login",fun1)
//	指定函数处理post请求
	r.POST("/login",fun2)
//	启动并监听服务
	r.Run(":9090")
}
