package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {

	var m runtime.MemStats
	for i := 0; i < 10; i++ {
		s := bigBytes()
		if s == nil {
			log.Println("oh noes")
		}
	}

	for i := 0; i < 20000; i++ {
		var c = new(int)
		*c = i
		fmt.Println(c)
	}
	runtime.ReadMemStats(&m)
	fmt.Printf("%v\n", bToMb(m.TotalAlloc))
	runtime.GC()

	time.Sleep(time.Second * 20)

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func bigBytes() *[]byte {
	s := make([]byte, 1000000)
	return &s
}
