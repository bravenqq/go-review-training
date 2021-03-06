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
	"time"
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

	who := conn.RemoteAddr().String()
	entering <- Client{name: who, c: ch}

	ch <- "You are " + who
	messages <- who + " has arrived"

	input := bufio.NewScanner(conn)
	timeout := 10 * time.Second
	timer := time.NewTimer(timeout)
	msgs := make(chan string)
	go func() {
		for input.Scan() {
			msgs <- input.Text()
		}
	}()

	for {
		select {
		case message := <-msgs:
			messages <- who + ": " + message
			timer.Reset(timeout)
		case <-timer.C:
			timer.Stop()
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
