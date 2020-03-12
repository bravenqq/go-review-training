// Package work provides ...
package work

import "sync"

type Pool struct {
	works chan Worker
	wg    sync.WaitGroup
}

type Worker interface {
	Task()
}

func New(maxGoroutines int) *Pool {
	pool := &Pool{works: make(chan Worker)}
	pool.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for t := range pool.works {
				t.Task()
			}
			defer pool.wg.Done()
		}()
	}
	return pool
}

func (p *Pool) Run(work Worker) {
	p.works <- work
}

func (p *Pool) Shutdown() {
	close(p.works)
	p.wg.Wait()
}
