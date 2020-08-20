package main

import (
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "iohttps.com/nqq/go-review-training/fileServe/file"
)

const (
	port string = ":8080"
)

var root string

type server struct {
	pb.UnimplementedFileServeServer
}

func (s *server) ServeContent(r *pb.ServeContentRequest, fs pb.FileServe_ServeContentServer) error {
	return errors.New("file not found")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("start port:%s err:%v\n", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterFileServeServer(s, &server{})
	s.Serve(lis)
}
