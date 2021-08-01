package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done!"
}

func otherService() {
	fmt.Println("Working on something else...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherService()
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherService()
	fmt.Println(<-retCh)
}
