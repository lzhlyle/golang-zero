package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	a, b := 1, 1
	for i := 0; i < 10; i++ {
		fmt.Print(a, " ")
		//t.Log(a)
		a, b = b, a+b
	}
	fmt.Println()
}
