package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func FirstResponse() string {
	runnerLimit := 10
	ch := make(chan string, runnerLimit)
	for i := 0; i < runnerLimit; i++ {
		go func(i int) {
			ch <- runTask(i)
		}(i)
	}
	return <-ch
}

func runTask(i int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("from [%d]", i)
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) // 当前系统协程数 2
	t.Log(FirstResponse())
	time.Sleep(20 * time.Millisecond)
	t.Log("After:", runtime.NumGoroutine()) // 11 协程未释放，被阻塞
}
