package net

import (
	"MyTcpFrame/zinx/iface"
	"bytes"
	"encoding/binary"
)

// 专门用来拆包封包
type DataPack struct {
}
// 初始化
func NewDp()*DataPack{
	return &DataPack{}
}
// 封
func (dp *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	// 封包首先要制定规则  这里我们的规则就是 给数据拼接为 长度,id,内容
	// 其中长度id为消息头  内容为消息体

	len := msg.GetLen()
	id := msg.GetId()
	data := msg.GetData()

	var buff bytes.Buffer
	// 写消息头
	//存入长度
	err := binary.Write(&buff, binary.LittleEndian, len)
	if err != nil {
		return nil, err
	}
	//存入id
	err = binary.Write(&buff, binary.LittleEndian, id)
	if err != nil {
		return nil, err
	}
	// 写消息体
	err = binary.Write(&buff, binary.LittleEndian, data)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), err
}

// 拆
// 这里只对消息头进行解析
func (dp *DataPack) Unpack(data []byte) (iface.IMessage, error) {
	//执行本方法只处理前八个字节,读取对应的消息头
	var msg Message
	red := bytes.NewReader(data)

	// 读取长度
	err := binary.Read(red, binary.LittleEndian, &msg.len)
	if err != nil {
		return nil, err
	}
	// 读取id
	err = binary.Read(red, binary.LittleEndian, &msg.msgid)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
