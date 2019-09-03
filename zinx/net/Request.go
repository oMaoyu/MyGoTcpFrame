package net

import "MyTcpFrame/zinx/iface"

type Request struct {
	conn iface.IConnection
	msg  iface.IMessage
}

func NewRequest(conn iface.IConnection, msg iface.IMessage) iface.IRequest {
	return &Request{
		conn: conn,
		msg:  msg,
	}
}

func (req *Request) GetConn() iface.IConnection {
	return req.conn
}
func (req *Request) GetMsg() iface.IMessage {
	return req.msg
}

