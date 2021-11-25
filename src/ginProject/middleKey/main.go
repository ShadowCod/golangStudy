package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
	gin框架中的中间件
*/
/*
	注意：中间件函数必须是gin.HandlerFunc 即：函数的参数是gin.Context
	注意：如果中间件中要使用goroutine必须使用c.Copy()	只能使用拷贝的，不能修改
*/
//中间件函数
func m1(c *gin.Context){
	fmt.Println("m1 in...")
	//计算后面处理函数使用时间
	t:=time.Now()
	//继续执行下一个处理函数
	c.Next()
	//不再执行后面的处理函数 直接返回到调用者
	//c.Abort()
	cost:=time.Since(t)
	fmt.Println("cost:---",cost)
	fmt.Println("m1 out...")
}

//中间件函数
func m2(c *gin.Context){
	fmt.Println("m2 in...")
	//计算后面处理函数使用时间
	//t:=time.Now()
	//继续执行下一个处理函数
	//在中间件插入数据给后面执行的函数中使用
	c.Set("name","tom")
	c.Next()
	//不再执行后面的处理函数 直接返回到调用者
	//c.Abort()
	//cost:=time.Since(t)
	//fmt.Println("cost:---",cost)
	fmt.Println("m2 out...")
}

func main() {
//	注意：默认创建的路由注册了两个中间键logger()和Recovery()
	//	创建一个全新的路由，没有注册任何中间键 r:=gin.New()
	r:=gin.Default()
	//如果需要注册一个全局的中间键，即r路由中的所有函数都需要调用的中间件
	//r.Use(m1)
	//给单个处理函数注册中间件
	//r.GET("/index",m1, func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"mes":"index",
	//	})
	//})
//	使用全局中间件时不需要特地添加中间件了
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"mes":"index",
		})
	})
//	给某个路由组注册中间件	注意：会先执行注册了的全局中间件
	indexGroup:=r.Group("/index")
	{
		indexGroup.Use(m2)
		indexGroup.GET("/1",func(c *gin.Context){
			//取出中间件中传递的数据
			name,ok:=c.Get("name")
			if !ok{
				name = "无效"
			}
			c.JSON(http.StatusOK,gin.H{
				"mes":name,
			})
		})
	}
//	启动并监听服务,默认监听8080端口
	r.Run(":9090")
}
