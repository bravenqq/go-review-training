package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./article")
	if err != nil {
		log.Fatal("open file err:", err)
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 10)
	// fmt.Println("buffered count:", r.Buffered())
	// bt, err := r.ReadByte()
	// if err != nil {
	// 	log.Fatal("read byte err:", err)
	// }
	// fmt.Println("buffered count:", r.Buffered())
	// p := make([]byte, 50)
	// n, err := r.Read(p)
	// if err != nil {
	// 	log.Fatal("read err:", err)
	// }
	// fmt.Printf("read to p count:%d, data:%v\n", n, string(p))
	//
	// bt, err = r.ReadByte()
	// if err != nil {
	// 	log.Fatal("read byte err:", err)
	// }
	// fmt.Println("buffered count:", r.Buffered())
	// fmt.Printf("reader byte data:%c\n", bt)

	// _, err = r.Discard(21)
	// if err != nil {
	// 	log.Fatal("discard err:", err)
	// }
	// p := make([]byte, 50)
	// n, err := r.Read(p)
	// if err != nil {
	// 	log.Fatal("read err:", err)
	// }
	// fmt.Printf("read to p count:%d, data:%v\n", n, string(p))

	// data, err := r.ReadBytes('\n')
	// if err != nil {
	// 	log.Fatal("read bytes err:,", err)
	// }
	// fmt.Printf("read bytes data:%s", string(data))
	// data, _, _ = r.ReadLine()
	// fmt.Printf("read line data:%s\n", string(data))
	data, err := r.ReadBytes('\n')
	if err != nil {
		log.Fatal("read slice err:", err)
	}
	fmt.Printf("read slice data:%s", string(data))
	// err = r.UnreadByte()
	// if err != nil {
	// 	log.Fatal("unread rune err:", err)
	// }
	//
	// rn, _, err := r.ReadRune()
	// if err != nil {
	// 	log.Fatal("read rune err:", err)
	// }
	// fmt.Printf("read rune data:%c\n", rn)

	file, err := os.OpenFile("./file", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal("open file err:", err)
	}
	defer file.Close()
	n, err := r.WriteTo(file)
	if err != nil {
		log.Fatal("write to file err:", err)
	}
	fmt.Println("write count:", n)

}
