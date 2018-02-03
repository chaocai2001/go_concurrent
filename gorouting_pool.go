/**
The package provides the common APIs for concurrent processing
The gorouting pool.
Normally it is not worth to use a pool for gorouting creationg.
Because gorouting creation is quite cheap.
But, when many "more.stack" operations happen, it is better to use a pool to avoid of extending gorouting repeatly

@Auth: Chao Cai
Created at 2018-2-3
**/
package go_concurrent

import (
	"sync"
	"time"
)

type Pool struct {
	tasks chan Runnable
}

var once sync.Once
var poolPtr *Pool

func worker(runnables chan Runnable) {
	for {
		select {
		case runnable, ok := <-runnables:
			if !ok {
				return
			}
			time.Sleep(1 * time.Millisecond)
			runnable.Run()

		}

	}

}

func GetGoroutingPool(taskQueueSize int, numOfWorker int) *Pool {
	once.Do(
		func() {

			pool := Pool{

				tasks: make(chan Runnable, taskQueueSize),
			}
			poolPtr = &pool

			for i := 0; i < numOfWorker; i++ {
				go worker(pool.tasks)
			}

		})
	return poolPtr
}

func (pool *Pool) Submit(runner Runnable) {
	pool.tasks <- runner
}

func (pool *Pool) Stop() {
	close(pool.tasks)
}
