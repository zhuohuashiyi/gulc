package gulc

import (
	"testing"
)

func TestHeap(t *testing.T) {
	arr := []int{1, 4, 2, -1}
	h := NewHeapWithArr(func(item1, item2 int) bool {
		return item1 > item2
	}, arr)
	if h.Top() != 4 {
		t.Errorf("err: h.Top() not equal to %d, actual it is %d", 4, h.Top())
	}
	h.Pop()
	if h.Top() != 2 {
		t.Errorf("err: h.Top() not equal to %d, actual it is %d", 2, h.Top())
	}
	h.Push(99)
	if h.Top() != 99 {
		t.Errorf("err: h.Top() not equal to %d, actual it is %d", 99, h.Top())
	}
}