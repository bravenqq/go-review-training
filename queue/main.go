// Package main provides ...
package main

import (
	"fmt"
	"sync"
)

func main() {
	q := NewQueue()
	for i := 0; i < 100; i++ {
		fmt.Println("i:", i)
		q.Join(i)
	}

	for {
		a := q.Leave()
		if a == nil {
			break
		}
		fmt.Println(a.(int))
	}
}

type Queue struct {
	Items  []interface{}
	locker sync.Mutex
}

func NewQueue() *Queue {
	q := &Queue{}
	q.Items = []interface{}{}
	return q
}

func (q *Queue) Join(item interface{}) {
	q.locker.Lock()
	defer q.locker.Unlock()
	if item != nil {
		q.Items = append(q.Items, item)
	}
}

func (q *Queue) Leave() interface{} {
	q.locker.Lock()
	defer q.locker.Unlock()
	if len(q.Items) == 0 {
		return nil
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}
