package net

import (
	"MyTcpFrame/zinx/iface"
	"fmt"
	"net"
	"sync"
)

type Connection struct {
	conn     *net.TCPConn
	connId   uint32
	isClosed bool
	//router   iface.IRouter
	routers *Routers
	msgChan chan []byte
	server  iface.IServer
	//自定义连接设置属性
	properit map[string]interface{}
	pLock sync.RWMutex

}

// 实现接口方法  进行多态
func (c *Connection) Start() {
	go c.startRead()
	go c.startWrite()
	c.server.ConnStartRun(c)
}

// 对conn只进行读
func (c *Connection) startRead() {
	defer c.Stop()
	for {
		msg, err := GetMsg(c.conn)
		if err != nil {
			return
		}
		req := NewRequest(c, msg)
		if MyConfig.Worker.Size > 0 {
			c.routers.SendReqToQueue(req)
		} else {
			c.routers.PreOneRouterFunc(req)
		}
	}
}

// 对conn只进行写
func (c *Connection) startWrite() {
	for buff := range c.msgChan {
		_, _ = c.conn.Write(buff)
	}

}

// 关闭客户端
func (c *Connection) Stop() {

	if !c.isClosed {
		return
	}
	c.server.ConnStopRun(c)
	c.server.GetConnMan().Remove(int(c.GetConnId()))
	_ = c.conn.Close()
	close(c.msgChan)
	c.isClosed = false
}

// 往客户端写数据
func (c *Connection) Send(buf []byte, id uint32) (int, error) {
	fmt.Println(string(buf))
	dp := NewDp()
	buff, err := dp.Pack(NewMessage(buf, uint32(len(buf)), id))
	c.msgChan <- buff
	return 0, err
}

func (c *Connection) GetConnId() uint32 {
	return c.connId
}

func (c *Connection) GetTcpConn() *net.TCPConn {
	return c.conn
}
func (c *Connection) SetProperity(key string, value interface{}) {
	c.pLock.Lock()
	defer c.pLock.Unlock()
	c.properit[key] = value
}
func (c *Connection)GetProperity(key string) interface{} {
	c.pLock.RLock()
	defer c.pLock.RUnlock()
	return c.properit[key]
}
func (c *Connection)RemoveProperity(key string){
	c.pLock.Lock()
	defer c.pLock.Unlock()
	delete(c.properit, key)
}

func NewConnection(conn *net.TCPConn, cid uint32, block *Routers, server iface.IServer) iface.IConnection {
	return &Connection{
		conn:     conn,
		connId:   cid,
		isClosed: true,
		routers:  block,
		msgChan:  make(chan []byte),
		server:   server,
		properit:make(map[string]interface{}),
	}
}
