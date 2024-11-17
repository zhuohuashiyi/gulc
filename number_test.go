package gulc


import (
	"testing"
)


func TestAddNumber(t *testing.T) {
	num1 := "1b"
	num2 := "2x"
	res := AddNumber(num1, num2, 16)
	if res != "48" {
		t.Errorf("AddNumber(base %d) error: %s+%s != %s, actually result is %s", 16, num1, num2, "48", res)
	}
}