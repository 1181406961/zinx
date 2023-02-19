package main

import "zinx/znet"

func main() {
	// create
	s := znet.NewServer("[zinx V0.1]")
	// run server
	s.Serve()
}
