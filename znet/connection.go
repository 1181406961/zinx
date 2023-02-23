package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
)

// 对原生的Connection进行封装
type Connection struct {
	// 最原始的Conn连接
	Conn *net.TCPConn
	// 唯一ID
	ConnID uint32
	// 链接状态
	isClosed bool

	// 通知并广播当前的连接是否退出
	ExitChan chan bool
	// 该连接处理的方法Router
	Router ziface.IRouter
}

// init connection
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// 链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID= ", c.ConnID, " Reader is exit,remote addr is ", c.Conn.RemoteAddr().String())
	defer c.Stop()
	for {
		// 读取客户端端数据道buffer中，目前最大512字节
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			continue
		}
		// 得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			data: buf,
		}
		// 执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}

}
func (c *Connection) Start() {
	fmt.Println("Conn Start()..ConnID= ", c.ConnID)
	// 启动从当前链接读数据的业务
	go c.StartReader()
}
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID= ", c.ConnID)
	if c.isClosed {
		return
	}
	c.isClosed = true
	// close connection
	c.Conn.Close()
	close(c.ExitChan)
}
func (c *Connection) GetConnID() uint32 {
	fmt.Println("Conn Start()..ConnID= ", c.ConnID)
	return c.ConnID
}
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) Send(data []byte) error {
	return errors.New("Send not implement")
}
