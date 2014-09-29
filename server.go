package main

import (
	"fmt"
	"net"
)

func main(){
	tcpaddr,err := net.ResolveTCPAddr("tcp4",":7777")
	if err!=nil {
		panic(err.Error())
	}
	listener,err := net.ListenTCP("tcp",tcpaddr)
	if err != nil {
		panic(err.Error())
	}

	for{
		conn,err := listener.Accept()
		if err != nil {
			continue
		}
		go dowork(conn)
	}
}

func dowork(conn net.Conn){
		defer conn.Close()
		r := make([]byte,100)
		n,err := conn.Read(r)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(string(r[:n]))
		if err !=nil {
			
			panic(err.Error())
		}
		conn.Write([]byte("i am server from xiecheng"))
//		conn.close()
}

