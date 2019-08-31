package iface

import "net"

type IConnection interface {
	Start() // 读写方法
	Stop()
	Send([]byte)error // 向conn发送数据
	GetConnId() uint32 // 每个链接都有属于自己的id
	GetTcpConn() *net.TCPConn
}

// 定义一个回调  让用户提供 处理用户指定业务

type IConnBlock func(IConnection,[]byte)