package type_test

import "testing"

func TestPoint(t *testing.T) {
	a := 1
	aPrt := &a
	t.Log(a, aPrt)
	t.Logf("%T %T", a, aPrt)
}
