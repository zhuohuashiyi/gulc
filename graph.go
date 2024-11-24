package gulc


// Graph定义图的抽象接口
type Graph interface {
	PointCounts() int      // 返回图中点的个数
	EdgeCounts() int       // 返回图中边的个数
	IsDirect() bool        // 返回图是否是一个有向图
	HasCycle() (bool, []int)        // 判断图中是否存在环路,当不存在环路时，输出一个可行的输出队列
}


type MatrixGraph struct {
	matrix [][]int 
	edgeCount int
	isDirect bool
}


func NewMatrixGraph(edges [][]int, isDirect bool) Graph {
	return &MatrixGraph{}   // TODO 待实现
}


func (g *MatrixGraph) PointCounts() int {
	return len(g.matrix)
}


func (g *MatrixGraph) EdgeCounts() int {
	return g.edgeCount
}

func (g *MatrixGraph) IsDirect() bool {
	return g.isDirect
}


func (g *MatrixGraph) HasCycle() (bool, []int) {
	return false, nil
}

type AdjTableNode struct {
	Point int 
	Cost  int
}

type AdjTable struct {
	tab []*LinkedList[AdjTableNode]
	edgeCount int
	isDirect bool
	edges [][]int
}


func NewAdjTable(edges [][]int, isDirect bool, n int) Graph {
	tab := make([]*LinkedList[AdjTableNode], n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		c := 1
		if len(edge) > 2 {   // 兼容有权图以及无权图
			c = edge[2]
		}
		var insertNode func(a, b, c int) = func(a, b, c int) {
			if tab[a] == nil {
				tab[a] = NewLinkedList([]AdjTableNode{{b, c}})
			} else {
				tab[a].InsertNode(AdjTableNode{b, c})
			}
		}
		insertNode(a, b, c)
		if !isDirect {  // 兼容无向图
			insertNode(b, a, c)
		}
	}
	edgeCount := len(edges)
	if !isDirect {
		edgeCount *= 2
	}
	return &AdjTable{tab: tab, edgeCount: edgeCount, isDirect: isDirect, edges: edges}
}


func (g *AdjTable) PointCounts() int {
	return len(g.tab)
}


func (g *AdjTable) EdgeCounts() int {
	return g.edgeCount
}

func (g *AdjTable) IsDirect() bool {
	return g.isDirect
}

func (g *AdjTable) HasCycle() (bool, []int) {
	if g.isDirect {
		return g.hasCycleDFS()
	}
	return g.hasCycleUnionSet(), nil
}


func (g *AdjTable) topoSort() (bool, []int) {
	n := g.PointCounts()
	// 首先统计各个节点的入度
	inDegrees := make([]int, n)
	for i := 0; i < n; i++ {
		if g.tab[i] == nil {
			continue
		}
		node := g.tab[i].Head
		for node != nil {
			p, _ := node.Val.Point, node.Val.Cost
			inDegrees[p]++
			node = node.Next
		}
	}
	// 初始化队列，收集入度为零的节点
	deque := NewList[int]()
	for i := 0; i < n; i++ {
		if inDegrees[i] == 0 {
			deque.AppendTail(i)
		}
	}

	// 开始拓扑排序
	topoSorted := make([]int, 0)
	for !deque.IsEmpty() {
		node := deque.RemoveHead()
		topoSorted = append(topoSorted, node)
		if g.tab[node] == nil {
			continue
		}
		head := g.tab[node].Head
		for head != nil {
			p, _ := head.Val.Point, head.Val.Cost
			inDegrees[p]--
			if inDegrees[p] == 0 {
				deque.AppendTail(p)
			}
			head = head.Next
		}
	}
	return len(topoSorted) != n, topoSorted
}

func (g *AdjTable) hasCycleDFS() (bool, []int) {
	n := g.PointCounts()
	colors := make([]int, n)   // 三色标记
	hasCycle := false
	s := NewStack[int]()

	var dfs func(cur, parent int)
	dfs = func(cur, parent int) {
		if colors[cur] != 0 {
			return
		}
		colors[cur] = 1   // 灰色
		if g.tab[cur] == nil {
			colors[cur] = 2   // 黑色
			s.Push(cur)
			return
		}
		head := g.tab[cur].Head
		for head != nil {
			p, _ := head.Val.Point, head.Val.Cost
			if (g.isDirect|| p != parent) && colors[p] == 1 {
				hasCycle = true
				return
			}
			dfs(p, cur)
			if hasCycle {
				return
			}
			head = head.Next
		}
		colors[cur] = 2   // 黑色
		s.Push(cur)
	}

	for i := 0; i < n; i++ {
		if colors[i] == 0 {  // 未被访问
			dfs(i, -1)
			if hasCycle {
				return hasCycle, nil
			}
		}
	}
	topoSorted := make([]int, 0, n)
	for !s.IsEmpty() {
		topoSorted = append(topoSorted, s.Pop())
	}
	return hasCycle, topoSorted
}


func (g *AdjTable) hasCycleUnionSet() bool {
	n := g.PointCounts()
	unionSet := NewUnionSet(n)
	for _, edge := range g.edges {
		a, b := edge[0], edge[1]
		if unionSet.InSame(a, b) {
			return true
		}
		unionSet.Merge(a, b)
	}
	return false
}