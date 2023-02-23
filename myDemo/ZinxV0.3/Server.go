package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

// ping test自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping...ping...ping error")
	}

}

// Post Handle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}

func main() {
	// 创建一个router
	s := znet.NewServer("[zinx V0.3]")
	// 添加一个自定义的router
	s.AddRouter(&PingRouter{})
	// 启动server
	s.Serve()
}
