package main
import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "wehome/whinvitesvr/proto"
	"os"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedWhinviteSvrServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoReq) (*pb.EchoRsp, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.EchoRsp{Message: "Hello " + in.GetName()}, nil
}

func main() {
	log.SetPrefix("whinvitesvr_");
	log.SetFlags(log.Ldate|log.Lshortfile)
	errFile,err:=os.OpenFile("/home/ubuntu/log/errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		panic("OpenlogFileFail")
	}
	log.SetOutput(errFile)	

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterWhinviteSvrServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
