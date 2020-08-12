// Package main provides ...
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "iohttps.com/nqq/go-review-training/tutorial/helloworld"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReplay, error) {
	log.Println("Received:", in.GetName())
	return &pb.HelloReplay{Message: "Hello" + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReplay, error) {
	return &pb.HelloReplay{Message: "Hello again" + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("faild to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
