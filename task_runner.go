/**
The package provides the common APIs for concurrent processing

@Auth: Chao Cai
Created at 2018-1-26
**/
package go_concurrent

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Runnable interface {
	Run()
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

//The following is the example of waiting for all tasks done
type enumTask struct {
	startFrom int
	result    int
	err       error
}

func (t *enumTask) Run() {
	fmt.Printf("Start %d ", t.startFrom)
	for i := t.startFrom; i < t.startFrom+10; i++ {
		t.result += i
	}
	fmt.Println(t.result)
}

func ExampleUtilAllTaskFinished() (int, error) {
	t1 := enumTask{1, 0, nil}
	t2 := enumTask{11, 0, nil}
	t3 := enumTask{21, 0, nil}

	tasks := []Runnable{&t1, &t2, &t3}

	err = UtilAllTaskFinished(tasks)
	if err != nil {
		return 0, err
	}
	ret := 0
	for _, task := range tasks {
		ret += (task.(*enumTask)).result
	}
	return ret, nil
}
