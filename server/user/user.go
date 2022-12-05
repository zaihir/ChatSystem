package user

import (
	"DemoServer/ChatSystem/server/object"
	"fmt"
	"net"
	"strings"
)


type UserInfo struct{
	Uid int
	MyConn net.Conn
	RevMsgChan chan string
	RevMsgBuff []byte
	SendMsgChan  chan string
	GotNewMsgTag  chan bool
	UserManage object.MyUserManger



}


func(this* UserInfo )GetUid() int {
	return this.Uid
}



func (this* UserInfo)Init(userManage object.MyUserManger){
	this.SendMsgChan = make(chan string,1000)
	this.RevMsgBuff = make ([]byte, 1024)
	this.GotNewMsgTag  = make( chan bool , 1)
	this.UserManage = userManage


	fmt.Printf("init this.UserManage  = %+v \n" , this.UserManage )

}
func (this *UserInfo) SetUid(uid int)int{
	this.Uid = uid
	return 0
}

func (this *UserInfo)SetConn(conn net.Conn)int{
	if conn == nil {

		return -1
		
	}

	this.MyConn = conn


	return 0
}



func (this* UserInfo) SendMsg(msg []byte)int{

	fmt.Println("receive msg " ,string(msg) )
	this.MyConn.Write([]byte(string(msg)  + "\n"))
	 //this.SendMsgChan <- string(msg)
	return 0
}



func (this* UserInfo)ProcessRecMsgFromConn()int {

	fmt.Printf("Enter ProcessRecMsgFromConn [%d]\n" , this.Uid , )
	fmt.Printf("before get  messgae if [%+v] \n" , this.RevMsgBuff )
	TmpMsg := string(this.RevMsgBuff)

	for  {



		num , err := this.MyConn.Read(this.RevMsgBuff )
		fmt.Printf("then  get  tranform to string [%s]\n" ,TmpMsg)

		fmt.Printf("num = [%d]  erro = %+v  , \n", num , err)
		if err != nil{
			fmt.Println(err)
			break
		}

		if num <= 0 {

			fmt.Print("rev msg len is %d /n " ,num)
			break
		}

		fmt.Printf("get orignal messgae if [%+v]" , this.RevMsgBuff )



		this.GotNewMsgTag <- true
		
	}



	return 0
}


func (this* UserInfo)	ProcessMsg()int{

	for  {
		fmt.Println("Prcess Msf" , this.Uid)
		select {

				case	tmpMsg ,ok := <- this.SendMsgChan :


					if ok {
						fmt.Println("receive msg " ,tmpMsg )
						this.MyConn.Write([]byte(tmpMsg + "\n"))
					}else{
						this.MyConn.Close()
					}
					break

				case  <- this.GotNewMsgTag :

					fmt.Printf("get GotNewMsgTag singer !%s!\n ",string(this.RevMsgBuff))

					SendMsgInfo := strings.Split(string(this.RevMsgBuff),"|")

					fmt.Printf("Get mes %+v  len = %d\n" , SendMsgInfo[0] ,len(SendMsgInfo) )

					if  len(SendMsgInfo[0]) != 0  && len(SendMsgInfo) == 2{

						fmt.Printf("this.UserManage  = %v  , this uid is %d \n", &this.UserManage , this.Uid)
						this.UserManage.BroadMassage([]byte(SendMsgInfo[1]))
						
					}
					break

			}


	}

	return 0
}


type User  struct {
	UserPool map[int] UserInfo
}


func (this *User) New () object.UserInfoObject{
	return new(UserInfo)
}







func init(){
	fmt.Println("enter init func")
}