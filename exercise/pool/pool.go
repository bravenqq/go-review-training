// Package pool provides ...
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

var ErrPoolClosed = errors.New("error pool has closed")

/* pool管理共享资源的使用*/

type Pool struct {
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
	m         sync.Mutex
}

func New(size int, factory func() (io.Closer, error)) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}

	return &Pool{resources: make(chan io.Closer, size), factory: factory}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case res, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		return res, nil
	default:
		return p.factory()

	}
}

func (p *Pool) Release(res io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		res.Close()
		return
	}

	select {
	case p.resources <- res:
		log.Println("release in queue")
	default:
		log.Println("release closed")
		res.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}
