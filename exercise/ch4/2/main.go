package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	t := flag.Uint("type", 0, "please input code type")
	flag.Parse()
	fmt.Println(*t)
	v := flag.String("value", "x", "请输入需要编码的值")
	switch *t {
	case 0:
		fmt.Printf("%x\n", sha256.Sum256([]byte(*v)))
	case 1:
		fmt.Printf("%x\n", sha512.Sum384([]byte(*v)))
	case 2:
		fmt.Printf("%x\n", sha512.Sum512([]byte(*v)))

	}
}
