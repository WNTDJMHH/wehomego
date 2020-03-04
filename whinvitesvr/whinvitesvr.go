package main
import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "wehome/whinvitesvr/proto"
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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWhinviteSvrServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
