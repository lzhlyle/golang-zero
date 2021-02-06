package _map

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMyIntSet(t *testing.T) {
	mySet := NewMyIntSet(8)

	t.Log(mySet.Contains(5)) // false
	t.Log(mySet.Remove(5))   // false
	t.Log(mySet.Add(5))      // true
	t.Log(mySet.Contains(5)) // true
	t.Log(mySet.Add(5))      // false
	t.Log(mySet.Remove(5))   // true
}

type MyIntSet struct {
	m map[int]bool
}

func NewMyIntSet(size int) *MyIntSet {
	return &MyIntSet{
		m: make(map[int]bool, size),
	}
}

func (s *MyIntSet) Contains(v int) bool {
	_, ex := s.m[v]
	return ex
}

func (s *MyIntSet) Add(v int) bool {
	if _, ex := s.m[v]; ex {
		return false
	} else {
		s.m[v] = true
		return true
	}
}

func (s *MyIntSet) Remove(v int) bool {
	if _, ex := s.m[v]; ex {
		delete(s.m, v)
		return true
	} else {
		return false
	}
}
