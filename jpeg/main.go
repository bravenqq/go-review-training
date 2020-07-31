// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	q := flag.Int("quality", 80, "pelease input quality")
	flag.Parse()
	file, err := os.Open("./dfmcd.jpeg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Create("./test.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	img, err := jpeg.Decode(file) //解码
	if err != nil {
		fmt.Println(err)
	}
	jpeg.Encode(file1, img, &jpeg.Options{Quality: *q}) //编码，但是将图像质量从100改成5
}
