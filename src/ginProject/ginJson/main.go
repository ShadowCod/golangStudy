package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	gin框架模版的渲染
*/
//注意：使用tag来对结构体定制化操作	后面的要用双引号包裹起来
type Msg struct{
	Name string `json:"name"`
	Age int	`json:"age"`
	Say string `json:"say"`
}
func structJson(c *gin.Context){
	msg:=&Msg{
		Name:"jerry",
		Age:20,
		Say:"hello golang",
	}
	c.JSON(http.StatusOK,msg)
}

func main(){
//	生成默认路由
	router := gin.Default()
//	关联请求对应的处理函数
	router.GET("/json",func(c *gin.Context){
		//方式一：使用map[string]interface{}返回
		//data:=map[string]interface{}{
		//	"name":"tom",
		//	"age":18,
		//	"say":"hello world!",
		//}
		//注意：gin框架中已经封装了一个gin.H	这个类型就是map[string]interface{}类型
		data:=gin.H{
				"name":"tom",
				"age":18,
				"say":"hello world!",
		}
		c.JSON(http.StatusOK,data) //渲染模版
	})

	router.GET("/another_json",structJson)

	router.Run(":9090")
}
