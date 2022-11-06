package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pbMessages "grpc_test/proto/build/go/messages"
	pbServer "grpc_test/proto/build/go/services"
)

func main() {
	client, err := grpc.Dial(
		"127.0.0.1:9000",
		grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	helloClient := pbServer.NewHelloClient(client)
	response, err := helloClient.SayHello(context.Background(), &pbMessages.HelloRequest{Name: "sepehr"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("response : %v", response.String())
}
