package main

import (
	"fmt"
	"net"
	"time"
)

// mock client
func main() {
	fmt.Println("client start……")
	// connect server
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	for {
		_, err = conn.Write([]byte("Hello Zinx V0.2……"))
		if err != nil {
			fmt.Println("write conn err ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error ", err)
			return
		}
		fmt.Printf("server call back: %s, cnt=%d\n", buf, cnt)
		//  sleep
		time.Sleep(1 * time.Second)
	}
}
