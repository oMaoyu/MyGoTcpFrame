package iface

type IMessage interface {
	GetData() []byte
	GetLen() uint32
	GetId() uint32
	SetData([]byte)
	SetLen(uint32)
	SetId(uint32)
}
