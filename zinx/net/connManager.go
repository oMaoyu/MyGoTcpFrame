package net

import (
	"MyTcpFrame/zinx/iface"
	"sync"
)

type ConnManager struct {
	//连接组
	conns map[int]iface.IConnection
	// 读写锁
	lock sync.RWMutex
}

func NewConnManager()*ConnManager{
	return &ConnManager{
		conns: make(map[int]iface.IConnection),
	}
}
//增加连接
func (cm *ConnManager) AddConn(int, iface.IConnection) {
}

//删除连接
func (cm *ConnManager) Remove(int) {

}

//给定cid，返回连接句柄
func (cm *ConnManager) GetConn(int) iface.IConnection {
	return nil
}

//获取当前所有的连接的总数
func (cm *ConnManager) GetConnCount() int {
	return 0
}

//清除所有的连接
func (cm *ConnManager) ClearConn() {

}
