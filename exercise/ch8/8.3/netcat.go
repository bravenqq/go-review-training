// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		tcpCon := conn.(*net.TCPConn)
		tcpCon.CloseWrite()
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	// mustCopy(conn, os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "EOF" {
			os.Stdin.Close()
			break
		}
		_, err = conn.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			log.Println(err)
		}
	}

	tcpCon := conn.(*net.TCPConn)
	tcpCon.CloseRead()
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}
