package iface

type IRequest interface {
	GetConn() IConnection
	GetLen() uint32
	GetData() []byte
}
