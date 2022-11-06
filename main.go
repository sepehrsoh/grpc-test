package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pbMessages "grpc_test/proto/build/go/messages"
	pbService "grpc_test/proto/build/go/services"
	"log"
	"net"
	"net/http"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	grpcServer := grpc.NewServer()
	log.Println("Start HTTP Server on :", listener.Addr())
	pbService.RegisterHelloServer(grpcServer, &Server{})
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	ctx := context.Background()
	if err = pbService.RegisterHelloHandlerFromEndpoint(ctx, mux, ":9000", opts); err != nil {
		panic(err)
	}
	go grpcServer.Serve(listener)
	srv := http.Server{Addr: ":8000", Handler: mux}

	if err = srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, req *pbMessages.HelloRequest) (*pbMessages.HelloReply, error) {
	log.Println("message to say hello")
	name := req.GetName()
	return &pbMessages.HelloReply{Message: fmt.Sprintf("hello %v", name)}, nil
}
