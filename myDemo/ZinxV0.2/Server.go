package main

import "zinx/znet"

func main() {
	// create
	s := znet.NewServer("[zinx V0.2]")
	// run server
	s.Serve()
}
