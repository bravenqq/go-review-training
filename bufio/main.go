package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("./file", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("open file err:", err)
	}

	writer := bufio.NewWriterSize(f, 20)
	for i := 0; i < 10; i++ {
		_, err = writer.WriteString("hello,nieqianqian!\n")
		if err != nil {
			log.Fatal("write to file err:", err)
		}
		fmt.Println("buffered:", writer.Buffered())
		// writer.Flush()

	}
	f.Close()

}
