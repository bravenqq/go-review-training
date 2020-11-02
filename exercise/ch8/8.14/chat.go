// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan Client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

type Client struct {
	name   string
	status bool
	c      client
}

func broadcaster() {
	clients := make(map[client]Client) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli, c := range clients {
				if c.status {
					cli <- msg
				}
			}

		case cli := <-entering:
			for _, client := range clients {
				if client.status {
					cli.c <- client.name + " has arrived"
				} else {
					cli.c <- client.name + " has left"
				}
			}
			cli.status = true
			clients[cli.c] = cli

		case cli := <-leaving:
			// delete(clients, cli)
			c := clients[cli]
			c.status = false
			clients[cli] = c
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	ch <- "please input your name which formate is name:your name"

	// messages <- who + " has arrived"

	input := bufio.NewScanner(conn)
	// timeout := 10 * time.Second
	// timer := time.NewTimer(timeout)
	msgs := make(chan string)
	var who string
	go func() {
		for input.Scan() {
			//判断输入是否包含name
			if strings.Contains(input.Text(), "name") {
				//解析出名字
				n := strings.Split(input.Text(), ":")
				if len(n) <= 1 {
					ch <- "please input your name use corrent formate"
					continue
				}
				who = n[1]
				//提示you are + name
				ch <- "You are " + who
				//entering
				entering <- Client{name: who, c: ch}
				//messages <- who
				msgs <- "has arrived"
			} else {
				msgs <- input.Text()
			}
		}
	}()

	for {
		select {
		case message := <-msgs:
			messages <- who + ": " + message
			// timer.Reset(timeout)
			// case <-timer.C:
			// 	timer.Stop()
			leaving <- ch
			messages <- who + " has time out"
			conn.Close()
			break
		}
	}

	// NOTE: ignoring potential errors from input.Err()

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
