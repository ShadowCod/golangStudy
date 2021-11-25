package dao

import (
	"strconv"
	"studyGo/bookstore/model"
	"studyGo/bookstore/utils"
)

// GetAllBooks 获取数据库中的所有图书信息
func GetAllBooks() (books []*model.Book, err error) {
	// 书写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	// 执行sql语句
	rows, err := utils.Db.Query(sqlStr)
	// 遍历row
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := &model.Book{}
		// 给book字段赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return
}

// AddBook 向数据库中添加图书
func AddBook(book *model.Book) (err error) {
	//写sql语句
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	// 执行sql语句
	_, err = utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	return
}

// DelBook 删除图书
func DelBook(id string) (err error) {
	// 写sql语句
	sqlStr := "delete from books where id = ?"
	// 执行sql语句
	_, err = utils.Db.Exec(sqlStr, id)
	return
}

// OneBook 获取一本图书的信息
func OneBook(id string) (book *model.Book, err error) {
	// 写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id =?"
	// 执行sql语句
	row := utils.Db.QueryRow(sqlStr, id)
	// 生成model.Book实例
	book = &model.Book{}
	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if err != nil {
		return nil, err
	}
	return book, err
}

// UpdataBook 更新图书信息
func UpdataBook(book *model.Book) (err error) {
	// 写sql语句
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id =?"
	// 执行sql
	_, err = utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	return
}

// GetPageBooks 获取带分页的图书信息
func GetPageBooks(pageNo string) (page *model.Page, err error) {
	IpageNo, _ := strconv.ParseInt(pageNo, 10, 0)
	// 获取数据库中图书的总记录数
	sqlStr := "select count(*) from books"
	// 执行sql语句
	row := utils.Db.QueryRow(sqlStr)
	// 获取查询到的条数
	var totalCount int64
	row.Scan(&totalCount)
	// 设置每页显示的条数	可以通过参数传入
	var pageSize int64 = 4
	// 计算出总页数
	var totalPageNo int64
	if totalCount%pageSize == 0 {
		totalPageNo = totalCount / pageSize
	} else {
		totalPageNo = totalCount/pageSize + 1
	}
	// 获取当前页中的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	// 执行
	rows, err := utils.Db.Query(sqlStr2, (IpageNo-1)*pageSize, pageSize)
	var books []*model.Book
	// 获取每一条数据
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	page = &model.Page{
		Books:       books,
		PageNo:      IpageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalCount,
	}
	return
}

// GetPageBooksByPrice 根据价格区间查询图书信息
func GetPageBooksByPrice(min, max, pageNo string) (page *model.Page, err error) {
	// 查询一共有多少条的语句
	sqlStr := "select count(*) from books where price between ? and ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, min, max)
	// 获得总条数
	var totalCount int64
	row.Scan(&totalCount)
	// 写sql语句
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	// 执行
	var pageSize int64 = 4
	IpageNo, _ := strconv.ParseInt(pageNo, 10, 0)
	rows, err := utils.Db.Query(sqlStr2, min, max, (IpageNo-1)*pageSize, pageSize)
	// 遍历查询得到的数据
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	// 计算得到总页数
	var totalPage int64
	if totalCount%pageSize == 0 {
		totalPage = totalCount / pageSize
	} else {
		totalPage = totalCount/pageSize + 1
	}
	page = &model.Page{
		Books:       books,
		PageNo:      IpageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPage,
		TotalRecord: totalCount,
	}
	return
}
