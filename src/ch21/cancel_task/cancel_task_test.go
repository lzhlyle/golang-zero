package cancel_task

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Cancel(t *testing.T) {
	var wg sync.WaitGroup
	cancelCh := make(chan struct{}, 0)
	// 开 5 个协程，每个都在执行任务，知道任务取消
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, ch chan struct{}, wg *sync.WaitGroup) {
			for {
				if isCanceled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Printf("Task-%d canceled.\n", i)
			wg.Done()
		}(i, cancelCh, &wg)
	}
	//cancelByEle(cancelCh) // 只有一组收到
	cancelByClose(cancelCh) // 所有组都收到
	wg.Wait()
}

func cancelByClose(ch chan struct{}) {
	close(ch)
}

func cancelByEle(ch chan struct{}) {
	ch <- struct{}{}
}

func isCanceled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
