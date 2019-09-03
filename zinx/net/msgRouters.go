package net

import "MyTcpFrame/zinx/iface"

type Routers struct {
	// 存放多个接口
	routers map[uint32]iface.IRouter
}

func NewRouters() *Routers {
	return &Routers{routers: make(map[uint32]iface.IRouter)}
}
func (msgRouters *Routers)AddRouter(key uint32,value iface.IRouter){
	if _,ok := msgRouters.routers[key];ok{
		return
	}
	msgRouters.routers[key] = value
}
// 执行对应方法
func (mr *Routers)PreOneRouterFunc(req iface.IRequest){
	 id := req.GetMsg().GetId()
	 if router,ok := mr.routers[id];ok {
	 	router.PreHandle(req)
	 	router.Handle(req)
	 	router.PostHandle(req)
	 }
}
