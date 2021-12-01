package ziface

/*
	路由的抽象接口
*/
type IRouter interface {
	//在处理业务之前的钩子方法
	PreHandle(request IRequest)
	//处理业务的方法
	Handle(request IRequest)
	//在处理业务之后的钩子方法
	PostHandle(request IRequest)
}