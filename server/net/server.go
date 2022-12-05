package net

import (

	"DemoServer/ChatSystem/server/object"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type TcpServer struct {
	Ip string
	Port int
	UserFactory object.MyUserFactory
	UserManger object.MyUserManger

}

func (this *TcpServer)Handle( myConn net.Conn){
	fmt.Println("connect succ")

}

func (this* TcpServer)Init(UserObject object.MyUserFactory)int  {
	fmt.Println("  enter Init ")
	if nil == UserObject {
		fmt.Println("UserObject is nill ,out  ")
		return -1
	}
	this.UserFactory =  UserObject


	this.UserManger.Init()
	return 0

}

func (this *TcpServer)BoardLogin(){

}

func (this *TcpServer) Start (){
	fmt.Println("Start")


	this.Ip = "127.0.0.1"
	this.Port = 8888
	//fmt.Println("%+v " ,)
	adress := fmt.Sprintf("%s:%d" , this.Ip,this.Port )

	fmt.Println("listen ,%s" ,adress )

	fmt.Println("11111111" )

	MyListenr,err := net.Listen("tcp",adress )

	//defer MyListenr.Close()

	fmt.Println("MyListenr [%+v] . [%+v]",MyListenr ,err)

	if err != nil{
		fmt.Println(err)
	}else{
		for {
			conn ,err := MyListenr.Accept()

			if err != nil{

				continue
			}else{

				TmpUser := this.UserFactory.New()
				fmt.Println("connet user is [%s] ",conn.RemoteAddr().String())

				UserIdStr := strings.Split(conn.RemoteAddr().String(),":")[1]




				UserId,_ := strconv.Atoi( UserIdStr)
				fmt.Println("uid =" ,UserId)
				fmt.Printf("this.UserManger adress = %+v ",&this.UserManger)
				TmpUser.Init(this.UserManger)


				TmpUser.SetUid(UserId)
				TmpUser.SetConn(conn)
				this.UserManger.AddUser(UserId,TmpUser)


				go TmpUser.ProcessRecMsgFromConn()

				go TmpUser.ProcessMsg()


				LoginMsg :=  fmt.Sprintf("user %d login  " ,UserId )

				fmt.Println(LoginMsg)

				this.UserManger.BroadMassage([]byte(LoginMsg))
			}
		}
	}
}

