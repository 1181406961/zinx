package ziface

import "net"

type IConnection interface {
	// start connection
	Start()
	// stop connection
	Stop()
	// get bind socket conn
	GetTCPConnection() *net.TCPConn
	// get connection id
	GetConnID() uint32
	// get client tcp status ip port
	RemoteAddr() net.Addr
	// send data to client
	Send(data []byte) error
}
// define handle method
type HandleFunc func(*net.TCPConn, []byte,int) error