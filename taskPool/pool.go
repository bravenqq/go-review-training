// Package async provides ...
package async

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type Task func() error

type TaskPool struct {
	max int
	sem *semaphore.Weighted
}

func NewTaskPool(max int) *TaskPool {
	if max <= 0 {
		panic("max must be a value of >= 1")
	}

	return &TaskPool{max: max, sem: semaphore.NewWeighted(int64(max))}
}

func (p *TaskPool) Run(ctx context.Context, task Task) <-chan error {
	errc := make(chan error, 1)
	err := p.sem.Acquire(ctx, 1)
	if err != nil {
		errc <- err
		close(errc)
		return errc
	}
	go func() {
		defer p.sem.Release(1)
		defer close(errc)
		err = task()
		if err != nil {
			errc <- err
		}
	}()
	return errc
}

func (p *TaskPool) Wait() error {
	for i := 0; i < p.max; i++ {
		err := p.sem.Acquire(context.Background(), 1)
		if err != nil {
			return err
		}
	}
	p.sem.Release(int64(p.max))
	return nil
}
