// Package main provides ...
package main

func main() {
	tasks := make(chan func() int)
	done := make(chan struct{})
	for {
		select {
		case task := <-tasks:
			task()
		case <-done:
			return
		}
	}
}
