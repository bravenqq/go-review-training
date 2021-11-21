// Package main provides ...
package main

// golang spec
// TypeDecl = "type" (TypeSpec | "(" { TypeSpec ";" } ")" ) .
// TypeSpec = AliasDecl | TypeDef .

// |   alternation
// ()  grouping
// []  option (0 or 1 times)
// {}  repetition (0 to n times)
func main() {

}

func Join(s string, str string) { // slice []string is point
	r := s + str
	r = r + "he"

	for i := 0; i < 2; i++ {
		r = r + ","
	}
	// var s, sep string // s, sep is string type, is imultiable.
	// for i := 0; i < len(slice); i++ {
	// 	s += sep + slice[i] // s += sep + slice[i] // s = string1 + string2 // +, += operator // s multiable?
	// 	// malloc(sep + slice[i])
	// 	// string imultiable type heap
	//
	// 	// "abc" + "xyz" string literal
	// 	// 1. lfh := "abc" + "xyz"
	// 	// 2. lfh += lfh + "aa"   // lfh1 := lfh + "aa" (unsafe) (operator + , +=) lfh string
	// }
	// return s
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

// Builder data builder
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
