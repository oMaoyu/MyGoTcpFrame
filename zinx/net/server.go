package net

import (
	"MyTcpFrame/zinx/iface"
	"fmt"
	"net"
)

type Server struct {
	// 地址
	IP string
	// 端口
	Port uint32
	// 名字
	Name string
	//版本 tcp4  tcp6
	Version string
	//配置路由
	Router iface.IRouter
}

func NewServer(name string) iface.IServer {
	return &Server{
		IP:      "0.0.0.0",
		Port:    8888,
		Name:    name,
		Version: "tcp4",
		Router: &Router{},
	}
}

func userBlock(request iface.IRequest) {

}


func (s *Server) Start() {
	add := fmt.Sprintf("%s:%d",s.IP,s.Port)
	tcpAdd,err := net.ResolveTCPAddr(s.Version,add)
	if err != nil {
		fmt.Println(err)
		return
	}
	listent,err := net.ListenTCP(s.Version,tcpAdd)
	if err != nil {
		fmt.Println(err)
		return
	}
	var cid uint32
	cid = 0

	go func() {
		for {
			// 监听
			con,err := listent.AcceptTCP()
			if err != nil {
				fmt.Println(err)
				return
			}
			// 使用自己封装的conn
			MyConn := NewConnection(con,cid,s.Router)
			cid ++
			go MyConn.Start()


		}
	}()
}
func (s *Server) Stop() {

}

func (s *Server) Server() {
	s.Start()
	for {
		;
	}
}
func (s *Server)AddRouter(router iface.IRouter){
	s.Router = router
}