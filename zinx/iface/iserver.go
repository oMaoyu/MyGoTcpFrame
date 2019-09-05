package iface

type IServer interface {
	Start()
	Stop()
	Server()
	AddRouter(uint32,IRouter)
	GetConnMan()IConnManager
}