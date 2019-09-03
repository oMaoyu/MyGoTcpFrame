package iface

type IRequest interface {
	GetConn() IConnection
	GetMsg() IMessage
}
