package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter_Unsafe(t *testing.T) {
	cnt := 0
	for i := 0; i < 5000; i++ {
		go func() {
			cnt++
		}()
	}
	time.Sleep(time.Millisecond * 50)
	t.Logf("counter: %d", cnt)
}

func TestCounter_ThreadSafe(t *testing.T) {
	var mut sync.Mutex
	cnt := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer mut.Unlock()
			mut.Lock()
			cnt++
		}()
	}
	time.Sleep(time.Millisecond * 50)
	t.Logf("counter: %d", cnt)
}

func TestCounter_ThreadSafe_WaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	cnt := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer mut.Unlock()
			mut.Lock()
			cnt++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter: %d", cnt)
}
