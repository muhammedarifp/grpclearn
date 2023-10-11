package main

import (
	"context"
	"fmt"
	"grpclearn/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	client := proto.NewGreetClient(clientConn)
	res, nwerr := client.SayHello(context.Background(), &proto.Request{
		Name: "Arif",
	})

	rs := []proto.Request{
		proto.Request{
			Name: "One",
		},
		proto.Request{
			Name: "Two",
		},
		proto.Request{
			Name: "One",
		},
		proto.Request{
			Name: "Two",
		},
		proto.Request{
			Name: "One",
		},
		proto.Request{
			Name: "Two",
		},
		proto.Request{
			Name: "One",
		},
		{
			Name: "Two",
		},
	}

	stream, sterr := client.SayHalloClientStreming(context.Background())
	if sterr != nil {
		log.Fatal("Stream Error found")
	}
	for _, val := range rs {
		time.Sleep(time.Second * 1)
		stream.Send(&val)
	}

	if nwerr != nil {
		fmt.Println("new error is : ", nwerr)
	}
	fmt.Println(res)
}
