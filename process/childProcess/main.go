// Package main provides ...
package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("find", "/Users/abbynie")
	//fork and exec sleep,wait sleep to finsh
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Println("err:", err, ":", stderr.String())
	}
	fmt.Printf("%+v\n", cmd.Process)

	c := exec.Command("find", "/Users/abbynie")
	c.Stderr = &stderr
	//fork and exec sleep,wait sleep to finsh
	err = c.Run()
	if err != nil {
		log.Println("err:", err, ":", stderr.String())
	}
	fmt.Printf("%+v\n", c.Process)

	//fork and exec sleep,but don't wait sleep to finsh
	// err := cmd.Start()
	// if err != nil {
	// 	log.Println("err:", err)
	// }
	// fmt.Printf("%+v", cmd.Process)
}
