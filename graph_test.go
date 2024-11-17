package gulc


import (
	"testing"
)


func TestTopoSort(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}}
	g1 := NewAdjTable(edges, true, 4)
	hasCycle, path := g1.HasCycle()
	t.Logf("get result: hasCycle=%v, path=%+v", hasCycle, path)

	edges = [][]int{{0, 1}, {1, 2}, {2, 0}}
	g2 := NewAdjTable(edges, true, 3)
	hasCycle, path = g2.HasCycle()
	t.Logf("get result: hasCycle=%v, path=%+v", hasCycle, path)
}