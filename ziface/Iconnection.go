package ziface

import "net"

// 对连接的封装
type IConnection interface {
	// 启动当前连接处理
	Start()
	// 关闭当前连接
	Stop()
	// 获取绑定的原始连接
	GetTCPConnection() *net.TCPConn
	// 获取连接唯一ID
	GetConnID() uint32
	// 获取连接地址状态
	RemoteAddr() net.Addr
	// 发送数据给客户端
	Send(data []byte) error
}
