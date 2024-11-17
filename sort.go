package gulc

type Sorter[T any] struct {
	less func(item1, item2 T) bool
}

func NewSorter[T any](less func(item1, item2 T) bool) *Sorter[T] {
	return &Sorter[T]{less: less}
}

// HeapSort 堆排序，时间复杂度:O(nlogn), 空间复杂度：O(1), 不稳定
func (s *Sorter[T]) HeapSort(arr []T) {
	h := NewHeapWithArr(s.less, arr)
	k := len(arr) - 1
	for !h.IsEmpty() {
		arr[k] = h.Pop()
		k--
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

// MergeSort 归并排序，时间复杂度：O(nlogn), 空间复杂度：O(n), 稳定
func (s *Sorter[T]) MergeSort(arr []T) []T {
	return s.mergeSortV1(arr)
}

// 递归形式的归并排序
func (s *Sorter[T]) mergeSortV1(arr []T) []T {
	if len(arr) <= 1 { // 递归终点
		return arr
	}
	// 递归过程
	leftArr := s.mergeSortV1(append([]T(nil), arr[:len(arr)/2]...))
	rightArr := s.mergeSortV1(append([]T(nil), arr[len(arr)/2:]...))
	// 合并
	i := 0
	j := 0
	k := 0
	for i < len(leftArr) || j < len(rightArr) {
		if i >= len(leftArr) || (j < len(rightArr) && s.less(rightArr[j], leftArr[i])) {
			arr[k] = rightArr[j]
			j++
			k++
		} else {
			arr[k] = leftArr[i]
			i++
			k++
		}
	}
	return arr
}

// 非递归形式的归并排序
func (s *Sorter[T]) mergeSortV2(arr []T) []T {
	segLen := 1 // 比较的段长
	for segLen < len(arr) {
		for p := 0; p < len(arr); p += 2 * segLen {
			leftArr := append([]T(nil), arr[p:min(p+segLen, len(arr))]...)
			rightArr := append([]T(nil), arr[min(p+segLen, len(arr)):min(p+2*segLen, len(arr))]...)
			// 合并
			i := 0
			j := 0
			k := 0
			for i < len(leftArr) || j < len(rightArr) {
				if i >= len(leftArr) || (j < len(rightArr) && s.less(rightArr[j], leftArr[i])) {
					arr[p+k] = rightArr[j]
					j++
					k++
				} else {
					arr[p+k] = leftArr[i]
					i++
					k++
				}
			}
		}
		segLen *= 2
	}
	return arr
}

// 快速排序实现，时间复杂度：O(nlogn)， 空间复杂度：O(1), 不稳定
func (s *Sorter[T]) QuickSort(arr []T) {
	s.quickSortV2(arr, 0, len(arr)-1)
}

// 递归形式的排序
func (s *Sorter[T]) quickSortV1(arr []T, left, right int) {
	if left >= right { // 递归终点
		return
	}
	base := arr[left]
	basePtr := left
	i := left + 1
	for i <= right {
		if s.less(arr[i], base) {
			basePtr++
			if basePtr != i {
				arr[basePtr], arr[i] = arr[i], arr[basePtr]
			}
		}
		i++
	}
	arr[left] = arr[basePtr]
	arr[basePtr] = base
	s.quickSortV1(arr, left, basePtr-1)
	s.quickSortV1(arr, basePtr+1, right)
}

// 非递归实现，借助栈来模拟递归
func (s *Sorter[T]) quickSortV2(arr []T, left, right int) {
	stack := NewStack[[2]int]()
	stack.Push([2]int{left, right})
	for !stack.IsEmpty() {
		state := stack.Pop()
		left, right := state[0], state[1]
		if left >= right {
			continue
		}
		base := arr[left]
		basePtr := left
		i := left + 1
		for i <= right {
			if s.less(arr[i], base) {
				basePtr++
				if basePtr != i {
					arr[basePtr], arr[i] = arr[i], arr[basePtr]
				}
			}
			i++
		}
		arr[left] = arr[basePtr]
		arr[basePtr] = base
		stack.Push([2]int{left, basePtr - 1})
		stack.Push([2]int{basePtr + 1, right})
	}
}

// quickSortV3优化实现，基准取头尾中三数中的平均数
func (s *Sorter[T]) quickSortV3(arr []T, left, right int) {
	if left >= right { // 递归终点
		return
	}
	s.adjustBase(arr, left, right)
	base := arr[left]
	basePtr := left
	i := left + 1
	for i <= right {
		if s.less(arr[i], base) {
			basePtr++
			if basePtr != i {
				arr[basePtr], arr[i] = arr[i], arr[basePtr]
			}
		}
		i++
	}
	arr[left] = arr[basePtr]
	arr[basePtr] = base
	s.quickSortV3(arr, left, basePtr-1)
	s.quickSortV3(arr, basePtr+1, right)
}

// adjustBase 调整基准值，将头尾中三数中的平均数放在头部作为基准
func (s *Sorter[T]) adjustBase(arr []T, left, right int) {
	basePtrs := []int{left, right, (left + right) / 2}
	s1 := NewSorter(func(i, j int) bool {
		return s.less(arr[i], arr[j])
	})
	s1.InsertSort(basePtrs)
	arr[left], arr[basePtrs[1]] = arr[basePtrs[1]], arr[left]
}

// BubbleSort 冒泡排序，时间复杂度：O(n^2) 空间复杂度：O(1), 稳定
func (s *Sorter[T]) BubbleSort(arr []T) {
	n := len(arr)
	for i := 0; i < n; i++ {
		var bubbleFlag bool // 是否有冒泡
		for j := i; j+1 < n; j++ {
			if s.less(arr[j+1], arr[j]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				bubbleFlag = true
			}
		}
		if !bubbleFlag { // 无冒泡，说明已经排序完成，提前退出
			break
		}
	}
}

// SelectSort 选择排序，时间复杂度：(n ^ 2), 空间复杂度:O(1)，不稳定
func (s *Sorter[T]) SelectSort(arr []T) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minPtr := i
		for j := i + 1; j < n; j++ {
			if s.less(arr[j], arr[minPtr]) {
				minPtr = j
			}
		}
		arr[minPtr], arr[i] = arr[i], arr[minPtr]
	}
}

// InsertSort 插入排序，时间复杂度:O(n^2) 空间复杂度:O(1), 稳定
func (s *Sorter[T]) InsertSort(arr []T) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if s.less(arr[j], arr[j-1]) {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}

// 希尔排序，时间复杂度: O(n^(3 / 2)) 空间复杂度：O(1)， 不稳定
func (s *Sorter[T]) ShellSort(arr []T) {
	n := len(arr)
	gap := n >> 1
	for gap > 0 {
		for i := 0; i < gap; i++ {
			for j := i; j < n; j += gap {
				for k := i + gap; k < n; k += gap {
					if s.less(arr[k], arr[k-gap]) {
						arr[k], arr[k-gap] = arr[k-gap], arr[k]
					}
				}
			}
		}
		gap >>= 1
	}
}

type IntSorter struct {
}

func NewIntSorter() *IntSorter {
	return &IntSorter{}
}

// 计数排序，时间复杂度：O(n+k) k为数据范围 空间复杂度：O(k) 稳定
// 适用于数据范围不大的情况下
func (s *IntSorter) CountSort(arr []int64) {
	minv := arr[0]
	maxv := arr[0]
	for _, item := range arr {
		minv = min(minv, item)
		maxv = max(maxv, item)
	}
	counts := make([]int64, maxv-minv+1)
	for _, item := range arr {
		counts[item-minv]++
	}
	k := 0
	for i := int64(0); i < maxv-minv+1; i++ {
		for j := int64(0); j < counts[i]; j++ {
			arr[k] = i + minv
			k++
		}
	}
}

func (s *IntSorter) RadixSort(arr []int) {
	var newBuckets func() []*List[int] = func() []*List[int] {
		buckets := make([]*List[int], 10)
		for i := 0; i < 10; i++ {
			buckets[i] = NewList[int]()
		}
		return buckets
	}
	buckets := newBuckets()
	oldBuckets := newBuckets()
	for _, num := range arr {
		oldBuckets[num%10].AppendNodeTail(NewListNode[int](num/10, num))
	}
	flag := true // 停止标志
	for flag {
		flag = false
		for i := 0; i < 10; i++ {
			for node := oldBuckets[i].Begin(); node != oldBuckets[i].End(); node = node.Next {
				if node.Key != 0 {
					flag = true
				}
				buckets[node.Key%10].AppendNodeTail(NewListNode[int](node.Key/10, node.Val))
			}
		}
		oldBuckets = buckets
		buckets = newBuckets()
	}
	// 此时所有数据都在零号桶里
	k := 0
	for node := oldBuckets[0].Begin(); node != oldBuckets[0].End(); node = node.Next {
		arr[k] = node.Val
		k++
	}
}
