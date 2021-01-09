package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	d := net.Dialer{Timeout: time.Second * 0}
	con, err := d.Dial("tcp", "104.193.88.123:8080")
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
