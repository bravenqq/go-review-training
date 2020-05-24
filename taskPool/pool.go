// Package pool provides ...
package pool

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type Pool struct {
	max int
	sem *semaphore.Weighted
}

func NewPool(max int) *Pool {
	//TODO
	return &Pool{}
}

func (p *Pool) Get() {
	p.sem.Acquire(context.Background(), 1)

}
