// Package main provides ...
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Second * 0}
	URL := "http://www.baidu.com:8080/"
	resp, err := client.Get(URL)
	if err != nil {
		log.Fatalf("Get url:%s err:%v\n", URL, err)
	}
	fmt.Printf("resp:%+v\n", resp)

}
