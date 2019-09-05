package iface

type IConnManager interface {
	AddConn(int, IConnection) //增加连接
	Remove(int)               //删除连接
	GetConn(int) IConnection  //给定cid，返回连接句柄
	GetConnCount() int        //获取当前所有的连接的总数
	ClearConn()               //清除所有的连接
}
