// Package main provides ...
package main

import (
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	pb "iohttps.com/nqq/go-review-training/tutorial/helloworld"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(in *pb.HelloRequest, gs pb.Greeter_SayHelloServer) error {
	log.Println("Received:", in.GetName())
	name := in.Name
	for i := 0; i < 100; i++ {
		gs.Send(&pb.HelloReplay{Message: "Hello " + name + strconv.Itoa(i)})
	}
	return nil
}

func (s *server) SayHelloAgain(gs pb.Greeter_SayHelloAgainServer) error {
	var names []string
	for {
		in, err := gs.Recv()
		if err == io.EOF {
			gs.SendAndClose(&pb.HelloReplay{Message: "Hello " + strings.Join(names, ",")})
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}
		names = append(names, in.Name)
	}
}

func (s *server) SayHelloAgainAll(gs pb.Greeter_SayHelloAgainAllServer) error {
	for {
		in, err := gs.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Println("failed to recv:", err)
			return err
		}
		gs.Send(&pb.HelloReplay{Message: "Hello " + in.Name})
	}
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
