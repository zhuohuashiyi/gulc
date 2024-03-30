package gulc


import (
	"testing"
)


func TestBinaryIndexTree(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	tree := NewBinaryIndexTree(nums)
	if tree.Sum(1) != 1 {
		t.Error(1)
	}
	if tree.Sum(2) != 3 {
		t.Error(2)
	}
	t.Log(tree.Sum(3))
	t.Log(tree.Sum(4))
}