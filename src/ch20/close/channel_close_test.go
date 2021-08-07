package close

import (
	"fmt"
	"sync"
	"testing"
)

func publish(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func Test_Main(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	publish(ch, &wg)

	wg.Add(1)
	consumer(ch, &wg)

	wg.Add(1)
	consumer(ch, &wg)

	wg.Wait()
}
