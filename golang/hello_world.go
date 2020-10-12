package main

import (
	"context"
	"log"
	"net"

	"greet/pb" // ローカルをパッケージ化しているので、必要に応じてgo mod init greet を実行してください。
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	server := grpc.NewServer()
	hello.RegisterGreeterServer(server, &Service{})
	log.Printf("Greeter service is running!")
	server.Serve(listener)
}

type Service struct{}

func (*Service) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	log.Printf("Helloが呼び出されました！")

	return &hello.HelloResponse{
		Message: "Hello, " + req.Name,
	}, nil
}
