/*“strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。实现一个简单版本的NewReader，用它来构造一个接收字符串输入的HTML解析器”*/
package reader

import (
	"io"
)

//StringReader make a reader from string
type StringReader struct {
	s string
}

func NewReader(s string) *StringReader {
	return &StringReader{s}
}

func (sr *StringReader) Read(p []byte) (int, error) {
	n := copy(p, sr.s)
	sr.s = sr.s[n:]
	if len(sr.s) == 0 {
		return n, io.EOF
	}
	return n, nil
}
