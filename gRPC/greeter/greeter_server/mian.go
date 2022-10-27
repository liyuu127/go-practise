package main

import (
	context "context"
	"flag"
	"fmt"
	pb "github.com/liyuu127/go-practise/gRPC/greeter/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "the greetServer port")
)

var (
	userId   = "sdc-add-dad"
	userName = "liyu"
)

// greetServer is used to implement helloworld.GreeterServer.
type greetServer struct {
	pb.UnimplementedGreeterServer
}
type UserServer struct {
	pb.UnimplementedUserServer
}

func (s *greetServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", request.GetName())
	return &pb.HelloReply{Message: "hello " + request.GetName()}, nil
}

func (s *UserServer) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Received: %+v", request)
	if request.Username == nil {
		request.Username = &userName
	}

	if request.UserId == nil {
		request.UserId = &userId
	}

	rep := &pb.GetUserResponse{
		UserId:   request.GetUserId(),
		Username: request.GetUsername(),
	}
	return rep, nil
}

func main() {

	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greetServer{})
	pb.RegisterUserServer(s, &UserServer{})
	log.Printf("greetServer listening at %v", listener.Addr())
	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// s.GracefulStop()

}
