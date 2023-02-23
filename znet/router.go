package znet

import "zinx/ziface"

// 实现router时，先嵌入BaseRouter，用户根据需要进行重新定制
type BaseRouter struct{}

// 在处理conn业务之前的hook
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

// 在处理conn业务的主方法hook
func (br *BaseRouter) Handle(request ziface.IRequest) {}

// 在处理conn业务之后的hook
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
