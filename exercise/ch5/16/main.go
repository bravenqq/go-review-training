/*
“编写多参数版本的strings.Join”
*/
package main

import (
	"strings"
)

func main() {

}

func Join(sep string, a ...string) string {
	return strings.Join(a, sep)
}
