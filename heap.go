package gulc

type Heap[T any] struct {
	less   func(item1, item2 T) bool
	array  []T
	length int // 标识实际长度
}

// NewHeap 构建一个空的堆
func NewHeap[T any](less func(item1, item2 T) bool) *Heap[T] {
	return &Heap[T]{less: less, array: make([]T, 0), length: 0}
}

// NewHeapWithArr 根据数组构建一个堆
func NewHeapWithArr[T any](less func(item1, item2 T) bool, arr []T) *Heap[T] {
	h := &Heap[T]{less: less, array: arr, length: len(arr)}
	for i := 1; i < len(arr); i++ {
		h.siftUp(i)
	}
	return h
}

// Push 往堆h中插入一个元素
func (h *Heap[T]) Push(v T) {
	h.pushBack(v)
	h.siftUp(h.length - 1)
}

// Top 返回堆顶的元素
func (h *Heap[T]) Top() T {
	return h.array[0]
}

// Pop 移除堆顶的元素并返回
func (h *Heap[T]) Pop() T {
	ret := h.array[0]
	h.array[0] = h.array[h.length-1]
	h.length--
	h.siftDown(0)
	return ret
}

func (h *Heap[T]) ToArr() []T {
	return h.array
}

func (h *Heap[T]) IsEmpty() bool {
	return h.length == 0
}

// pushBack 在尾部插入一个元素
func (h *Heap[T]) pushBack(v T) {
	if len(h.array) == h.length {
		h.array = append(h.array, v)
	} else {
		h.array[h.length] = v
	}
	h.length++
}

// siftUp 从pos位置处开始往上调整
func (h *Heap[T]) siftUp(pos int) {
	if pos < 0 || pos >= h.length {
		return
	}
	i := pos
	j := (pos - 1) / 2
	for j >= 0 && h.less(h.array[i], h.array[j]) {
		h.array[i], h.array[j] = h.array[j], h.array[i]
		i = j
		j = (i - 1) / 2
	}
}

// siftDown 从pos位置处开始往下调整
func (h *Heap[T]) siftDown(pos int) {
	if pos < 0 || pos >= len(h.array) {
		return
	}
	i := 0
	j1 := 2*i + 1
	j2 := 2*i + 2
	for j1 < h.length {
		if h.less(h.array[j1], h.array[i]) && ((j2 >= h.length) || h.less(h.array[j1], h.array[j2])) {
			h.array[i], h.array[j1] = h.array[j1], h.array[i]
			i = j1
		} else if h.less(h.array[j2], h.array[i]) {
			h.array[i], h.array[j2] = h.array[j2], h.array[i]
			i = j2
		} else {
			break
		}
		j1 = 2*i + 1
		j2 = 2*i + 2
	}
}
