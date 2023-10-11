package main

import (
	"context"
	"fmt"
	"grpclearn/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServer struct {
	proto.GreetServer
}

func (*HelloServer) SayHello(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	return &proto.Response{
		Message: "Hello " + req.Name,
	}, nil
}

func (*HelloServer) SayHalloClientStreming(stream proto.Greet_SayHalloClientStremingServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			fmt.Println("Ended ....... !")
			break
		} else {
			fmt.Println(req)
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("error : %s\n", err.Error())
	}

	grpcServer := grpc.NewServer()

	proto.RegisterGreetServer(grpcServer, &HelloServer{})

	fmt.Println("Connected")
	fmt.Println(grpcServer.Serve(lis))
}
