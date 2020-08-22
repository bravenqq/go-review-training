package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "iohttps.com/nqq/go-review-training/fileServe/file"
)

const (
	address = "localhost:8080"
)

func main() {
	file := flag.String("file", "main.go", "please input file")
	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewFileServeClient(conn)
	fc, err := c.ServeContent(context.Background(), &pb.ServeContentRequest{Path: *file})
	if err != nil {
		log.Println(err)
		return
	}
	for {
		data, err := fc.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Println("recv err:", err)
			return
		}
		fmt.Println(string(data.GetContent()))
	}
}
