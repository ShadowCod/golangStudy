package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"studyGo/bookstore/dao"
	"studyGo/bookstore/model"
	"studyGo/bookstore/utils"
)

// AddBook2Cart 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	// 判断是否登陆
	flag, session := dao.IsLogin(r)
	if !flag{
		w.Write([]byte("请先登陆"))
	}else{
	// 获取客户端参数值
	bookID := r.FormValue("bookID")
	// fmt.Println("要添加的图书ID是:", bookID)
	// 查询图书信息
	book, _ := dao.OneBook(bookID)
	// 获取用户id
	userID := session.UserID
	// fmt.Println("userID:", userID)
	// 根据userID获取购物车信息
	cart, _ := dao.GetCartByUserID(strconv.Itoa(userID))
	if cart != nil {
		cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
		if cartItem != nil {
			// 购物车的购物项中有这本图书
			// cartItem.Count = cartItem.Count + 1  //不能使用上面的购物项，因为购物车中的购物项切片也需要修改
			// 获取购物车切片
			cts := cart.CartItems
			// 遍历得到每一项
			for _, v := range cts {
				// 找到当前购物项
				if v.Book.ID == cartItem.Book.ID {
					// 将购物项中的图书数量修改
					v.Count = v.Count + 1
					// 将数据库中购物项的图书数量修改
					dao.UpdateBookCount(v)
				}
			}
		} else {
			// 购物车的购物项中没有有这本图书，需要在购物项中添加
			// 1.创建新的购物项
			newCartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cart.CartID,
			}
			// 将购物项添加到当前购物车切片中
			cart.CartItems = append(cart.CartItems, newCartItem)
			// 将购物项添加到数据库中
			dao.AddCartItem(newCartItem)
		}
		// 无论之前购物车中是否有图书对应的购物项都要更新购物车
		dao.UpdateCart(cart)
	} else {
		// 没有购物车,需要创建购物车
		// 1.创建购物项
		// 生成购物车的id
		cartID := utils.CreateUUID()
		newCart := &model.Cart{
			CartID: cartID,
			UserID: userID,
		}
		// 2.创建购物车中的购物项
		var cartItems []*model.CartItem
		newCartItem := &model.CartItem{
			Book:   book,
			Count:  1,
			CartID: cartID,
		}
		// 将购物项添加到切片中
		cartItems = append(cartItems, newCartItem)
		// 将切片添加到购物车中
		newCart.CartItems = cartItems
		// 将购物车保存到数据库
		dao.AddCart(newCart)
	}
	w.Write([]byte("将"+book.Title+"加入到购物车"))
	}
}

// GetCartInfo	根据用户ID获取购物车详情
func GetCartInfo(w http.ResponseWriter,r *http.Request){
//是否登陆
	flag,session :=dao.IsLogin(r)
	if flag{
	//	从session中获取UserID
		userID :=session.UserID
	//	使用dao中的方法获取购物车信息
		cart,_ := dao.GetCartByUserID(strconv.Itoa(userID))
		if cart !=nil{
			cart.UserName = session.UserName
		//	解析模版
			t:=template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//	渲染模版
			t.Execute(w,cart)
		}else{//该用户还没有购物车
			c:=&model.Cart{}
			c.UserName = session.UserName
			//	解析模版
			t:=template.Must(template.ParseFiles("views/pages/cart/cart.html"))
			//	渲染模版
			t.Execute(w,c)
		}
	}else{
		//没登陆就提示登陆
		w.Write([]byte("请先登陆"))
	}
}
// 清空购物车
func ClearCart(w http.ResponseWriter,r *http.Request){
//判断是否登陆
	flag,session:=dao.IsLogin(r)
	if flag{//已经登陆
	//获取购物车信息
		cart,_:=dao.GetCartByUserID(strconv.Itoa(session.UserID))
		if cart!=nil{//有购物车信息

		}else{//没有购物车信息

		}
	}else{//没有登陆
		w.Write([]byte("请先登陆"))
	}
}