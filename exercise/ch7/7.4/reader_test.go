package reader

import (
	"bytes"
	"testing"
)

func TestStringReader(t *testing.T) {
	s := "hello,world"
	r := NewReader(s)
	buffer := bytes.Buffer{}
	n, err := buffer.ReadFrom(r)
	if n != int64(len(s)) || err != nil {
		t.Fatalf("n=%d,err=%s", n, err)
	}
	if buffer.String() != s {
		t.Errorf("intput=%s,get=:%s", s, buffer.String())
	}
}
