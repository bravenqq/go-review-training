// Package main provides ...
package main

import (
	"fmt"
	"sync"
)

func main() {
	s := NewStack()

	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	for s.Pop() != nil {
		fmt.Println(s.Pop().(int))
	}
}

type Stack struct {
	Items []interface{}
	lock  sync.Mutex
}

func NewStack() *Stack {
	s := &Stack{}
	s.Items = []interface{}{}
	return s
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.Items) == 0 {
		return nil
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}
