package oop

import (
	"fmt"
	"math/rand"
	"testing"
	"unsafe"
)

type Employee struct {
	Id     string
	Name   string
	Age    int
	RdVal  int
	secret string
}

func (e Employee) ChangeRandValueObj() {
	fmt.Printf("Addr: %x\n", unsafe.Pointer(&e))
	e.RdVal = rand.Int()
}

func (e *Employee) ChangeRandValuePtr() {
	fmt.Printf("Addr: %x\n", unsafe.Pointer(e))
	e.RdVal = rand.Int()
}

func TestChangeRandValueObj(t *testing.T) {
	e := Employee{}
	fmt.Printf("Addr: %x\n", unsafe.Pointer(&e))

	e.ChangeRandValueObj()
	t.Log(e.RdVal)
	e.ChangeRandValuePtr()
	t.Log(e.RdVal)
}

func TestChangeRandValuePtr(t *testing.T) {
	e := &Employee{}
	fmt.Printf("Addr: %x\n", unsafe.Pointer(e))

	e.ChangeRandValueObj()
	t.Log(e.RdVal)
	e.ChangeRandValuePtr()
	t.Log(e.RdVal)
}
