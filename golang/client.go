package main

import (
    "context"
	"log"

    "greet/pb"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
    if err != nil {
        log.Printf("grpc.Dial: %v\n", err)
        return
    }
    defer conn.Close()
    client := hello.NewGreeterClient(conn)
    req := &hello.HelloRequest{
    	Name: "Yuta Nakamura",
	}
    message, err := client.Hello(context.Background(), req)
    if err != nil {
        log.Printf("%v\n", err)
        return
    }
    log.Printf( "%s\n", message)
}