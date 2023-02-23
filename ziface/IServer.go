package ziface

type IServer interface {
	// start server
	Start()
	// stop server
	Stop()
	// run server
	Serve()
	// 路由功能，给当前的服务注册一个路由方法，给客户端使用
	AddRouter(router IRouter)
}
