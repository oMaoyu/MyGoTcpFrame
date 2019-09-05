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

	ConnMan iface.IConnManager
	// 连接开始执行函数
	ConnStateFunc func(connection iface.IConnection)
	// 连接结束执行函数
	ConnStopFunc func(connection iface.IConnection)

}

func NewServer() iface.IServer {
	return &Server{
		IP:      MyConfig.IP,
		Port:    MyConfig.Port,
		Name:    MyConfig.Name,
		Version: MyConfig.Version,
		//Router:  &Router{},
		Routers: NewRouters(),
		ConnMan: NewConnManager(),
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
			// 控制最大连接数
			if s.ConnMan.GetConnCount() >= MyConfig.Worker.ConnSize{
				_ = con.Close()
				continue
			}

			// 使用自己封装的conn
			MyConn := NewConnection(con, cid, s.Routers,s)
			//建立连接时,添加到连接管理器当中
			s.ConnMan.AddConn(MyConn)
			cid++
			go MyConn.Start()

		}
	}()
}
func (s *Server) Stop() {
	s.GetConnMan().ClearConn()
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
func (s *Server)GetConnMan()iface.IConnManager{
	return s.ConnMan
}
// 注册
func (s *Server)RegisterStart(f func(connection iface.IConnection)){
	s.ConnStateFunc = f
}
func (s *Server)RegisterStop(f func(connection iface.IConnection)) {
	s.ConnStopFunc = f
}
//执行
func (s *Server)ConnStartRun(connection iface.IConnection){
	if s.ConnStateFunc == nil {
		return
	}
	s.ConnStateFunc(connection)

}
func (s *Server)ConnStopRun(connection iface.IConnection){
	if s.ConnStopFunc == nil {
		return
	}
	s.ConnStopFunc(connection)
}

