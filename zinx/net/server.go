package net

import (
	"MyTcpFrame/zinx/iface"
	"fmt"
	"net"
	"strings"
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
}

func NewServer(name string) iface.IServer {
	return &Server{
		IP:      "0.0.0.0",
		Port:    8888,
		Name:    name,
		Version: "tcp4",
	}
}

func userBlock(conn iface.IConnection, data []byte) {
	//用户的业务处理逻辑
	buf := []byte(strings.ToUpper(string(data)))
	err := conn.Send(buf)
	if err != nil {
		fmt.Println(err)
	}
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
			MyConn := NewConnection(con,cid,userBlock)
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
