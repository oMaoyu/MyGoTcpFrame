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
	router    iface.IRouter
}

// 实现接口方法  进行多态
func (c *Connection)  Start() {
	for {
		buf := make([]byte, 512)
		end, err := c.conn.Read(buf)
		if err != nil {
			return
		}
		req := NewRequest(c,buf[:end],uint32(end))
		c.router.Handle(req)
		c.router.PostHandle(req)
		c.router.PreHandle(req)
	}
}

// 关闭客户端
func (c *Connection) Stop() {
	if !c.isClosed {
		return
	}
	_ = c.conn.Close()
}
// 往客户端写数据
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
func NewConnection(conn *net.TCPConn, cid uint32, block iface.IRouter) iface.IConnection {
	return &Connection{
		conn:     conn,
		connId:   cid,
		isClosed: false,
		router:    block,

	}
}
