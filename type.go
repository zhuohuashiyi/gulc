/*
	基础类型包装
*/
package gulc

type String struct {
	s string
}

// 生成包装类型
func StringPacket(s string) String {
	return String{s: s}
}

func (s String) DePacket() string {
	return s.s
}

// 基础类型实现Hashable
func (s String) Hash() []byte {
	return []byte(s.s)
}

type Integer struct {
	i int
}

func IntegerPacket(i int) Integer {
	return Integer{i: i}
}

func (i Integer) Depacket() int {
	return i.i
}

func (i Integer) Hash() []byte {
	return Int2Byte(i.i)
}



