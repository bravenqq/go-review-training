// Package main provides ...
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w, wr io.Writer
	var buffer *bytes.Buffer
	wr = buffer
	//类型断言断言w的动态类型是否是*os.File类型，返回动态值
	f, ok := wr.(*os.File)
	fmt.Println(ok)
	fmt.Printf("type:%T value:%v\n", f, f)
	w = os.Stderr
	// 断言w是否是类型*os.File
	f, ok = w.(*os.File)
	fmt.Println(ok)
	fmt.Printf("type:%T value:%v\n", f, f)

	rw := wr.(io.ReadWriter)
	fmt.Println(rw)
	rwc, ok := wr.(io.ReadWriteCloser)
	fmt.Println(ok)
	fmt.Println(rwc)
}
