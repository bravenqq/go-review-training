package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
)

var pc [256]byte

func main() {
	v1 := flag.String("value1", "x", "请输入第一个消息值")
	v2 := flag.String("value2", "X", "请输入第二个消息值")
	flag.Parse()
	c1 := sha256.Sum256([]byte(*v1))
	c2 := sha256.Sum256([]byte(*v2))
	fmt.Printf("c1%x\n", c1)
	fmt.Printf("c2%x\n:", c2)
	pc = func() [256]byte {
		var p [256]byte
		for i := range p {
			p[i] = p[i/2] + byte(i&1)
		}
		return p
	}()
	fmt.Println(pc)

	fmt.Println(DiffCount(c1, c2))
}

func DiffCount(c1, c2 [32]byte) uint8 {

	var dif [32]byte
	for i := 0; i < len(c1); i++ {
		dif[i] = c1[i] ^ c2[i]
		fmt.Printf("c1:%d c2:%d dif:%d\n", c1[i], c2[i], dif[i])
	}

	fmt.Println(dif)
	var count uint8
	for _, v := range dif {
		count += PopCount(v)
	}
	return count
}

func PopCount(x byte) uint8 {
	return pc[x]
}
