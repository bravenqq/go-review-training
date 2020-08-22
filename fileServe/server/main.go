package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"sort"

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
	p := r.Path
	log.Println("path:", p)
	ph := root + p
	log.Println("path:", ph)
	fi, err := os.Stat(ph)
	if err != nil {
		return err
	}
	f, err := os.Open(ph)
	defer f.Close()
	if err != nil {
		return err
	}
	if fi.IsDir() {
		err = dirList(fs, f)
		return err
	}
	err = writeContent(fs, f)
	if err != nil {
		return err
	}
	return nil
}

func writeContent(fs pb.FileServe_ServeContentServer, f *os.File) error {
	r := bufio.NewReader(f)
	buf := make([]byte, 100)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		fs.Send(&pb.Chunk{Content: buf[:n]})
	}
	return nil
}

func dirList(fs pb.FileServe_ServeContentServer, f *os.File) error {
	dirs, err := f.Readdir(-1)
	if err != nil {
		return err
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })
	for _, d := range dirs {
		name := d.Name()
		if d.IsDir() {
			name += "/"
		}
		fs.Send(&pb.Chunk{Content: []byte(name)})
	}
	return nil
}

func main() {
	flag.StringVar(&root, "path", "./", "please input root path")
	flag.Parse()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("start port:%s err:%v\n", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterFileServeServer(s, &server{})
	s.Serve(lis)
}
