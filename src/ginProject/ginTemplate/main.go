package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)
/*
	静态文件：html页面上用到的样式文件.css js文件	图片
*/

func main(){
//	生成默认路由
	r:=gin.Default()
//	加载静态文件需要在解析模版之前	第一个参数是:引用静态文件以啥开头	第二个参数是：去那个文件中找
//	所有以xxx开头的都去statics文件中找
	r.Static("xxx","./statics")
//	在gin框架中添加自定义的函数(需要在解析模版之前注册自定义函数)
	r.SetFuncMap(template.FuncMap{
		"safe":func(s string) template.HTML{
			return template.HTML(s)
		},
	})
	//	解析模版
//	r.LoadHTMLFiles("templates/posts/index.tmpl")
	r.LoadHTMLGlob("templates/**/*")//解析template下面的所有模版
//	指定函数处理请求	注意：处理函数的参数必须是 *gin.Context
	r.GET("/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{//渲染模版	注意：name中的名称应该和模版名称一致
			"context":"tom",
		})
	})

//	指定函数处理请求
	r.GET("/users/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
			"title":"user/index.tmpl",
		})
	})
//启动server并监听端口	注意：必须加上":"
	r.Run(":9090")

}