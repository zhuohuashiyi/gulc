package gulc

import "testing"


func TestPow(t *testing.T) {
	a := 2
	b := 4
	c := Pow(a, b)
	t.Logf("get c=%d", c)	
}

