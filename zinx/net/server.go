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
	go func() {
		for {
			con,err := listent.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}
			go func() {
				defer con.Close()
				for {
					buf := make([]byte, 512)
					cnt,err := con.Read(buf)
					if err != nil {
						return
					}
					buf = []byte(strings.ToUpper(string(buf[:cnt])))
					fmt.Println(string(buf))
					_,err = con.Write(buf)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}()

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
