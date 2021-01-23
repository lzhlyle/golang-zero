package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Open    = 1 << iota // 第 1 位为 1，其他为 0
	Close               // 第 2 位为 1，其他为 0
	Pending             // 第 3 位为 1，其他为 0
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantDay(t *testing.T) {
	t.Log(Monday, Tuesday, Sunday)
}

func TestStatus(t *testing.T) {
	t.Log(Open, Close, Pending)
}

func TestPermission(t *testing.T) {
	curr := 5
	t.Log(curr&Readable == Readable, curr&Writable == Writable, curr&Executable == Executable)
}
