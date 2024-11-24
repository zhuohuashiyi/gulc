package gulc


import (
	"testing"
)


func TestTopoSort(t *testing.T) {
	// 以下两个测试用例测试有向图有环无环两种情况
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}}
	g1 := NewAdjTable(edges, true, 4)
	hasCycle, path := g1.HasCycle()
	t.Logf("get result: hasCycle=%v, path=%+v", hasCycle, path)

	edges = [][]int{{0, 1}, {1, 2}, {2, 0}}
	g2 := NewAdjTable(edges, true, 3)
	hasCycle, path = g2.HasCycle()
	t.Logf("get result: hasCycle=%v, path=%+v", hasCycle, path)

	// 测试无向图有环无环
	edges = [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}}
	g3 := NewAdjTable(edges, false, 4)
	hasCycle, path = g3.HasCycle()
	t.Logf("get result: hasCycle=%v, path=%+v", hasCycle, path)

	edges = [][]int{{0, 1}, {1, 2}}
	g4 := NewAdjTable(edges, false, 3)
	hasCycle, path = g4.HasCycle()
	t.Logf("get result: hasCycle=%v, path=%+v", hasCycle, path)
}