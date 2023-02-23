package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	// name
	Name string
	// address version
	IPVersion string
	// listen ip
	IP string
	// listen port
	Port int
	// 当前Server添加一个router，server注册的连接对应的处理业务
	Router ziface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP:%s, Port %d, is starting\n", s.IP, s.Port)
	go func() {
		// 获取一个tcp连接
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		// 监听地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen, ", s.IPVersion, " err ", err)
			return
		}
		fmt.Println("start Zinx server success, ", s.Name, " success listening……")
		var cid uint32
		cid = 0
		// 阻塞并等待连接
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			//将处理新连接的业务方法和conn进行绑定，得到我们的链接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			// 启动当前的链接业务处理
			go dealConn.Start()
		}
	}()

}
func (s *Server) Stop() {
	// stop server
}
func (s *Server) Serve() {
	// start server
	s.Start()
	// do something after start

	// block, because Start is async
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ!!")

}

// 初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}
