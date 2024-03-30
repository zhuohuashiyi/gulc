package gulc

type Stack struct {
	elems []int
	ptr   int
}

func NewStack(n int) *Stack {
	return &Stack{make([]int, n), -1}
}

func (s *Stack) IsEmpty() bool {
	return s.ptr == -1
}

func (s *Stack) Push(v int) {
	s.ptr++
	s.elems[s.ptr] = v
}

func (s *Stack) Top() int {
	return s.elems[s.ptr]
}

func (s *Stack) Pop() int {
	res := s.Top()
	s.ptr--
	return res
}

func (s *Stack) Clear() {
	s.ptr = -1
}
