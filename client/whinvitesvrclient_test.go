package main
import (
	"testing"
	"fmt"
	pb "wehome/whinvitesvr/proto"
)

func TestEcho(t * testing.T){
	reqEcho := pb.EchoReq{}
	reqEcho.Name = "jessehou"
	inviteCli := NewWhinviteSvrClient()
 	rspEcho, err := inviteCli.Echo(&reqEcho) 
	fmt.Println("Rsp", rspEcho, err)
	if rspEcho.Message != "Hello jessehou"{
		t.Error("RspMessageErr")
	}else{
		t.Log("TestOk")
	}
}
