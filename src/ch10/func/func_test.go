package _func

import (
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Logf("a:%v, b:%v", a, b)
}

func timeSpent(inner func(op int) int) func(op int) (int, float64) {
	return func(op int) (int, float64) {
		start := time.Now()
		res := inner(op)
		return res, time.Since(start).Seconds()
	}
}

func TestTimeSpent(t *testing.T) {
	fib := func(n int) int {
		a, b := 1, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
		}
		return a
	}
	tsFn := timeSpent(fib)
	res, ts := tsFn(50)
	t.Log(res, ts)
	t.Log(tsFn(50)) // t.Log(timeSpent(fib)(50))
}

func sum(ops ...int) int {
	res := 0
	for _, v := range ops { // []int
		res += v
	}
	return res
}

func TestSum(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
	t.Log(sum(1, 2, 3, 4, 5))
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("defer...")
	}()

	t.Log("run...")

	t.Fail()
}
