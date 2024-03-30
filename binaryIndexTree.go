package gulc


type BinaryIndexTree struct {
	elems []int
}


func NewBinaryIndexTree(nums []int) *BinaryIndexTree {
	n := len(nums)
	tree := &BinaryIndexTree{make([]int, n + 1)}
	for i, num := range nums {
		tree.Add(i + 1, num)
	}
	return tree
}


// 求取前i个数的和
func (tree *BinaryIndexTree) Sum(i int) int {
	sum := 0
	for i > 0 {
		sum += tree.elems[i]
		i -= (i & -i)
	}
	return sum
}


// 第i个数加上delta
func (tree *BinaryIndexTree) Add(i, delta int) {
	for i < len(tree.elems) {
		tree.elems[i] += delta
		i += ( i & -i)
	}
}