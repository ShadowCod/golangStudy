package main

import (
	"net/http"
	"studyGo/bookstore/controller"
)

func main() {
	// // 创建多路复用器
	// mux := http.NewServeMux()
	// // 指定处理器(映射处理器),注册处理器
	// // http.HandleFunc("/", handler) //使用http中的
	// mux.HandleFunc("/", handler)
	// // 创建路由
	// // http.ListenAndServe(":8080", nil) //使用默认的多路复用器
	// http.ListenAndServe(":8080", mux) //使用自己创建的多路复用器

	// 设置处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	// 直接去main页面  已经被分页handle代替
	// http.HandleFunc("/main", controller.IndexHandler)
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	// 获取所有图书信息
	// http.HandleFunc("/getBooks", controller.GetAllBooks)
	// 添加图书
	// http.HandleFunc("/addBook", controller.AddBook)

	// 去登陆
	http.HandleFunc("/login", controller.Login)
	// 去注册
	http.HandleFunc("/regist", controller.Regist)
	// 处理判断用户名是否正确（客户端发送的ajax）
	http.HandleFunc("/checkUserName", controller.CheckName)
	// 删除图书
	http.HandleFunc("/delBook", controller.DelBook)
	// 去图书修改页面
	http.HandleFunc("/toUpdataPage", controller.ToUpdataPage)
	// 更新或添加图书
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)
	// 获取分页图书信息
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	// 通过价格查询图书
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	// 注销用户
	http.HandleFunc("/logout", controller.Logout)
	// 添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	// 获取购物车信息
	http.HandleFunc("/getCartInfo",controller.GetCartInfo)
	// 清空购物车
	http.HandleFunc("/clearCart",controller.ClearCart)

	http.ListenAndServe(":8080", nil)
}
