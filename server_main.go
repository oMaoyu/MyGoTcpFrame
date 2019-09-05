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
	data := req.GetMsg().GetData()
	conn := req.GetConn()
	//用户的业务处理逻辑
	buf := []byte(strings.ToUpper(string(data)))
	_,err := conn.Send(buf,200)
	if err != nil {
		fmt.Println(err)
	}
}
//=======
type DemoRouter struct {
	net.Router
}

func (r *DemoRouter) Handle(req iface.IRequest) {
	conn := req.GetConn()
	//用户的业务处理逻辑
	buf := []byte("????")
	_,err := conn.Send(buf,200)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	server := net.NewServer()
	defer server.Stop()
	server.AddRouter(1,&TestRouter{})
	server.AddRouter(2,&DemoRouter{})
	server.Server()
}
