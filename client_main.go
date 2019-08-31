package main

import (
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
	for {
		_,err = conn.Write(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		time.Sleep(1 * time.Second)

	}
}
