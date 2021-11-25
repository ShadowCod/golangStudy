package model

// Cart 购物车结构体
type Cart struct {
	CartID      string      //购物车ID
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int64       //图书总数量
	TotalAmount float64     //图书总金额
	UserID      int         //用户ID
	UserName 	string		//用户名称
}

// GetTotalCount 计算图书总数量
func (c *Cart) GetTotalCount() int64 {
	var totalCount int64
	//变量购物项切片
	for _, v := range c.CartItems {
		totalCount += v.Count
	}
	return totalCount
}

// GetTotalAmount 计算图书总金额
func (c *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	// 遍历购物项切片
	for _, v := range c.CartItems {
		totalAmount += v.GetAmount()
	}
	return totalAmount
}
