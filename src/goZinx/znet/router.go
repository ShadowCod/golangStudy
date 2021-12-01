package znet

import "goZinx/ziface"

type BaseRouter struct{

}


//在处理业务之前的钩子方法
func (bs *BaseRouter)PreHandle(request ziface.IRequest){}
//处理业务的方法
func (bs *BaseRouter) Handle(request ziface.IRequest){}
//在处理业务之后的钩子方法
func (bs *BaseRouter) PostHandle(request ziface.IRequest){}