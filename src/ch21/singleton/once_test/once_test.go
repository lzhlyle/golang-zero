package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var (
	singletonInstance *Singleton
	singletonOnce     sync.Once
)

func NewSingleton() *Singleton {
	singletonOnce.Do(func() {
		fmt.Println("creating...")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := NewSingleton()
			fmt.Printf("%v\n", unsafe.Pointer(obj)) // same address
			wg.Done()
		}()
	}
	wg.Wait()
}
