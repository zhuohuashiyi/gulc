package gulc

import "testing"

func TestSkiplist(t *testing.T) {
	sl := NewSkiplist()
	sl.Insert(30)
	sl.Insert(40)
	ret := sl.Search(30)
	Assert(t, true, ret, "TestSkiplist search")
	sl.Delete(30)
	ret = sl.Search(30)
	Assert(t, false, ret, "TestSkiplist search")
}