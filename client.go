package main


import (
	"fmt"
	"net"
)

func main(){
	tcpaddr,err := net.ResolveTCPAddr("tcp","127.0.0.1:7777")
	if err!= nil {
		panic(err.Error())
	}

	conn,err := net.DialTCP("tcp",nil,tcpaddr)
	if err != nil {
		panic(err.Error())
	}
	_, err1 := conn.Write([]byte("i am client"))
	if err1 != nil {
		panic(err1.Error())
	}
	rsp := make([]byte,100)
	n,err :=conn.Read(rsp)  //尽量使用conn原生read/write 否则可能造成 connect reset
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(rsp[:n]))
}
