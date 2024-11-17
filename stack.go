package gulc

type Stack[T any] struct {
	elems []T
	ptr   int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{elems: make([]T, 0), ptr: -1}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.ptr == -1
}

func (s *Stack[T]) Push(v T) {
	s.ptr++
	if s.ptr + 1 <= len(s.elems) {
		s.elems[s.ptr] = v
	} else {
		s.elems = append(s.elems, v)
	}
}

func (s *Stack[T]) Top() T {
	return s.elems[s.ptr]
}

func (s *Stack[T]) Pop() T {
	res := s.Top()
	s.ptr--
	return res
}

func (s *Stack[T]) Clear() {
	s.ptr = -1
}
