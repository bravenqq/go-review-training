package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	con, err := net.Dial("unix", "/Users/abbynie/tmp/unix.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	reader := bufio.NewReader(con)
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			log.Println("read err:", err)
			return
		}
		fmt.Println(string(data))
	}
}
