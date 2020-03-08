// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}

//!+
func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	lines := make(chan string)
	go func(lines chan<- string) {
		for input.Scan() {
			lines <- input.Text()
		}

		if input.Err() != nil {
			log.Println(input.Err())
		}
	}(lines)

	go func() {
		wg.Wait()
	}()

	timeout := 10 * time.Second
	timer := time.NewTimer(timeout)
	for {
		select {
		case <-timer.C:
			fmt.Println("time out")
			timer.Stop()
			return
		case text := <-lines:
			timer.Reset(timeout)
			wg.Add(1)
			go echo(c, text, 1*time.Second)
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
