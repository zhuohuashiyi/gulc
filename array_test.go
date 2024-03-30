package gulc


import (
	"testing"
)


func TestFindKthNumber(t *testing.T) {
	if FindKthNumber([]int{3, 2, 5, 1}, 2) != 2 {
		t.Errorf("{3, 2, 5, 1}, k=2, it should be 2")
	}
}