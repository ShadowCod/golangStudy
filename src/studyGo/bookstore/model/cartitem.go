package model

// CartItem 结构体
type CartItem struct {
	CartItemID int64   //购物项的ID
	Book       *Book   //购物项中的图书信息
	Count      int64   //购物项中图书的数量
	Amount     float64 //购物项中图书的金额小计
	CartID     string  //购物车的ID
}

//GetAmount 获取金额小计
func (c *CartItem) GetAmount() float64 {
	// 获取当前购物项中图书的价格
	price := c.Book.Price
	return float64(c.Count) * price
}
