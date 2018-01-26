package go_concurrent

import (
	"testing"
)

func TestUtilAllTaskFinished(t *testing.T) {
	ret, err := ExampleUtilAllTaskFinished()
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
