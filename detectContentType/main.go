package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	tp1 := detectContentType("./5.jpeg")
	fmt.Println("./test.jpeg type:", tp1)
	tp2 := detectContentType("./fe6dae00-c0e2-11ea-9a0b-a45e60f2fb75")
	fmt.Println("type:", tp2)
}

func detectContentType(file string) string {

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // defer file.Close() must after checked err.
	buf := make([]byte, 512)
	fmt.Println("data:", string(buf))
	_, err = f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return http.DetectContentType(buf)
}
