package net

import "MyTcpFrame/zinx/iface"

type Request struct {
	conn iface.IConnection
	data []byte
	len  uint32
}

func NewRequest(conn iface.IConnection, data []byte, len uint32) iface.IRequest {
	return &Request{
		conn: conn,
		data: data,
		len:  len,
	}
}

func (req *Request) GetConn() iface.IConnection {
	return req.conn
}
func (req *Request) GetData() []byte {
	return req.data
}
func (req *Request) GetLen() uint32 {
	return req.len
}
