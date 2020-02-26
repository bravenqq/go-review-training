package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
		if strings.Contains(string(data), "EOF") {
			break
		}
		fmt.Println(string(data))
	}
}
