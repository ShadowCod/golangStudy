package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	gin框架中绑定参数的方法
*/
/*
	使用该方法的原因：
		每次接收一个参数就需要定义一个变量接收过于频繁麻烦，使用该方法会自动执行定义变量接收参数
*/
//注意：需要使用tag将对应的字段定制化为传入参数时的key
type UserInfo struct{
	Username string	`form:"username"`
	Password string	`form:"password"`
}

func main(){
//	生成默认路由器
	r:=gin.Default()
//	指定函数处理请求（绑定query string中传过来的参数值）
//	r.GET("/bind",func(c *gin.Context){
//		u:=&UserInfo{}
//	//	使用gin框架的方法绑定参数(注意：方法和函数中想要修改传入的参数需要传递指针类型)
//		err:=c.ShouldBind(u)//注意：该方法会使用反射取结构体中的值，结构体中的字段首字母要大写
//		if err!=nil{
//			fmt.Printf("bind fail err:%v\n",err)
//			return
//		}
//	//	返回响应
//		c.JSON(http.StatusOK,u)
//	})
//	指定函数处理请求（绑定post form中传递的参数值）
//	r.POST("/form",func(c *gin.Context){
//	//	定义一个结构体类型的变量
//		u:=&UserInfo{}
//	//	将参数值绑定到变量上
//		err:=c.ShouldBind(u)
//		if err !=nil{
//			c.JSON(http.StatusOK,gin.H{
//				"message":http.StatusBadRequest,
//				"err":err.Error(),
//			})
//		}else{
//			c.JSON(http.StatusOK,gin.H{
//				"message":http.StatusOK,
//				"value":u,
//			})
//		}
//	})
//	指定函数处理请求（绑定post json参数值）
	r.POST("/json",func(c *gin.Context){
		//	定义一个结构体类型的变量
		u:=&UserInfo{}
		//	将参数值绑定到变量上
		err:=c.ShouldBind(u)
		if err !=nil{
			c.JSON(http.StatusOK,gin.H{
				"message":http.StatusBadRequest,
				"err":err.Error(),
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"message":http.StatusOK,
				"value":u,
			})
		}
	})
//	启动并监听服务
	r.Run(":9090")
}
