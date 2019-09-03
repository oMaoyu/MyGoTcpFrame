package net

import (
	"MyTcpFrame/zinx/iface"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	conn     *net.TCPConn
	connId   uint32
	isClosed bool
	router   iface.IRouter
}

// 实现接口方法  进行多态
func (c *Connection) Start() {
	for {
		// 拆包
		dp := NewDp()
		// 读取数据头
		head := make([]byte, dp.GetHead())
		_, err := io.ReadFull(c.conn, head)
		if err != nil {
			fmt.Println(err)
			return
		}
		msg, err := dp.Unpack(head)
		if err != nil {
			fmt.Println(err)
			return
		}
		if msg.GetLen() == 0 {
			fmt.Println("数据长度为:", msg.GetLen())
			return
		}
		data := make([]byte, msg.GetLen())
		_, err = io.ReadFull(c.conn, data)
		if err != nil {
			fmt.Println(err)
			return
		}
		msg.SetData(data)
		req := NewRequest(c, msg)
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
		router:   block,
	}
}
