package array_test

import (
	"testing"
)

func TestArrayInt(t *testing.T) {
	arr := [...]int{1, 2, 3}
	t.Log(len(arr), arr)
}

func TestArraySection(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	t.Log(a)
	t.Log(a[1:2])
	t.Log(a[1:3])
	t.Log(a[1:len(a)])
	t.Log(a[1:])
	t.Log(a[:3])

	t.Logf("%T", a[1:3])
}
