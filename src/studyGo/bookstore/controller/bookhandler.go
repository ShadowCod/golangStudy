package controller

import (
	"net/http"
	"strconv"
	"studyGo/bookstore/dao"
	"studyGo/bookstore/model"
	"text/template"
)

//IndexHandler 创建处理器函数------去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 获取客户端参数的值
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	// 查询数据库信息
	page, _ := dao.GetPageBooks(pageNo)
	// 解析模版
	t := template.Must(template.ParseFiles("views/index.html"))
	// 执行
	t.Execute(w, page)
}

// GetAllBooks 获取所有图书信息
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetAllBooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

// AddBook 添加图书
func AddBook(w http.ResponseWriter, r *http.Request) {
	// 获取前台传入的参数值
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	priceStr := r.PostFormValue("price")
	salesStr := r.PostFormValue("sales")
	stockStr := r.PostFormValue("stock")
	// 需要将数字类型进行转换
	price, _ := strconv.ParseFloat(priceStr, 64)
	sales, _ := strconv.ParseInt(salesStr, 10, 0)
	stock, _ := strconv.ParseInt(stockStr, 10, 0)
	// 生成model.Book实例
	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   price,
		Sales:   int(sales),
		Stock:   int(stock),
		ImgPath: "static/img/default.jpg",
	}
	// 调用添加图书的方法
	dao.AddBook(book)
	// 调用GetAllBooks处理器函数，再一次查询数据库
	GetAllBooks(w, r)
}

// DelBook 删除图书
func DelBook(w http.ResponseWriter, r *http.Request) {
	// 获取客户端参数的值(此处不是提交表单，不能使用PostFormValue)
	idStr := r.FormValue("bookID")
	// 删除数据的方法
	dao.DelBook(idStr)
	// 调用GetAllBooks处理器函数，再一次查询数据库
	GetAllBooks(w, r)
}

// ToUpdataPage 去更新信息界面或者添加界面
func ToUpdataPage(w http.ResponseWriter, r *http.Request) {
	// 获取参数值
	id := r.FormValue("bookID")
	if id != "" { // 更新
		// 查询信息
		book, _ := dao.OneBook(id)
		// 解析模版
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		// 执行
		t.Execute(w, book)
	} else { //添加图书
		// 解析模版
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		// 执行
		t.Execute(w, "")
	}
}

// UpdateOrAddBook 更新或添加图书信息
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	// 获取前台传入的参数值
	idStr := r.PostFormValue("bookID")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	priceStr := r.PostFormValue("price")
	salesStr := r.PostFormValue("sales")
	stockStr := r.PostFormValue("stock")
	// 需要将数字类型进行转换
	price, _ := strconv.ParseFloat(priceStr, 64)
	id, _ := strconv.ParseInt(idStr, 10, 0)
	sales, _ := strconv.ParseInt(salesStr, 10, 0)
	stock, _ := strconv.ParseInt(stockStr, 10, 0)
	// 生成model.Book实例
	book := &model.Book{
		ID:      int(id),
		Title:   title,
		Author:  author,
		Price:   price,
		Sales:   int(sales),
		Stock:   int(stock),
		ImgPath: "static/img/default.jpg",
	}
	if book.ID > 0 {
		// 更新数据
		dao.UpdataBook(book)
	} else {
		// 添加图书
		dao.AddBook(book)
	}
	// 调用GetAllBooks处理器函数，再一次查询数据库
	GetAllBooks(w, r)
}

// GetPageBooks 获取分页的图书信息
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	// 获取客户端参数的值
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	// 查询数据库信息
	page, _ := dao.GetPageBooks(pageNo)
	// 解析模版
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	// 执行
	t.Execute(w, page)
}

//GetPageBooksByPrice 根据价格查询书籍
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	// 获取客户端参数的值
	min := r.FormValue("min")
	max := r.FormValue("max")
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if min != "" && max != "" {
		// 查询数据库
		page, _ = dao.GetPageBooksByPrice(min, max, pageNo)
		page.Min = min
		page.Max = max
	} else {
		// 查询数据库信息
		page, _ = dao.GetPageBooks(pageNo)
	}
	// 判断是否已经登陆
	flag, session := dao.IsLogin(r)
	// 判断是否存在cookie
	if flag {
		page.UserName = session.UserName
		page.IsLogin = true
	}
	// 解析模版
	t := template.Must(template.ParseFiles("views/index.html"))
	// 执行
	t.Execute(w, page)
}
