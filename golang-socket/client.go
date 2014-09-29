
ckage main


import (
    "fmt"
    "net"
)

func main(){
    tcpaddr,err := net.ResolveTCPAddr("tcp4","127.0.0.1:7777") //目标服务器

    if err!= nil {
        panic("client error1")
    }  

    conn,err := net.DialTCP("tcp",nil,tcpaddr)
    if err != nil {
        panic(err.Error())
    }  
    _, err1 := conn.Write([]byte("i am client11111111"))
    if err1 != nil {
        panic("client error3")
    }  
    rsp := make([]byte,100)
    n,err :=conn.Read(rsp)
    if err != nil {
        panic(err.Error())
    }  
    fmt.Println(string(rsp[:n]))
}


