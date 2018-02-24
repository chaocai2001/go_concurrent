package go_concurrent

import (
	"testing"
	"time"
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
		t.Logf("The sum is %d, which is less than the expected value %d for cancelling\n", ret, expRet)
	}
}

func TestExampleUtilAnyoneResponse(t *testing.T) {
	ret := ExampleUtilAnyoneResponse()
	d1, ok := ret.(time.Duration)
	if !ok {
		t.Errorf("The return value should be time.Duration")
	}
	if d1.Seconds() != 1 {
		t.Errorf("Expected the return value 1s")
	}
	t.Log(ret)
}
