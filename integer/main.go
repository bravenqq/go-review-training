// Package main provides ...
package main

import "fmt"

// START MIT
const (
	Uint8Min   uint8   = 0
	Uint8Max   uint8   = ^(Uint8Min)
	Uint16Min  uint16  = 0
	Uint16Max  uint16  = ^(Uint16Min)
	Uint32Min  uint32  = 0
	Uint32Max  uint32  = ^(Uint32Min)
	Uint64Min  uint64  = 0
	Uint64Max  uint64  = ^(Uint64Min)
	UintMin    uint    = 0
	UintMax    uint    = ^(UintMin)
	UintptrMin uintptr = 0
	UintptrMax uintptr = ^(UintptrMin)

	Int8Min int8 = int8(^Int8Max)
	//此处一定要使用无符号的0，若使用int8(0),^int8(0)因为有符号整数是补码形式保存的，
	//得到值11111111在计算机中为负数，仍然是int8类型，然后右移一位，负数移位符号位会保留，得到还是11111111，
	//是补码形式，对应的原码是10000001值为-1
	Int8Max  int8  = int8(^uint8(0) >> 1)
	Int16Min int16 = int16(^Int16Max)
	Int16Max int16 = int16(^uint16(0) >> 1)
	Int32Min int32 = int32(^Int32Max)
	Int32Max int32 = int32(^uint32(0) >> 1)
	Int64Min int64 = int64(^Int64Max)
	Int64Max int64 = int64(^uint64(0) >> 1)
	IntMin   int   = int(^IntMax)
	IntMax   int   = int(^uint(0) >> 1)
)

// END MIT

func main() {
	fmt.Printf("uint8 min value:%d, max value:%d\n", Uint8Min, Uint8Max)
	fmt.Printf("uint16 min value:%d, max value:%d\n", Uint16Min, Uint16Max)
	fmt.Printf("uint32 min value:%d, max value:%d\n", Uint32Min, Uint32Max)
	fmt.Printf("uint64 min value:%d, max value:%d\n", Uint64Min, Uint64Max)
	fmt.Printf("uint min value:%d, max value:%d\n", UintMin, UintMax)
	fmt.Printf("uintptr min value:%d, max value:%d\n", UintptrMin, UintptrMax)

	fmt.Printf("int8 min value:%d, max value:%d\n", Int8Min, Int8Max)
	fmt.Printf("int16 min value:%d, max value:%d\n", Int16Min, Int16Max)
	fmt.Printf("int32 min value:%d, max value:%d\n", Int32Min, Int32Max)
	fmt.Printf("int64 min value:%d, max value:%d\n", Int64Min, Int64Max)
	fmt.Printf("int min value:%d, max value:%d\n", IntMin, IntMax)

	arr := make([]int, 2147483648, 2147483648)
	fmt.Println(arr)
}
