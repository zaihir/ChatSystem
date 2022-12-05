package main

import (
	"DemoServer/ChatSystem/server/object"
	"DemoServer/ChatSystem/server/user"
	"DemoServer/ChatSystem/server/net"
	"fmt"
)
type  PeopleInfo struct {
	ago int
	name    string

}

func main(){
	fmt.Println("test ok")




	MyUser := new(user.User)
	UserManger := new(object.MyUserManger)

	myServer := net.TcpServer {"127.0.0.1" ,888 ,MyUser ,*UserManger}
	myServer.Init(MyUser)
	myServer.Start()
}



