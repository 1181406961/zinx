package znet

import (
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
	// 绑定的业务处理函数
	handleAPI ziface.HandleFunc
	// 通知并广播当前的连接是否退出
	ExitChan chan bool
}

// init connection
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callbackApi,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
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
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			continue
		}
		// 调用当前链接所绑定的handlerAPI
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID ", c.ConnID, "handler is error ", err)
			break
		}
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

// func (c *Connection) Send(data []byte) error {

// }
