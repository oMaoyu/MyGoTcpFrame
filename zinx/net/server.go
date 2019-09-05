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
	//Router iface.IRouter
	//配置路由群  根据客户端和服务端相同的id去执行对应路由方法
	Routers *Routers
}

func NewServer() iface.IServer {
	return &Server{
		IP:      MyConfig.IP,
		Port:    MyConfig.Port,
		Name:    MyConfig.Name,
		Version: MyConfig.Version,
		//Router:  &Router{},
		Routers: NewRouters(),
	}
}

func (s *Server) Start() {
	add := fmt.Sprintf("%s:%d", s.IP, s.Port)
	tcpAdd, err := net.ResolveTCPAddr(s.Version, add)
	if err != nil {
		fmt.Println(err)
		return
	}
	listent, err := net.ListenTCP(s.Version, tcpAdd)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Routers.StartWorkerPool()
	var cid uint32
	cid = 0

	go func() {
		for {
			// 监听
			con, err := listent.AcceptTCP()
			if err != nil {
				fmt.Println(err)
				continue
			}
			// 使用自己封装的conn
			MyConn := NewConnection(con, cid, s.Routers)
			cid++
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
func (s *Server) AddRouter(key uint32,router iface.IRouter) {
	s.Routers.AddRouter(key,router)
}
