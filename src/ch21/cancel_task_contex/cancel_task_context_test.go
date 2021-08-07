package cancel_task_contex

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Cancel(t *testing.T) {
	var wg sync.WaitGroup
	ctx, cancelFunc := context.WithCancel(context.Background())
	// 开 5 个协程，每个都在执行任务，知道任务取消
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context, wg *sync.WaitGroup) {
			for {
				if isCanceled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Printf("Task-%d canceled.\n", i)
			wg.Done()
		}(i, ctx, &wg)
	}
	cancelFunc()
	wg.Wait()
}

func isCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
