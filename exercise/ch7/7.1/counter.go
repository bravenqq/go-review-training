package counter

import (
	"bufio"
	"bytes"
)

type LineCounter struct {
	lines int
}

func (l *LineCounter) Write(p []byte) (int, error) {
	for _, b := range p {
		if b == '\n' {
			l.lines++
		}
	}
	return len(p), nil
}

type WordCounter struct {
	words int
}

func (w *WordCounter) Write(p []byte) (int, error) {
	scaner := bufio.NewScanner(bytes.NewReader(p))
	scaner.Split(bufio.ScanWords)
	for scaner.Scan() {
		w.words++
	}
	return len(p), nil
}
