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
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP:%s, Port %d, is starting\n", s.IP, s.Port)
	go func() {
		// get TCP addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		// listen addr
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen, ", s.IPVersion, " err ", err)
			return
		}
		fmt.Println("start Zinx server success, ", s.Name, " success listening……")
		// block wait, handle read and write
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			// connected
			go func() {
				for {
					// create buffer 512 bytes
					buf := make([]byte, 512)
					// read length
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
					// if err != nill read content write back
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err ", err)
						continue
					}
				}
			}()
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

// init server module method
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
