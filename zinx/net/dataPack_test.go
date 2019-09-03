package net

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestDataDemo(t *testing.T){
	fmt.Println("执行测试文件")

	// 测试数据
	// 服务器
	go func() {
		listen,err :=net.Listen("tcp","0.0.0.0:8888")
		if err != nil {
			t.Error(err)
			return
		}
		defer listen.Close()
		conn,err :=listen.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		for {
			// 设置读取数据  读取消息头
			head := make([]byte,8)
			// io.ReadFull 读取数据,可以读取指定Bety长度数据,未读取完不返回
			_,err = io.ReadFull(conn,head)
			if err != nil {
				t.Error(err)
				return
			}
			dp := NewDp()
			msg,err := dp.Unpack(head)
			if err != nil {
				t.Error(err)
				return
			}
 			fmt.Println("消息长度:",msg.GetLen())
			if msg.GetLen() == 0 {
				return
			}
			// 获取消息体
			data := make([]byte,msg.GetLen())
			_,err = io.ReadFull(conn,data)
			msg.SetData(data)
			fmt.Println("消息体:",string(data))
		}
	}()
	//客户端
	go func() {
		data1 := []byte("HI oMaoyu 你好么")
		data2 := []byte("我很好 不要担心我")
		data3 := []byte("by 卖毛玉的小贩")

		msg := NewMessage(data1,uint32(len(data1)),0)
		msg1 := NewMessage(data2,uint32(len(data2)),0)
		msg2 := NewMessage(data3,uint32(len(data3)),0)

		dp := NewDp()
		info1,_ := dp.Pack(msg)
		info2,_ := dp.Pack(msg1)
		info3,_ := dp.Pack(msg2)

		// 发送数据
		conn,err := net.Dial("tcp","127.0.0.1:8888")
		if err != nil {
			t.Error(err)
			return
		}

		 conn.Write(append(append(info1, info2...), info3...))
	}()

	time.Sleep(3 *time.Second)
}