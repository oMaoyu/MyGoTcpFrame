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
func (cm *ConnManager) AddConn(v iface.IConnection) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	if _, ok := cm.conns[int(v.GetConnId())]; !ok {
		cm.conns[int(v.GetConnId())] = v
	}

}

//删除连接
func (cm *ConnManager) Remove(i int) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	delete(cm.conns, i)
}

//给定cid，返回连接句柄
func (cm *ConnManager) GetConn(k int) iface.IConnection {
	cm.lock.RLock()
	defer cm.lock.RUnlock()
	if conn, ok := cm.conns[k]; ok {
		return conn
	}
	return nil
}

//获取当前所有的连接的总数
func (cm *ConnManager) GetConnCount() int {
	return len(cm.conns)
}

//清除所有的连接
func (cm *ConnManager) ClearConn() {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	for id, con := range cm.conns{
		con.Stop()
		delete(cm.conns, id)
	}
}
