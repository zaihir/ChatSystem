package object

import (
	"fmt"
	"net"
	"sync"
)

type UserInfoObject interface {
	SetUid(uid int)int
	SetConn(conn net.Conn)int
	SendMsg(msg []byte)int
	ProcessMsg()int
	Init(MyUserManger)
	GetUid() int
	ProcessRecMsgFromConn()int

}


type MyUserFactory interface{
	New() UserInfoObject
}


type MyUserManger struct {
	MyUserPool  map[int]UserInfoObject
	myMutex    sync.RWMutex

}

func (this* MyUserManger )Init(){
	this.MyUserPool = make(map[int]UserInfoObject)
}

func (this* MyUserManger )AddUser(uid int,userObject UserInfoObject) int{

	if  userObject == nil  || uid < 0 {
		return -1
	}
	
	_, ok := this.MyUserPool[uid]

	fmt.Print("find user [%d]  result = %b" , uid , ok)

	if ok  {
		return 0
	}

	this.MyUserPool[uid] = userObject
	return 0
}
func (this* MyUserManger ) GetUser(uid int) UserInfoObject{

	if uid < 0 {
		return nil
	}

	tmpUser := this.MyUserPool[uid]

	if nil == tmpUser {
		return nil
	}

	return tmpUser

}


func (this* MyUserManger) DeleteUser( uid int ) int{

	if uid< 0  {

		return -1
	}	


	tempUser := this.MyUserPool[uid]

	if tempUser != nil {


		delete(this.MyUserPool ,uid)
	}
	
	
	return 0
}


func (this *MyUserManger) BroadMassage (msg []byte) int{

fmt.Println("enter BroadMassage	")

    this.myMutex.Lock()

	for uid,userObject := range this.MyUserPool {


		fmt.Printf("send msg %s to user[%d] ,[%+v]\n  ",string(msg),uid,userObject )

		userObject.SendMsg(msg)

	}
	this.myMutex.Unlock()
	return  0

}