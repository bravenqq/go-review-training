// Package main provides ...
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("find", "/Users/abbynie")
	//fork and exec sleep,wait sleep to finsh
	err := cmd.Start()
	if err != nil {
		log.Println("err:", err)
		fmt.Fprintln(cmd.Stderr)
	}
	fmt.Printf("%+v", cmd.Process)
	c := exec.Command("find", "/Users/abbynie")

	//fork and exec sleep,wait sleep to finsh
	err = c.Start()
	if err != nil {
		log.Println("err:", err)
		fmt.Fprintln(c.Stderr)
	}
	fmt.Printf("%+v", c.Process)

	//fork and exec sleep,but don't wait sleep to finsh
	// err := cmd.Start()
	// if err != nil {
	// 	log.Println("err:", err)
	// }
	// fmt.Printf("%+v", cmd.Process)
}
