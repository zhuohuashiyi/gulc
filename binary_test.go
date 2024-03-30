package gulc


import (
	"testing"
)


func TestReverseBits(t *testing.T) {
	t.Log(ReverseBits(0x80000000))
	t.Log(ReverseBits(2147483648))
}