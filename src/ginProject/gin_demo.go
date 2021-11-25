package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"hello golang",
	})
}

func main(){
	//创建一个默认的路由引擎
	r:=gin.Default()
//	指定访问方法对应的出来方法
	r.GET("/hello",sayHello)
//	RESTful 指的是一系列规范	get获取信息，post新建，put更新，delete删除
	r.GET("/book",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method":"GET",
		})
	})
	r.POST("/book",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method":"POST",
		})
	})
	r.PUT("/book",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method":"PUT",
		})
	})
	r.DELETE("/book",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method":"DELETE",
		})
	})
//	启动
	r.Run(":9091")
}
