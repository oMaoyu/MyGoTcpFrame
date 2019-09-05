package iface

type IServer interface {
	Start()
	Stop()
	Server()
	AddRouter(uint32,IRouter)
	GetConnMan()IConnManager
	// 用于注册
	RegisterStart(func(IConnection))
	RegisterStop(func(IConnection))
	// 提供执行调用
	ConnStartRun(IConnection)
	ConnStopRun(IConnection)

}