// Package main provides ...
package main

import (
	"container/heap"
	"math/rand"
	"time"
)

var nWorker int = 20

type Request struct {
	fn func() int
	c  chan int
}

func requester(work chan Request) {
	c := make(chan int)
	for {
		time.Sleep(rand.Int63(nWorker * 2e9))
		work <- Request{workFn, c}
		result := <-c
		furtherProcess(result)
	}
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests
		req.c <- req.fn()
		done <- w
	}
}

type Pool []*Worker
type Balancer struct {
	pool Pool
	down chan *Worker
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work:
			b.dispatch(req)
		case w := <-b.down:
			b.complete(w)

		}
	}
}

func (b *Balancer) dispatch(req Request) {
	w := heap.Pop(&b.pool).(*Worker)
	w.requests <- req
	w.peding++
	heap.Push(&b.pool, w)
}

func (b *Balancer) complete(w *Worker) {
	w.peding--
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w)
}

func (p Pool) Less(i, j int) bool {
	return p[i].peding < p[j].peding
}

type Worker struct {
	requests chan Request
	peding   int
	index    int
}
