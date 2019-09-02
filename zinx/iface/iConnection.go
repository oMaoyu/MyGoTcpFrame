package iface

import "net"

type IConnection interface {
	Start() // 读写方法
	Stop()
	Send([]byte)error // 向conn发送数据
	GetConnId() uint32 // 每个链接都有属于自己的id
	GetTcpConn() *net.TCPConn
}
