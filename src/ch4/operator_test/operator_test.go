package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{2, 3, 4}
	//c := [...]int{1, 2, 3, 4}
	d := [...]int{1, 2, 3}

	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	curr := 0b0111
	t.Log(curr&Readable == Readable, curr&Writable == Writable, curr&Executable == Executable)

	curr &^= Readable // 清掉可读性
	t.Log(curr&Readable == Readable, curr&Writable == Writable, curr&Executable == Executable)
}
