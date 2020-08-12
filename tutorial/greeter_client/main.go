// Package main provides ...
package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "iohttps.com/nqq/go-review-training/tutorial/helloworld"
)

const (
	address     = "localhost:8080"
	defaultName = "World"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting %s\n", r.GetMessage())
}
