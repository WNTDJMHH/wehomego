package client 

import (
	"context"
	"log"
	"time"
	"math/rand"
	pb "wehome/whstorekv/proto"
)

type WhStoreKvClient struct{
}

func NewWhStoreKvClient() * WhStoreKvClient{
			
	return &WhStoreKvClient{}
}

/*
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
*/

func (client *WhStoreKvClient) Get(strKey * string) (strValue * string, err error){
	cc, err := GetConnForUin(rand.Uint64() + 1000000)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return 
	}
	defer cc.Close()
	c := pb.NewWhStoreKvClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	objReq := pb.GetReq{}
	objRsp, err := c.Get(ctx, &objReq)
	*strValue = objRsp.Data.Value;
	return 
}


