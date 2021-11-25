package dao

import (
	"studyGo/bookstore/model"
	"studyGo/bookstore/utils"
)

// AddCart 插入购物车数据
func AddCart(c *model.Cart) (err error) {
	// 写sql语句
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	// 执行sql语句
	_, err = utils.Db.Exec(sqlStr, c.CartID, c.GetTotalCount(), c.GetTotalAmount(), c.UserID)
	if err != nil {
		return
	}
	// 获取购物车中所有的购物项
	cartItems := c.CartItems
	// 遍历得到每个购物项
	for _, v := range cartItems {
		// 将购物项添加到数据库中
		AddCartItem(v)
	}
	return
}

// GetCartByUserID 根据用户的ID从数据库中查询对应的购物车
func GetCartByUserID(id string) (c *model.Cart, err error) {
	// 写sql语句
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, id)
	// 将查询到的数据库数据赋值给结构体
	c = &model.Cart{}
	err = row.Scan(&c.CartID, &c.TotalCount, &c.TotalAmount, &c.UserID)
	if err != nil {
		return nil, err
	}
	// 获取所有购物项
	cartItems, err := GetCartItemsByCartID(c.CartID)
	c.CartItems = cartItems
	return
}

// UpdateCart 更新购物车中图书的总数量和总金额
func UpdateCart(c *model.Cart) error {
	// 编写sql语句
	sqlStr := "update carts set total_count=?,total_amount=? where id =?"
	// 执行sql语句
	_, err := utils.Db.Exec(sqlStr, c.GetTotalCount(), c.GetTotalAmount(), c.CartID)
	if err != nil {
		return err
	}
	return nil
}
