package reader

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "hello,world"
	buffer := bytes.Buffer{}
	n, err := buffer.ReadFrom(LimitReader(strings.NewReader(s), 5))
	if n != 5 || err != nil {
		t.Fatalf("n=%d,err=%s", n, err)
	}
	if buffer.String() != s[:5] {
		t.Errorf("input=%s,get=%s,want=%s", s, buffer.String(), s[:5])
	}
}
