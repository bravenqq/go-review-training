/*
将一个文件的内容复制到另一个文件
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("请输入源文件的名称")
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("获取输入文件错误", err)
	}
	in = strings.Replace(in, "\n", "", -1)
	fmt.Println("in:", in)

	if !isFileExist(in) {
		log.Fatalf("cp file %s not exist\n", in)
	}
	inFile, err := os.Open(in)
	if err != nil {
		log.Fatalf("open in file:%s err:%v\n", in, err)
	}
	defer inFile.Close()

	fmt.Println("请输入目标文件名称")
	out, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("获取输出文件错误", err)
	}
	out = strings.Replace(out, "\n", "", -1)
	fmt.Println("out:", out)

	var outFile *os.File
	// if !isFileExist(out) {
	outFile, err = os.OpenFile(out, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	// } else {
	// 	outFile, err = os.Open(out)
	// 	if err != nil {
	// 		log.Fatalf("create out file:%s err:%v\n", out, err)
	// 	}
	// 	defer outFile.Close()
	// 	err = outFile.Truncate(0)
	// }
	if err != nil {
		log.Fatalf("open out file:%s err:%v\n", out, err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		log.Println("copy err:", err)
	}
}

func isFileExist(in string) bool {
	_, err := os.Stat(in)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
