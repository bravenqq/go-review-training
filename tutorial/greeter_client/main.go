// Package main provides ...
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

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

	sm, err := c.SayHelloAgain(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		sm.Send(&pb.HelloRequest{Name: name + strconv.Itoa(i)})
	}
	reply, err := sm.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Greeting:", reply.Message)
	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	//
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// for {
	// 	reply, err := r.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Println("faild to rec:", err)
	// 	}
	// 	fmt.Println("Greeting:", reply.Message)
	// }
}
