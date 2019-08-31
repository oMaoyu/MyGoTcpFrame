package net

import (
	"MyTcpFrame/zinx/iface"
	"fmt"
	"net"
)

type Connection struct {
	conn     *net.TCPConn
	connId   uint32
	isClosed bool
	block    iface.IConnBlock
}

// 实现接口方法  进行多态
func (c *Connection) Start() {
	for {
		buf := make([]byte, 512)
		_, err := c.conn.Read(buf)
		if err != nil {
			return
		}
		c.block(c, buf)
	}
}

//
func (c *Connection) Stop() {
	if !c.isClosed {
		return
	}
	_ = c.conn.Close()
}

func (c *Connection) Send(buf []byte) error {
	_, err := c.conn.Write(buf)
	fmt.Println(string(buf))
	return err
}

func (c *Connection) GetConnId() uint32 {
	return c.connId
}

func (c *Connection) GetTcpConn() *net.TCPConn {
	return c.conn
}
func NewConnection(conn *net.TCPConn, cid uint32, block iface.IConnBlock) iface.IConnection {
	return &Connection{
		conn:     conn,
		connId:   cid,
		isClosed: false,
		block:    block,
	}
}
