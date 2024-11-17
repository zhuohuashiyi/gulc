package gulc


import (
	"testing"
)


func TestReverseBits(t *testing.T) {
	t.Log(ReverseBits(0x80000000))
	t.Log(ReverseBits(2147483648))
}

func TestInt2Byte(t *testing.T) {
	i := 1
	data := Int2Byte(i)
	t.Logf("%+v", data) 
}
