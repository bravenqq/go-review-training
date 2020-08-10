package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) < 2 {
		log.Fatal("请输入地区和主机")
	}
	for i := 1; i < len(os.Args); i++ {
		strs := strings.Split(os.Args[i], "=")
		c := Clock{name: strs[0], host: strs[1]}
		con, err := net.Dial("tcp", c.host)
		if err != nil {
			log.Fatalf("connect:%s err:%v", c.host, err)
			continue
		}
		wg.Add(1)
		defer con.Close()
		go c.Watch(os.Stdout, con)
	}
	wg.Wait()
}

type Clock struct {
	name string
	host string
}

func (c *Clock) Watch(w io.Writer, r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Fprintf(w, "%s:%s\n", c.name, scanner.Text())
	}
}
