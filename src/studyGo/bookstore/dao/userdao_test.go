package dao

import (
	"fmt"
	"studyGo/bookstore/model"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("测试userdao中的函数")
	// t.Run("测试用户名|密码:", testLogin)
	// t.Run("测试用户名:", testRegist)
	// t.Run("测试用存储:", testSave)
	// t.Run("测试查询所有图书信息:", testGetAllBook)
	// t.Run("测试添加图书:", testAddBook)
	// t.Run("测试删除图书:", testDelBook)
	// t.Run("测试获取一本图书信息:", testOneBook)
	// t.Run("测试更新图书信息:", testUpdataBook)
	// t.Run("测试获取分页图书信息:", testGetPageBook)
	// t.Run("测试根据价格获取分页图书信息:", testGetPageBooksByPrice)
	// t.Run("测试添加session:", testAddSession)
	// t.Run("测试删除session:", testDelSession)
	// t.Run("测试向购物车中加入:", testAddCart)
	// t.Run("测试向购物车中加入:", testGetCartItemsByCartID)
	t.Run("测试向购物车中加入:", testGetCartByUserID)
}

// testLogin 测试登陆
func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassWord("lisi", "666")
	fmt.Println("获取的用户信息:", user)
}

// testRegist 测试注册
func testRegist(t *testing.T) {
	res, _ := CheckUserName("admin")
	fmt.Println("验证结果:", res)
}

// testSave 测试存储
func testSave(t *testing.T) {
	SaveUser("lisi", "666")
}

// testGetAllBook 查询所有图书
func testGetAllBook(t *testing.T) {
	books, _ := GetAllBooks()
	fmt.Println(books)
}

// testAddBook 添加图书
func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "红楼梦",
		Author:  "曹雪芹",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "static/img/default.jpg",
	}
	AddBook(book)
}

// testDelBook 删除图书
func testDelBook(t *testing.T) {
	DelBook("33")
}

// testOneBook  获取一本图书
func testOneBook(t *testing.T) {
	book, _ := OneBook("30")
	fmt.Println(book)
}

// testUpdataBook 更新图书信息
func testUpdataBook(t *testing.T) {
	book := &model.Book{
		ID:      1,
		Title:   "解忧杂货店",
		Author:  "东野圭吾",
		Price:   30.3,
		Sales:   100,
		Stock:   100,
		ImgPath: "static/img/default.jpg",
	}
	UpdataBook(book)
}

// testGetPageBook 获取分页图书信息
func testGetPageBook(t *testing.T) {
	page, _ := GetPageBooks("1")
	fmt.Println(page)
}

// testGetPageBooksByPrice 测试根据价格查询图书信息
func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("20", "30", "1")
	for _, v := range page.Books {
		fmt.Println("书籍信息:", v)
	}
}

// testAddSession
func testAddSession(t *testing.T) {
	session := &model.Session{
		SessionID: "one",
		UserName:  "李四",
		UserID:    1,
	}
	AddSession(session)
}

//	testDelSession
func testDelSession(t *testing.T) {
	DelSession("one")
}

//	testAddCart
func testAddCart(t *testing.T) {
	// 设置要买的一本书
	book := &model.Book{
		ID:    1,
		Price: 30.00,
	}
	// 创建购物项切片
	var cartItemSlice []*model.CartItem
	// 设置购物项
	cartItem := &model.CartItem{
		Book:   book,
		Count:  1,
		CartID: "1",
	}
	cartItemSlice = append(cartItemSlice, cartItem)
	// 声明一个cart结构体
	cart := &model.Cart{
		CartID:    "1",
		CartItems: cartItemSlice,
		UserID:    1,
	}
	AddCart(cart)
}

// testGetCartItemByBookID
func testGetCartItemByBookID(t *testing.T) {
	c, _ := GetCartItemByBookID("1")
	fmt.Println(c)
	fmt.Println("count", c.Count)
}

// testGetCartItemsByCartID
func testGetCartItemsByCartID(t *testing.T) {
	c, _ := GetCartItemsByCartID("1")
	for _, v := range c {
		fmt.Println(v)
	}
}

// testGetCartByUserID
func testGetCartByUserID(t *testing.T) {
	c, _ := GetCartByUserID("1")
	fmt.Println(c)
}
