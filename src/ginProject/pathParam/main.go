package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	获取path(url)参数值	返回值都是string类型
*/
func main(){
	//生成默认路由
	r:=gin.Default()
//	指定函数处理请求	注意：为了url的匹配不冲突 需要修改为/user/:name/:age
	r.GET("/user/:name/:age",func(c *gin.Context){
	//	取得path参数的值
		name:=c.Param("name")
		age:=c.Param("age")
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})
	//注意：这里path路径已经被上面的匹配上了，编译器运行时就会出现抛错
	r.GET("/blog/:year/:month",func(c *gin.Context){
		year:=c.Param("year")
		month:=c.Param("month")
		c.JSON(http.StatusOK,gin.H{
			"year":year,
			"month":month,
		})
	})
//	启动并监听服务
	r.Run(":9090")
}
