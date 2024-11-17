package gulc

type Deque[T any] interface {
	RemoveHead() T
	AppendTail(v T)
}