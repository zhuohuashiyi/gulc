package gulc

type Character struct {
	idx int
	count int
}


type Heap1 []Character

func (h Heap1) Len() int {
	return len(h)
}

func (h Heap1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap1) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count < h[j].count
	} 
	return h[i].idx < h[j].idx
} 

func (h *Heap1) Push(v interface{}) {
	*h = append(*h, v.(Character))
}

func (h *Heap1) Pop() interface{} {
	ret := (*h)[len(*h) - 1]
	*h = (*h)[: len(*h) - 1]
	return ret
}
