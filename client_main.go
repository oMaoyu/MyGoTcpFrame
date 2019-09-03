package main

import (
	net2 "MyTcpFrame/zinx/net"
	"fmt"
	"net"
	"time"
)

func main(){

	conn,err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := []byte("hi oMaoyu")
	msg := net2.NewMessage(buf, uint32(len(buf)),0)

	dp := net2.NewDp()
	buf ,err = dp.Pack(msg)
	if err != nil {
		fmt.Println(buf)
		return
	}
	for {
		_,err = conn.Write(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		time.Sleep(1 * time.Second)

	}
}
