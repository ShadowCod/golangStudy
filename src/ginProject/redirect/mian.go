package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	gin框架中的重定向
*/
// 直接重定向到其他地址
func Redirect(c *gin.Context){
	c.Redirect(http.StatusMovedPermanently,"https://www.sogo.com")
}

// 重定向到另一个处理函数	这种形式需要使用到r，不能将处理函数单独写
func funcA(c *gin.Context,r *gin.Engine){
//将url中的path路径修改为/b
	c.Request.URL.Path = "/b"
//	继续执行
	r.HandleContext(c)
}

func funcB(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"b",
	})
}

func main(){
	//生成默认路由
	r:=gin.Default()
	//指定处理函数处理请求
	r.GET("/redirect",Redirect)
	//处理funA请求
	r.GET("/a",func(c *gin.Context){
	//	修改url中的path
		c.Request.URL.Path = "/b"
	//	继续执行处理函数
		r.HandleContext(c)
	})
	////处理funB请求
	//r.GET("/b",func(c *gin.Context){
	//	c.JSON(http.StatusOK,gin.H{
	//		"message":"b",
	//	})
	//})
	//r.GET("/a",funcA)  修改路径的处理函数不能单独提到外面写
	r.GET("/b",funcB)
	//启动并监听服务
	r.Run(":9090")
}
