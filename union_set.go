package gulc


type UnionSet struct {
	parent []int
}

// 新建一个长度为n的并查集
func NewUnionSet(n int) *UnionSet {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	return &UnionSet{parent: parent}
}

// FindRoot 找到其所在的集合的根节点
func (u *UnionSet) FindRoot(node int) int {
	root := node
	for u.parent[root] >= 0 {
		root = u.parent[root]
	}
	return root
}

// MergeRoot 合并两个集合，传入的应该是集合的根节点
func (u *UnionSet) MergeRoot(root1, root2 int) {
	if u.parent[root1] < u.parent[root2] {  // 说明root2集合元素更多
		u.parent[root2] += u.parent[root1]
		u.parent[root1] = root2
		return
	} 
	u.parent[root1] += u.parent[root2]
	u.parent[root2] = root1
}

// Merge 合并两个集合
func (u *UnionSet) Merge(node1, node2 int) {
	root1 := u.FindRoot(node1)
	root2 := u.FindRoot(node2)
	u.MergeRoot(root1, root2)
}

func (u *UnionSet) InSame(node1, node2 int) bool {
	root1 := u.FindRoot(node1)
	root2 := u.FindRoot(node2)
	return root1 == root2
}