// Package main provides ...
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

var values map[reflect.Type]reflect.Value

func main() {
	// var a int64
	// a = 3
	// t := reflect.TypeOf(a)
	// fmt.Println(t)
	// v := reflect.ValueOf(a)
	// fmt.Println(v)
	// a = v.Interface().(int64)
	// ch := make(chan int)
	// f := func() {
	// 	fmt.Println("test")
	// }
	// slice := []int{1, 2, 3}
	// m := make(map[string]string)
	// p := struct {
	// 	name string
	// 	age  int
	// }{"nqq", 26}
	// var w io.Writer
	// w = os.Stdout
	// tests := []interface{}{-10, 10, 3.12, false, "hello world", ch, f, slice, m, 6.18, p, w}
	// for _, v := range tests {
	// 	fmt.Println("src:", v, " dest:", Any(v))
	// }

	// display("rV", reflect.ValueOf(os.Stderr))

	// strangelove := Movie{
	// 	Title:    "Dr. Strangelove",
	// 	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	// 	Year:     1964,
	// 	Color:    false,
	// 	Actor: map[string]string{
	// 		"Dr. Strangelove":            "Peter Sellers",
	// 		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
	// 		"Pres. Merkin Muffley":       "Peter Sellers",
	// 		"Gen. Buck Turgidson":        "George C. Scott",
	// 		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
	// 		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	// 	},
	// 	Oscars: []string{
	// 		"Best Actor (Nomin.)",
	// 		"Best Adapted Screenplay (Nomin.)",
	// 		"Best Director (Nomin.)",
	// 		"Best Picture (Nomin.)",
	// 	},
	// }

	// display("strangelove", reflect.ValueOf(strangelove))

	//reflect.Value修改值
	// x := 2
	// a := reflect.ValueOf(2)
	// fmt.Println("a:", a)
	// b := reflect.ValueOf(x)
	// fmt.Println("b:", b)
	// c := reflect.ValueOf(&x)
	// fmt.Println("c:", c)
	// d := c.Elem()
	// fmt.Println("d:", d)
	// fmt.Println(a.CanAddr())
	// fmt.Println(b.CanAddr())
	// fmt.Println(c.CanAddr())
	// fmt.Println(d.CanAddr())
	//通过reflect.Value 来修改变量的值
	// px := d.Addr().Interface().(*int)
	// *px = 3
	// px := c.Interface().(*int)
	// *px = 5
	// fmt.Println("x:", x)
	// d.Set(reflect.ValueOf(4))
	// fmt.Println("x:", x)
	var w1 io.Writer = os.Stdout
	var w2 io.Writer = &bufio.Writer{}

	values = make(map[reflect.Type]reflect.Value)
	values[InterfaceOf((*io.Writer)(nil))] = reflect.ValueOf(w2)
	values[InterfaceOf((*io.Writer)(nil))] = reflect.ValueOf(w1)
	for k, v := range values {
		fmt.Println("k:", k, " value:", v)
	}
	out := Get(InterfaceOf((*io.Writer)(nil))).Interface().(*os.File)
	out.WriteString("hello world\n")
	display("", reflect.ValueOf(w1))
}

func InterfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("Called inject.InterfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
	}

	return t
}

func Get(t reflect.Type) reflect.Value {
	val := values[t]

	if val.IsValid() {
		return val
	}

	// no concrete types found, try to find implementors
	// if t is an interface
	if t.Kind() == reflect.Interface {
		for k, v := range values {
			if k.Implements(t) {
				val = v
				break
			}
		}
	}

	// // Still no type found, try to look it up on the parent
	// if !val.IsValid() && i.parent != nil {
	// 	val = i.parent.Get(t)
	// }

	return val

}

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + "Ox" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + "value"
	}
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s=invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			filedPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(filedPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s=nil\n", path)
		} else {
			display(fmt.Sprintf("*%s", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s=nil\n", path)
		} else {
			fmt.Printf("%s.type=%s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
