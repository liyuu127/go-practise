package main

import (
	"context"
	"flag"
	pb "github.com/liyuu127/go-practise/gRPC/greeter/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	dial, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer dial.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()

	// client := pb.NewGreeterClient(dial)
	// reply, err := client.SayHello(ctx, &pb.HelloRequest{
	// 	Name: *name,
	// })

	userClient := pb.NewUserClient(dial)
	id := "123"
	reply, err := userClient.GetUser(ctx, &pb.GetUserRequest{
		UserId: &id,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %+#v", reply)

}
