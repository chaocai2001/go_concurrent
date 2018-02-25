/**
The package provides the common APIs for concurrent processing

@Auth: Chao Cai
Created at 2018-1-26
**/
package go_concurrent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Callable interface {
	Call() interface{}
}

type Runnable interface {
	Run()
}

type RunnableAndCancellable interface {
	RunWithContext(ctx context.Context)
}

//The current process will be blocked util all the tasks get done
func UtilAllTaskFinished(runners []Runnable) {
	var wg sync.WaitGroup
	for _, runner := range runners {
		wg.Add(1)
		go func(r Runnable) {
			r.Run()
			wg.Done()
		}(runner)
	}
	wg.Wait()
}

var TimeOutError = errors.New("Timeout occured!")

func UtilAllTaskFinishedWithTimeout(runners []RunnableAndCancellable, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	var endChan chan struct{} = make(chan struct{}, 1)
	var wg sync.WaitGroup
	//defer close(endChan) : Can't put the statement here, otherwise "panic for sending on closed chan" would occur
	for _, runner := range runners {
		wg.Add(1)
		go func(r RunnableAndCancellable, ctx context.Context) {
			r.RunWithContext(ctx)
			wg.Done()
		}(runner, ctx)
	}
	go func(eChan chan struct{}) {
		wg.Wait()
		endChan <- struct{}{}
	}(endChan)

	select {
	case <-endChan:
		return nil
	case <-time.After(timeout):
		cancel()
		return TimeOutError
	}

}

func UntilAnyoneResponse(callables []Callable) interface{} {
	ch := make(chan interface{}, len(callables))
	for _, callable := range callables {
		go func(c Callable) {
			ch <- c.Call()
		}(callable)
	}
	return <-ch
}

//The following is the example of waiting for all tasks done
type enumTask struct {
	startFrom int
	result    int
	err       error
	sleepTime time.Duration
}

//The task is required to implement Runnable interface
func (t *enumTask) Run() {
	for i := t.startFrom; i < t.startFrom+10; i++ {
		t.result += i
	}
	fmt.Printf("Start from %d, ret = %d \n ", t.startFrom, t.result)
}

//The task is required to implement RunnableAndCancellable interface
func (t *enumTask) RunWithContext(ctx context.Context) {
	for i := t.startFrom; i < t.startFrom+10; i++ {
		t.result += i
		select {
		case <-ctx.Done(): //ready for cancel
			return
		default:
			time.Sleep(t.sleepTime)
		}
	}
	fmt.Printf("Start from %d, ret = %d \n ", t.startFrom, t.result)
}

//All the example is about caculating the sum : 1+2+3+..+30
//Three goroutings are used to calcuate the sum:
//1+2+...+10
//11+12+...+20
//21+22+...+30
func ExampleUtilAllTaskFinished() (int, error) {
	t1 := enumTask{1, 0, nil, 0}  //1+2+3+...+10
	t2 := enumTask{11, 0, nil, 0} //11+12+13...+20
	t3 := enumTask{21, 0, nil, 0} //21+22+23+...+30
	tasks := []Runnable{&t1, &t2, &t3}
	UtilAllTaskFinished(tasks) //will be blocked until all tasks get done
	ret := 0
	for _, task := range tasks { //sum the results from the goroutings
		ret += (task.(*enumTask)).result
	}
	return ret, nil
	//Output 465, nil
}

func ExampleUtilAllTaskFinishedWithTimeout() (int, error) {
	t1 := enumTask{1, 0, nil, time.Millisecond * 1}
	t2 := enumTask{11, 0, nil, time.Millisecond * 1}
	t3 := enumTask{21, 0, nil, time.Millisecond * 1}
	tasks := []RunnableAndCancellable{&t1, &t2, &t3}
	err := UtilAllTaskFinishedWithTimeout(tasks, time.Millisecond*50)
	ret := 0
	for _, task := range tasks {
		ret += (task.(*enumTask)).result
	}
	return ret, err
	//Output 465,nil (All the tasks can be finished with the expected time spent)
}

func ExampleUtilAllTaskFinishedWithTimeout_TimeoutOccurred() (int, error) {
	t1 := enumTask{1, 0, nil, time.Millisecond * 4}
	t2 := enumTask{11, 0, nil, time.Millisecond * 4}
	t3 := enumTask{21, 0, nil, time.Millisecond * 4}
	tasks := []RunnableAndCancellable{&t1, &t2, &t3}
	err := UtilAllTaskFinishedWithTimeout(tasks, time.Millisecond*2)
	ret := 0
	for _, task := range tasks {
		ret += (task.(*enumTask)).result
	}
	return ret, err
	//Output ?,TimeOutError (The tasks can not be finished in the expected time duration)
}

//The following example is for only waiting the first response
type numberCreationTask struct {
	sleepTime time.Duration
}

//The task is required to implement Callable interface
func (t *numberCreationTask) Call() interface{} {
	time.Sleep(t.sleepTime)
	return t.sleepTime
}

func ExampleUtilAnyoneResponse() interface{} {
	callables := make([]Callable, 3)
	for i := 1; i <= 3; i++ {
		t := numberCreationTask{time.Second * time.Duration(i)}
		callables[i-1] = &t
	}
	return UntilAnyoneResponse(callables)
	//Ouput 1s
}
