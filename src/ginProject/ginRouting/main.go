package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	gin框架中的路由与路由组
*/
func main(){
//	创建默认路由
	r:=gin.Default()
//	使用any接收任意方式的请求
	r.Any("/any",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"mes":"any",
		})
	})
//	设置用户请求不存在的路径返回的信息
	r.NoRoute(func(c *gin.Context){
		c.JSON(http.StatusNotFound,gin.H{
			"mes":"invalid path",
		})})
//	路由组的使用
//eg	视频网站
	r.GET("/video/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"mes":"/video/index",
		})
	})
	r.GET("/video/home", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"mes":"/video/home",
		})
	})
//	上面写法会随着越来越多导致有很多重复的路径
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/video1", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"mes":"/video/video1",
			})
		})
		videoGroup.POST("/video2", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"mes":"/video/video2",
			})
		})
		//路由组也支持嵌套
		pageGroup:=videoGroup.Group("/page")
		pageGroup.GET("/1", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"mes":"/video/page/1",
			})
		})
	}
//	启动服务并监听
	r.Run(":9090")
}
