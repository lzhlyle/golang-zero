package slice_test

import (
	"testing"
	"unsafe"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//t.Log(s2[3]) // index out of range

	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2)) // 4, 5
	t.Log(s2[3])

	s2 = append(s2, 2, 3)
	t.Log(len(s2), cap(s2)) // 6, 10=5<<1
	t.Log(s2)
}

func TestSliceGrowing(t *testing.T) {
	s := make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceGrowingAndArray(t *testing.T) {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	t.Logf("%[1]v, %[1]T, %[2]x", arr, unsafe.Pointer(&arr[0]))

	s1 := arr[:]
	t.Logf("%[1]v, %[1]T, %[2]x", s1, unsafe.Pointer(&s1[0])) // 无需扩容，原数组下一个元素被修改

	s1 = append(s1, -1)
	t.Logf("%[1]v, %[1]T, %[2]x", s1, unsafe.Pointer(&s1[0])) // 需要扩容，指向一个新的数组
	t.Logf("%[1]v, %[1]T, %[2]x", arr, unsafe.Pointer(&arr[0]))

	// 故切片的 append 是先扩容(保证内存足够)再追加
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // 3, 9=到year最后一共9个

	Q3 := year[6:9]
	t.Log(Q3, len(Q3), cap(Q3)) // 3, 6=到year最后一共6个
}
