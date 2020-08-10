package counter

import (
	"testing"
)

func TestLineCounter(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"hfghgjfh\njgsjgjkgfjk\njsgdjkgfkjgf\n", 3},
	}
	l := &LineCounter{}

	for _, test := range tests {
		l.Write([]byte(test.input))
		if l.lines != test.output {
			t.Errorf("input:%s want:%d get:%d", test.input, test.output, l.lines)
		}
	}
}

func TestWordCounter(t *testing.T) {
	c := &WordCounter{}
	data := [][]byte{
		[]byte("The upcoming word is sp"),
		[]byte("lit across the buffer boundary. "),
		[]byte(" And this one ends on the buffer boundary."),
		[]byte(" Last words."),
	}
	for _, p := range data {
		n, err := c.Write(p)
		if n != len(p) || err != nil {
			t.Logf(`bad write: p="%s" n=%d err="%s"`, string(p), n, err)
			t.Fail()
		}
	}
	if c.words != 20 {
		t.Logf("words: %d != 19", c.words)
		t.Fail()
	}
}
