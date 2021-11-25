package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	gin框架中获取请求中的参数
*/
//Get请求中 URL?后面的是query string参数
func func1(c *gin.Context){
//	获取Get请求参数方式一:
	val := c.Query("name")
	//获取GET请求参数方式二:	第二个参数是设定的默认值
	//val :=c.DefaultQuery("name","noVal")
	//获取Get请求参数方式三:	存在返回(val true)	不存在返回("" false)
	//val,ok:=c.GetQuery("name")
	//if !ok{
	//	val="no value"
	//}
	c.JSON(http.StatusOK,gin.H{
		"name":val,
	})
}

func main(){
	//声明一个默认的路由
	r:=gin.Default()
//	指定函数处理对应的请求
	r.GET("/web",func1)
//	启动服务并监听服务端口
	r.Run(":9090")
}
