package gulc

type Character struct {
	idx int
	count int
}


type Heap []Character

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count < h[j].count
	} 
	return h[i].idx < h[j].idx
} 

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(Character))
}

func (h *Heap) Pop() interface{} {
	ret := (*h)[len(*h) - 1]
	*h = (*h)[: len(*h) - 1]
	return ret
}
