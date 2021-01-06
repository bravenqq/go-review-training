// Package main provides ...
package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("%p\n", &s)
	fmt.Printf("len:%d cap:%d\n", len(s), cap(s))
	fmt.Println(s[0])
}

func Join(slice []string, str string) string {

	var s, sep string
	for i := 0; i < len(slice); i++ {
		//1.sep + slice[i] 大小不确定，escape to heap
		//2.每次赋值之后之前sep+slice[i]的值不再使用，要被gc回收
		//所以这里会发生i-1次垃圾回收和堆内存分配
		s += sep + slice[i]
		sep = str
	}
	return s
}

func FastJoin(slice []string, str string) string {
	builder := &Builder{}
	if len(slice) == 0 {
		return ""
	}
	if len(slice) == 1 {
		return slice[1]
	}

	n := (len(slice) - 1) * len(str)
	for _, s := range slice {
		n += len(s)
	}

	builder.Grown(n)
	builder.Add(slice[0])
	for i := 1; i < len(slice); i++ {
		builder.Add(str)
		builder.Add(slice[i])
	}
	return builder.toString()
}

//Builder data builder
type Builder struct {
	index int
	data  []byte
}

//Grown grown data cap
func (builder *Builder) Grown(n int) {
	//make([]byte, len(builder.data), n + len(builder.data)) escapes to heap,不确定长度
	builder.data = make([]byte, len(builder.data), n+len(builder.data))
}

//Add add str to data
func (builder *Builder) Add(str string) (int, error) {
	builder.data = append(builder.data, str...)
	return len(str), nil
}

func (builder *Builder) toString() string {
	//string(builder.data) escape to heap
	return string(builder.data)
}
