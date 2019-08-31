package main

import (
	"MyTcpFrame/zinx/net"
)

func main() {
	server := net.NewServer("oMaoyu_Tcp 0.0.1")
	server.Server()
}
