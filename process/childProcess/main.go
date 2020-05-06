// Package main provides ...
package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "3600")
	//fork and exec sleep,wait sleep to finsh
	err := cmd.Run()
	if err != nil {
		log.Println("err:", err)
	}
}
