package obj_pool

import (
	"errors"
	"testing"
	"time"
)

type ReusableObj struct{}

type ObjPool struct {
	ch chan *ReusableObj
}

func NewMyObjPool(cap int) *ObjPool {
	res := &ObjPool{
		ch: make(chan *ReusableObj, cap),
	}
	for i := 0; i < cap; i++ {
		res.ch <- &ReusableObj{}
	}
	return res
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case res := <-p.ch:
		return res, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.ch <- obj:
		return nil
	default:
		return errors.New("over flow")
	}
}

func TestObjPool(t *testing.T) {
	limit := 5
	times := limit << 1
	pool := NewMyObjPool(limit)

	// test overflow
	if err := pool.ReleaseObj(nil); err != nil {
		t.Log("testing overflow: ", err)
	}

	for i := 0; i < times+1; i++ {
		if obj, err := pool.GetObj(100 * time.Millisecond); err != nil {
			// test timeout
			t.Log("testing timeout: ", err)
		} else {
			t.Logf("obj addr:[%x]", obj)
			if i < times>>1 {
				if err := pool.ReleaseObj(obj); err != nil {
					t.Log(err)
				}
			} else {
				// test not release
				t.Log("testing not release...")
			}
		}
	}
}
