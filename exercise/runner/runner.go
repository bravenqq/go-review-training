// Package runner provides ...
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

/*runner监控任务是否在规定时间内执行，并且在接收到中段信号时结束任务*/

var (
	ErrTimeOut   = errors.New("received timeout")
	ErrInterrupt = errors.New("received interrupt")
)

type Runner struct {
	interrupt chan os.Signal
	timeout   <-chan time.Time
	complete  chan error
	tasks     []func(int)
}

func NewRunner(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		timeout:   time.After(d),
		complete:  make(chan error),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
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
