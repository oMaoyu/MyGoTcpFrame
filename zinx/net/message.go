package net

import (
	"MyTcpFrame/zinx/iface"
)

type Message struct {
	//数据
	data []byte
	//长度
	len uint32
	//描述消息类型的字段，id
	msgid uint32
}

func NewMessage(data []byte, len uint32, msg uint32) iface.IMessage {
	return &Message{
		data:  data,
		len:   len,
		msgid: msg,
	}
}
// 获取对应字段数据
func (m *Message) GetData() []byte {
	return m.data
}
func (m *Message) GetLen() uint32 {
	return m.len
}
func (m *Message) GetId() uint32 {
	return m.msgid
}
// 更改对应字段数据
func (m *Message) SetData(data []byte) {
	m.data = data
}
func (m *Message) SetLen(len uint32) {
	m.len = len
}
func (m *Message) SetId(id uint32) {
	m.msgid = id
}
