// Package list provides ...
package list

import "testing"

var l = NewLinkList()

func TestInsert(t *testing.T) {
	tests := []struct {
		input Node
	}{
		{input: Node{Data: 10}},
		{input: Node{Data: 20}},
	}
	for _, test := range tests {
		l.Insert(test.input)
	}
	node := l.head
	for node != nil {
		if node.Data != tests[0].input.Data && node.Data != tests[1].input.Data {
			t.Error("l insert error")
		}
		node = node.next
	}
}
func TestRemove(t *testing.T) {
	tests := []struct {
		input Node
	}{
		{input: Node{Data: 10}},
	}
	for _, test := range tests {
		l.Remove(test.input)
	}
	node := l.head
	for node != nil {
		if node.Data == tests[0].input.Data {
			t.Error("l remove error")
		}
		node = node.next
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		out Node
	}{
		{out: Node{Data: 20}},
	}
	for _, test := range tests {
		n, err := l.Pop()
		if err != nil || n.Data != test.out.Data {
			t.Error("l pop error")
		}
	}
}
