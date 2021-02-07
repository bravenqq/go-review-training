package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"unsafe"
)

var memprofile = flag.String("memprofile", "mem.prof", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
	//最大成员类型FloatValue占4字节，按4字节对齐
	//FloatValue起始地址要按4字节对齐所以example占用12字节内存
	example := Example{false, 20.33, 20}
	alignBoundary := unsafe.Alignof(example)
	fmt.Printf("example alignBoundary:%d size:%d\n ", alignBoundary, unsafe.Sizeof(example))
	fmt.Printf("BoolValue size:%d offset:%d\n", unsafe.Sizeof(example.BoolValue), unsafe.Offsetof(example.BoolValue))
	fmt.Printf("IntValue size:%d offset:%d\n", unsafe.Sizeof(example.IntValue), unsafe.Offsetof(example.IntValue))
	fmt.Printf("FloatValue size:%d offset:%d\n", unsafe.Sizeof(example.FloatValue), unsafe.Offsetof(example.FloatValue))
	// fmt.Printf("StringValue size:%d offset:%d\n", unsafe.Sizeof(example.StringValue), unsafe.Offsetof(example.StringValue))
	var res bool
	fmt.Printf("res alignBoundary:%d size:%d", unsafe.Alignof(res), unsafe.Sizeof(res))

	//按机器字对齐，占用6*8=48字节
	var company Company
	fmt.Printf("company alignBoundary:%d size:%d\n", unsafe.Alignof(company), unsafe.Sizeof(company))
}

//Example example
type Example struct {
	BoolValue  bool
	FloatValue float32
	IntValue   int16
}

//Company company
type Company struct {
	BoolValue bool
	Name      string
	Employee  []int
}
