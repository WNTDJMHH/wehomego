package main

import (
	"context"
	"log"
	"time"
	"fmt"
	"math/rand"
	"google.golang.org/grpc"
	pb "wehome/whinvitesvr/proto"
  "wehome/comm/global"
)

type WhinviteSvrClient struct{
}

func NewWhinviteSvrClient() * WhinviteSvrClient{
			
	return &WhinviteSvrClient{}
}

func GetConnForUin(uId uint64) (cc *grpc.ClientConn, err error ){
	endPoint := global.GetEndPointForClient("whinvitesvr", uId)
	if endPoint == nil{ 
			log.Fatalf("%v NotFoundEndPoint", uId)
			panic("NotFoundEndPoint")
	}	
	strAddress := fmt.Sprintf("%v:%v", endPoint.Ip, endPoint.Port)
	cc, err = grpc.Dial(strAddress, grpc.WithInsecure(), grpc.WithBlock())
	return 
} 

func (client *WhinviteSvrClient) Echo(req * pb.EchoReq) (rsp * pb.EchoRsp, err error){
	cc, err := GetConnForUin(rand.Uint64() + 1000000)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return 
	}
	defer cc.Close()
	c := pb.NewWhinviteSvrClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rsp, err = c.Echo(ctx, req)
	return 
}


/*
func main() {
	// Set up a connection to the server.
	reqEcho := pb.EchoReq{}
	reqEcho.Name = "jessehou"
	inviteCli := NewWhinviteSvrClient()
 	rspEcho, err := inviteCli.Echo(&reqEcho) 
	fmt.Println("Rsp", rspEcho, err)
}
*/
