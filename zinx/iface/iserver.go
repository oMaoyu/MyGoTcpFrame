package iface

type IServer interface {
	Start()
	Stop()
	Server()
	AddRouter(IRouter)
}