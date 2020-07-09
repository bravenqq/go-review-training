// Package main provides ...
package main

import (
	"fmt"
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

	display("rV", reflect.ValueOf(os.Stderr))

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
