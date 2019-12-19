package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println("Println:")
	fmt.Println(sample)
	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println()
	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)
	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)
	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	data := []byte(sample)
	fmt.Println("Byte loop:")
	for i := 0; i < len(data); i++ {
		fmt.Printf("%q ", data[i])
	}
	fmt.Println()

	fmt.Println("rune loop:")
	for index, runeValue := range sample {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
