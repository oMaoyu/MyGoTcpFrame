package net

import "MyTcpFrame/zinx/iface"

type Routers struct {
	// 存放多个接口
	routers map[uint32]iface.IRouter
	// 线程池数量
	workerSize int
	// 线程池
	taskQueue [] chan iface.IRequest
}

func NewRouters() *Routers {
	return &Routers{
		routers:    make(map[uint32]iface.IRouter),
		workerSize: MyConfig.Worker.Size,
		taskQueue:  make([]chan iface.IRequest,MyConfig.Worker.Size),
	}
}
func (msgRouters *Routers) AddRouter(key uint32, value iface.IRouter) {
	if _, ok := msgRouters.routers[key]; ok {
		return
	}
	msgRouters.routers[key] = value
}

// 执行对应方法
func (mr *Routers) PreOneRouterFunc(req iface.IRequest) {
	id := req.GetMsg().GetId()
	if router, ok := mr.routers[id]; ok {
		router.PreHandle(req)
		router.Handle(req)
		router.PostHandle(req)
	}
}

// 启动线程池
func (mr *Routers) StartWorkerPool() {
	for i := 0; i < mr.workerSize; i++ {
		mr.taskQueue[i] = make(chan iface.IRequest,MyConfig.Worker.TaskQueSize)
		go func(i int) {
			for {
				req := <- mr.taskQueue[i]
				mr.PreOneRouterFunc(req)
			}
		}(i)
	}
}
// 像任务队列发送任务
func (mr *Routers)SendReqToQueue(req iface.IRequest){
	id := req.GetConn().GetConnId()
	// 根据对应id  分配到对应任务池
	workerId := int(id) % mr.workerSize

	mr.taskQueue[workerId] <- req

}