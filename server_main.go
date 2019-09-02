package main

import (
	"MyTcpFrame/zinx/iface"
	"MyTcpFrame/zinx/net"
	"fmt"
	"strings"
)
//具体业务应该由使用框架的人传入
type TestRouter struct {
	net.Router
}

// 继承重写
func (r *TestRouter) Handle(req iface.IRequest) {
	data := req.GetData()
	conn := req.GetConn()
	//用户的业务处理逻辑
	buf := []byte(strings.ToUpper(string(data)))
	err := conn.Send(buf)
	if err != nil {
		fmt.Println(err)
	}}

func main() {
	server := net.NewServer("oMaoyu_Tcp 0.0.1")
	server.AddRouter(&TestRouter{})
	server.Server()
}
