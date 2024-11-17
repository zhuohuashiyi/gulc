package gulc


import (
	"testing"
)


func TestBloomFilter(t *testing.T) {
	filter := NewBloomFilter[String](128)
	filter.Add(StringPacket("Alice"))
	filter.Add(StringPacket("Bob"))
	filter.Add(StringPacket("Clone"))
	filter.Add(StringPacket("Dawe"))
	filter.Add(StringPacket("Eleps"))

	ret := filter.Search(StringPacket("Dawe"))
	Assert(t, true, ret, "BloomFilter(Dawe)")

	ret = filter.Search(StringPacket("Helen"))
	Assert(t, false, ret, "BloomFilter(Helen)")
	
}


func TestCountBloomFilter(t *testing.T) {
	filter := NewCountBloomFilter[String](128)
	filter.Add(StringPacket("Alice"))
	filter.Add(StringPacket("Bob"))
	filter.Add(StringPacket("Clone"))
	filter.Add(StringPacket("Dawe"))
	filter.Add(StringPacket("Eleps"))

	ret := filter.Search(StringPacket("Dawe"))
	Assert(t, true, ret, "BloomFilter(Search)")

	ret, err := filter.Delete(StringPacket("Dawe"))
	Assert(t, true, ret, "BloomFilter(Delete Dawe)")
	AssertNil(t, err, "BloomFilter(Delete Dawe)")

	ret = filter.Search(StringPacket("Dawe"))
	Assert(t, false, ret, "BloomFilter(Search)")
}


func TestInvertibleBloomFilter(t *testing.T) {
	filter := NewInvertibleBloomFilter(128)
	filter.Add(IntegerPacket(123))
	filter.Add(IntegerPacket(89))
	filter.Add(IntegerPacket(34))
	filter.Add(IntegerPacket(123))

	ret := filter.Search(IntegerPacket(123))
	Assert(t, true, ret, "TestInvertibleBloomFilter(search 123)")

	ret, err := filter.Delete(IntegerPacket(89))
	Assert(t, true, ret, "TestInvertibleBloomFilter(delete 89)")
	AssertNil(t, err, "TestInvertibleBloomFilter(delete 89)")

	ret = filter.Search(IntegerPacket(89))
	Assert(t, false, ret, "TestInvertibleBloomFilter(search 89)")

	rets, err := filter.List()
	Assert(t, 2, len(rets), "TestInvertibleBloomFilter(List)")
	AssertNil(t, err, "TestInvertibleBloomFilter(List)")
	t.Logf("filter.List()=%+v", rets)
}