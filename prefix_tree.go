package gulc


type PrefixTreeNode[T comparable] struct {
	key T    // 当前树节点名称
	children map[T]*PrefixTreeNode[T]  // 子节点
	isEnd  bool  // 表示是否有一个元素结束在该节点上
	value interface{}
}

func NewPrefixTreeNode[T comparable](key T, value interface{}, isEnd bool) *PrefixTreeNode[T] {
	return &PrefixTreeNode[T]{key: key, value: value, isEnd: isEnd, children: make(map[T]*PrefixTreeNode[T])}
}


type PrefixTree[T comparable] struct {
	root *PrefixTreeNode[T]
}

func NewPrefixTree[T comparable]() *PrefixTree[T] {
	return &PrefixTree[T]{root: &PrefixTreeNode[T]{children: make(map[T]*PrefixTreeNode[T])}}
}

// Insert 根据keyPath插入一系列节点，最后一个节点的值赋值或者更新为value
func (t *PrefixTree[T]) Insert(keyPath []T, value interface{}) {
	tr := t.root
	n := len(keyPath)
	i := 0
	for i < n {
		if nextNode, ok := tr.children[keyPath[i]]; ok {
			tr = nextNode
			i++
		} else {
			for i < n {
				newNode := NewPrefixTreeNode(keyPath[i], nil, false)
				tr.children[keyPath[i]] = newNode
				tr = newNode
				i++
			}
		}
	}
	tr.isEnd = true
	tr.value = value
}


// Search 搜索前缀树，返回最后一个节点绑定的值，如果没找到，则第二个返回值为false
func (t *PrefixTree[T]) Search(keyPath []T) (interface{}, bool) {
	tr := t.root
	n := len(keyPath)
	i := 0
	for i < n {
		if nextNode, ok := tr.children[keyPath[i]]; ok {
			tr = nextNode
			i++
		} else {
			return nil, false
		}
	}
	return tr.value, true
}


// GetShortestKey 返回前缀树所有最长前缀，树种不存在其他前缀是这些前缀的前缀
func (t *PrefixTree[T]) GetShortestKey() [][]T {
	res := make([][]T, 0)
	var dfs func(curNode *PrefixTreeNode[T], cur []T) 
	dfs = func(curNode *PrefixTreeNode[T], cur []T) {
		if curNode.isEnd {
			res = append(res, append([]T(nil), cur...))
			return
		}
		for _, child := range curNode.children {
			dfs(child, append(cur, child.key))
		}
	}
	dfs(t.root, make([]T, 0))
	return res
}