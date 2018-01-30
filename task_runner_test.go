package go_concurrent

import (
	"fmt"
	"testing"
)

func TestUtilAllTaskFinished(t *testing.T) {
	ret, err := ExampleUtilAllTaskFinishedWithTimeout()
	if err != nil {
		t.Error(err)
	}
	expRet := 0
	for i := 1; i <= 30; i++ {
		expRet += i
	}
	if ret != expRet {
		t.Errorf("expected ret is %d, but actual ret is %d", expRet, ret)
	}

}

func TestUtilAllTaskFinishedWithTimeout(t *testing.T) {
	ret, err := ExampleUtilAllTaskFinishedWithTimeout()
	if err != nil {
		t.Error(err)
	}
	expRet := 0
	for i := 1; i <= 30; i++ {
		expRet += i
	}
	if ret != expRet {
		t.Errorf("expected ret is %d, but actual ret is %d", expRet, ret)
	}

}

func TestExampleUtilAllTaskFinishedWithTimeout_TimeoutOccurred(t *testing.T) {
	ret, err := ExampleUtilAllTaskFinishedWithTimeout_TimeoutOccurred()
	if err != TimeOutError {
		t.Error(err)
	}
	expRet := 0
	for i := 1; i <= 30; i++ {
		expRet += i
	}
	if ret >= expRet {
		t.Errorf("failed to cancel")
	} else {
		fmt.Printf("The sum is %d, which is less than the expected value %d for cancelling\n", ret, expRet)
	}
}
