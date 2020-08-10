// Package async provides ...
package async

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskPoolRunSuccess(t *testing.T) {
	startedTask1 := false
	finishedTask1 := false
	task1 := func() error {
		defer func() { finishedTask1 = true }()
		startedTask1 = true
		time.Sleep(time.Millisecond * 100)
		return nil
	}
	fmt.Println()

	startedTask2 := false
	finishedTask2 := false
	task2 := func() error {
		defer func() {
			finishedTask2 = true
		}()
		startedTask2 = true
		time.Sleep(time.Millisecond * 100)

		return nil
	}

	startedTask3 := false
	finishedTask3 := false
	started := make(chan bool, 1)
	task3 := func() error {
		defer func() {
			finishedTask3 = true
			started <- true
		}()
		startedTask3 = true
		time.Sleep(time.Millisecond * 200)
		return nil
	}

	pool := NewTaskPool(2)

	// act
	errc := pool.Run(context.Background(), task1)

	assert.NotNil(t, errc)

	errc = pool.Run(context.Background(), task2)
	assert.NotNil(t, errc)

	errc = pool.Run(context.Background(), task3)
	assert.NotNil(t, errc)

	<-started
	defer close(started)

	assert.True(t, startedTask1)
	assert.True(t, finishedTask1)

	assert.True(t, startedTask2)
	assert.True(t, finishedTask2)

	assert.True(t, startedTask3)
	assert.True(t, finishedTask3)
}
