package go_concurrent

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

type TestTask struct {
	wg *sync.WaitGroup
}

const NUM_OF_WORKER = 5
const LEN_OF_TASK_QUEUE = 5

func (self *TestTask) Run() {
	time.Sleep(time.Second * 1)
	fmt.Println("task was done!")
	self.wg.Done()
}

func TestSingletonCreation(t *testing.T) {
	poolPtr1 := GetGoroutingPool(LEN_OF_TASK_QUEUE, NUM_OF_WORKER)
	poolPtr2 := GetGoroutingPool(LEN_OF_TASK_QUEUE, NUM_OF_WORKER)
	if poolPtr1 != poolPtr2 {
		t.Error("2 pools have been created!")
	}
}

func TestMulitGoroutingRunning(t *testing.T) {
	var wg sync.WaitGroup
	poolPtr1 := GetGoroutingPool(LEN_OF_TASK_QUEUE, NUM_OF_WORKER)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		poolPtr1.Submit(&TestTask{&wg})
		numOfGoroutings := runtime.NumGoroutine()
		if numOfGoroutings != (NUM_OF_WORKER + 2) {
			t.Errorf("The expected number of the working goroutings is %d"+
				" but the actual number is %d", NUM_OF_WORKER, numOfGoroutings)
		}
	}
	wg.Wait()

}

func TestStopPool(t *testing.T) {
	var numOfGoroutings int
	poolPtr1 := GetGoroutingPool(LEN_OF_TASK_QUEUE, NUM_OF_WORKER)
	numOfGoroutings = runtime.NumGoroutine()
	if numOfGoroutings != (NUM_OF_WORKER + 2) {
		t.Errorf("The expected number of the working goroutings is %d"+
			" but the actual number is %d", NUM_OF_WORKER, numOfGoroutings-2)
	}
	poolPtr1.Stop()
	time.Sleep(time.Second * 1)
	numOfGoroutings = runtime.NumGoroutine()
	if numOfGoroutings != 2 {
		t.Errorf("The expected number of the working goroutings is %d"+
			" but the actual number is %d", 0, numOfGoroutings-2)
	}
}
