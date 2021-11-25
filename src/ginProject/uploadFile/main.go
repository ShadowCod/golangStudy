package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

/*
	gin框架中上传文件
*/
func main(){
//设置默认路由
	r:=gin.Default()
//	解析
	r.LoadHTMLFiles("./index.html")
//	指定函数处理请求
	r.GET("/upload",func(c *gin.Context){
		//渲染
		c.HTML(http.StatusOK,"index.html",nil)
	})
//	指定函数处理上传文件
//	r.POST("/upload",func(c *gin.Context){
//	//	使用接收方法
//		f,err:=c.FormFile("f")
//		if err !=nil{
//			c.JSON(http.StatusBadRequest,gin.H{
//				"err":err.Error(),
//			})
//		}else{
//		//	设定保存文件的路径
//			dir :=path.Join("./",f.Filename)
//		//	将接收到的文件保存到本地（服务器）
//			c.SaveUploadedFile(f,dir)
//		//	响应
//			c.JSON(http.StatusOK,gin.H{
//				"message":"ok",
//			})
//		}
//	})
//	指定函数处理上传多个文件
	r.POST("/upload",func(c *gin.Context){
	//	解析并将参数放到form中
	form,_:=c.MultipartForm()
	fmt.Println("form:---",form)
	//	取出上传的文件
	files:=form.File["f"]
	fmt.Println("files:---",files)
	//	通过遍历将文件保存到本地（服务器）
	for _,f:=range files{
		//	设定保存文件的路径
		dir :=path.Join("./",f.Filename)
		//	将接收到的文件保存到本地（服务器）
		c.SaveUploadedFile(f,dir)
		}
		//	响应
		c.JSON(http.StatusOK,gin.H{
			"message":"ok",
		})
	})
//	启动并监听服务
	r.Run(":9090")
}
