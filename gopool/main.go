// Package main provides ...
package main

import (
	"errors"
	"sync"
)

var ErrInvalidPoolCap = errors.New("invalid pool cap")

type Adder interface {
	Add() int
}

type Sum struct {
	a, b int
}

func (s Sum) Add() int {
	return s.a + s.b
}

type Pool struct {
	wg       sync.WaitGroup //监视所有work结束后结束任务
	closeC   chan struct{}  //发送结束通知给工作者
	capacity int            //并发控制数
	used     int            //已使用的work
	closed   bool
	works    chan Adder
}

func NewPool(capacity int) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}
	return &Pool{capacity: capacity, works: make(chan Adder, capacity), closeC: make(chan struct{})}, nil
}

func (p *Pool) Put(w Adder) {
	if p.used < p.capacity {
		p.used++
		p.wg.Add(1)
		p.run()
	}
	p.works <- w
}

func (p *Pool) Close() error {
	if p.closed {
		return errors.New("已关闭")
	}
	close(p.closeC)
	close(p.works)
	p.wg.Wait()
	p.closed = true
	return nil
}

func (p *Pool) run() {
	go func() {
		for {
			select {
			case w := <-p.works:
				w.Add()
			case <-p.closeC:
				p.wg.Done()
				return
			}
		}
	}()
}
