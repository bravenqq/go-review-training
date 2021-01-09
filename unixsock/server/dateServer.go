package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	l, err := net.Listen("unix", "/Users/abbynie/tmp/unix.sock")
	if err != nil {
		log.Fatalln("listen err:", err)

	}
	time.Sleep(time.Minute)
	for {
		con, err := l.Accept()
		if err != nil {
			log.Fatalln("connect err:", con)
		}
		defer con.Close()
		// _, err = con.Write([]byte("hi,nqq\n"))

		read := bufio.NewReader(os.Stdin)
		go func() {
			for {
				line, _, err := read.ReadLine()
				if err != nil {
					log.Println("read err:", err)
				}
				if string(line) == "EOF" {
					break
				}
				_, err = con.Write(line)
				if err != nil {
					log.Fatalln("write err:", err)
				}
				con.Write([]byte{'\n'})
			}
			con.Close()
		}()
	}
}
