package main

import (
    "fmt"
    "net"
)

func main(){
    tcpaddr,err := net.ResolveTCPAddr("tcp",":7777") //listen 7777
    if err!=nil {
        panic(err.Error())
    }
    listener,err := net.ListenTCP("tcp",tcpaddr)
    if err != nil {
        panic(err.Error())
    }

    for{
        conn,err := listener.Accept() //wait 
        if err != nil {
            continue
        }
        go dowork(conn)  //使用协程处理请求 高并发
    }
}

/**

处理业务	

**/

func dowork(conn net.Conn){
        defer conn.Close()
        r := make([]byte,100)
        n,err := conn.Read(r)
        if err != nil {
            panic("server error5")
        }
        fmt.Println(string(r[:n]))
        if err !=nil {
           
            panic("server error5")
        }
        conn.Write([]byte("i am server from xiecheng"))
//      conn.close()
}
