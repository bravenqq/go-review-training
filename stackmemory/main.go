// Package main provides ...
package main

func main() {
	strs := []string{"hell", "wor", "test", "nqq"}
	Join(strs, ",")
}

func Join(slice []string, str string) string {

	var s, sep string
	for i := 0; i < len(slice); i++ {
		f(sep + slice[i]) //此处sep+slice[i] does not escape
		//此处sep+slice escape逃逸了，
		//1.s+=sep+slice[i]=>tmp = sep+slice[i],s = tmp
		//2. 一次循环结束后循环内变量sep+slice[i]的内存要回收，但是sep+slice[i]被外部变量s引用了，所以将sep+slice[i]分配到堆中

		s += sep + slice[i]
		sep = str
	}
	return s
}
func f(s string) {
}
