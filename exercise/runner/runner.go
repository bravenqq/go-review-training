// Package runner provides ...
package runner

import (
	"errors"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*runner监控任务是否在规定时间内执行，并且在接收到中段信号时结束任务*/

var (
	ErrTimeOut   = errors.New("received timeout")
	ErrInterrupt = errors.New("received interrupt")
)

type Runner struct {
	worker    int
	interrupt chan os.Signal
	timeout   <-chan time.Time
	complete  chan error
	tasks     chan func(int)
	wg        sync.WaitGroup
}

func NewRunner(d time.Duration, worker int) *Runner {
	return &Runner{
		worker:    worker,
		interrupt: make(chan os.Signal, 1),
		timeout:   time.After(d),
		complete:  make(chan error),
		tasks:     make(chan func(int)),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	go func() {
		for _, task := range tasks {
			r.tasks <- task
		}
	}()
}

func (r *Runner) Start() error {
	r.wg.Add(r.worker)
	for i := 0; i < r.worker; i++ {
		go func() {
			r.complete <- r.run()
		}()
	}

	go func() {
		r.wg.Wait()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

func (r *Runner) run() error {
	for task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(1)
	}
	r.wg.Done()
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
