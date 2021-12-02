package ziface

// Interface 定义一个接口类型
type IServer interface {
	// 开始服务
	Start()
	// 停止服务
	Stop()
	// 运行服务
	Serve()
	// 添加路由
	AddRouter(router IRouter)
}
