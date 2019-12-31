package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("./file", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("open file err:", err)
	}
	writer := bufio.NewWriter(f)
	_, err = writer.WriteString("聂倩倩,你好\n")
	if err != nil {
		log.Fatal("write to file err:", err)
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal("flush to file err:", err)
	}
	f.Close()

	f, err = os.OpenFile("./file", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("open file err:", err)
	}

	reader := bufio.NewReader(f)
	for {
		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		os.Stdout.Write(data)
	}

}
