package gulc

import "testing"

func TestSort(t *testing.T) {
	arr := []int{15,11,101,223,52,1111}
	sorter := NewIntSorter()
	sorter.RadixSort(arr)
	t.Logf("sorted arr=%+v", arr)
}