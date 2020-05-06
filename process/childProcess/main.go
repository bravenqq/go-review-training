// Package main provides ...
package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "3600")
	//fork and exec sleep,wait sleep to finsh
	// err := cmd.Run()

	//fork and exec sleep,but don't wait sleep to finsh
	err := cmd.Start()
	if err != nil {
		log.Println("err:", err)
	}
}
