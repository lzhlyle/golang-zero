package string

import "testing"

func TestString(t *testing.T) {
	s := "中国"
	t.Logf("len(s): %v", len(s)) // 6

	c := []rune(s)
	t.Logf("len(c): %v, type(c): %T", len(c), c) // 2

	t.Logf("unicode: %x", c) // [4e2d 56fd]
	t.Logf("UTF8: %x", s)    // e4b8ade59bbd

	for i, ch := range s {
		t.Logf("i:(%v), ch:(%[2]v), type(ch):(%[2]T), unicode(ch):(%[2]x)", i, ch)
	}

}
