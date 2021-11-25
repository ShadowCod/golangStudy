package dao

import (
	"studyGo/bookstore/model"
	"studyGo/bookstore/utils"
)

// AddCartItem 插入购物项
func AddCartItem(c *model.CartItem) (err error) {
	// 写sql语句
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	// 执行sql语句
	_, err = utils.Db.Exec(sqlStr, c.Count, c.GetAmount(), c.Book.ID, c.CartID)
	return
}

// GetCartItemByBookID 根据图书ID查找购物项
func GetCartItemByBookID(id string) (c *model.CartItem, err error) {
	// 写sql语句
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, id)
	// 将数据库中的值赋值给结构体
	c = &model.CartItem{}
	err = row.Scan(&c.CartItemID, &c.Count, &c.Amount, &c.CartID)
	return
}

// GetCartItemsByCartID 根据购物车ID查所有购物项
func GetCartItemsByCartID(id string) (c []*model.CartItem, err error) {
	// 写sql语句
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	// 执行sql语句
	rows, err := utils.Db.Query(sqlStr, id)
	// 将查询到的多行数据放入切片中
	for rows.Next() {
		var bookID string
		item := &model.CartItem{}
		rows.Scan(&item.CartItemID, &item.Count, &item.Amount, &bookID, &item.CartID)
		book, _ := OneBook(bookID)
		item.Book = book
		c = append(c, item)
	}
	return
}

// GetCartItemByBookIDAndCartID 根据图书ID和用户ID查询购物项
func GetCartItemByBookIDAndCartID(bookID, cartID string) (c *model.CartItem, err error) {
	//fmt.Printf("book_id:%v,cart_id:%v\n",bookID, cartID)
	// 编写sql语句	注意：需要用到什么就查询什么
	sqlStr := "select id,count,amount from cart_items where book_id = ? and cart_id = ?"
	// 执行sql语句
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	//fmt.Println("row:---",row)
	c = &model.CartItem{}
	err = row.Scan(&c.CartItemID, &c.Count, &c.Amount)
	c.CartID = cartID
	if err != nil {
		return nil, err
	}
	// 根据图书的ID查询图书的信息
	book, _ := OneBook(bookID)
	// 将book设置到购物项中
	c.Book = book
	return
}

// UpdateCartItem 更新购物项信息
func UpdateCartItem(c *model.CartItem) (err error) {
	// 编写sql语句
	sqlStr := "update cart_items set count=?,amount=? where id = ?"
	// 执行sql语句
	_, err = utils.Db.Exec(sqlStr, c.Count, c.GetAmount(), c.CartItemID)
	return
}

// UpdateBookCount 根据购物项更新图书数量、金额信息
func UpdateBookCount(c *model.CartItem) (err error) {
	// 写sql
	sqlStr := "update cart_items set count=?,amount=? where book_id = ? and cart_id = ?"
	// 执行
	_, err = utils.Db.Exec(sqlStr, c.Count,c.GetAmount(), c.Book.ID, c.CartID)
	return
}
