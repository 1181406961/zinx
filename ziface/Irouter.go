package ziface
// 路由抽象接口

type IRouter interface{
	// 在处理conn业务之前的hook
	PreHandle(request IRequest)
	// 在处理conn业务的主方法hook
	Handle(request IRequest)
	// 在处理conn业务之后的hook
	PostHandle(request IRequest)
}