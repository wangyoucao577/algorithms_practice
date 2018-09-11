package graphsample1

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

func TestAdjacencyListGraphSample(t *testing.T) {

	want := [][]graph.NodeID{
		[]graph.NodeID{1, 4},
		[]graph.NodeID{0, 5},
		[]graph.NodeID{3, 5, 6},
		[]graph.NodeID{2, 7},
		[]graph.NodeID{0},
		[]graph.NodeID{1, 2, 6},
		[]graph.NodeID{2, 5, 7},
		[]graph.NodeID{3, 6},
	}

	g := AdjacencyListGraphSample()

	if g.NodeCount() != nodeCount {
		t.Errorf("new adjacency list graph sample, got node count %d, want %d", g.NodeCount(), nodeCount)
	}

	wantEdgeCount := 0
	g.IterateAllNodes(func(u graph.NodeID) {
		i := graph.NodeID(0)
		adjNodesCount := 0
		wantEdgeCount += len(want[u])

		g.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			adjNodesCount++
			if adjNodesCount > len(want[u]) {
				t.Errorf("adjacency list graph sample, node %v, got %d adjacency nodes, but want %d adjacency nodes",
					u, adjNodesCount, len(want[u]))
			} else {
				if v != want[u][i] {
					t.Errorf("adjacency list graph sample, node %v, %dth adjacency node, %v, but want %v",
						u, i, v, want[u][i])
				}
				i++
			}
		})
	})

	if g.EdgeCount() != wantEdgeCount {
		t.Errorf("new adjacency list graph sample, got edge count %d, want %d", g.EdgeCount(), wantEdgeCount)
	}

}

// AdjacencyMatrixGraphSample1 adjacency Matrix based graph sample 1
/*
  For this undirected graph, we can only store half of the matrix to save storage if needed

	  r s t u v w x y

  r   0 1 0 0 1 0 0 0
  s   1 0 0 0 0 1 0 0
  t   0 0 0 1 0 1 1 0
  u   0 0 1 0 0 0 0 1
  v   1 0 0 0 0 0 0 0
  w   0 1 1 0 0 0 1 0
  x   0 0 1 0 0 1 0 1
  y   0 0 0 1 0 0 1 0
*/

func TestAdjacencyMatrixGraphSample(t *testing.T) {

	want := [][]bool{
		{false, true, false, false, true, false, false, false},
		{true, false, false, false, false, true, false, false},
		{false, false, false, true, false, true, true, false},
		{false, false, true, false, false, false, false, true},
		{true, false, false, false, false, false, false, false},
		{false, true, true, false, false, false, true, false},
		{false, false, true, false, false, true, false, true},
		{false, false, false, true, false, false, true, false},
	}

	g := AdjacencyMatrixGraphSample()

	if g.NodeCount() != nodeCount {
		t.Errorf("new adjacency matrix graph sample, got node count %d, want %d", g.NodeCount(), nodeCount)
	}

	wantEdgeCount := 0
	for _, s := range want {
		for _, b := range s {
			if b {
				wantEdgeCount++
			}
		}
	}
	if g.EdgeCount() != wantEdgeCount {
		t.Errorf("new adjacency matrix graph sample, got edge count %d, want %d", g.EdgeCount(), wantEdgeCount)
	}

	g.IterateAllNodes(func(u graph.NodeID) {
		g.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			if !want[u][v] {
				t.Errorf("adjacency matrix graph sample, got %v->%v, want %v !-> %v",
					u, v, u, v)
			}
		})
	})

}

func TestIDNameMap(t *testing.T) {

	//id <-> name map verify
	for i, v := range nodeConverter.orderedNodesName {
		gotNodeID, ok := nodeConverter.nodeNameToIDMap[v]
		if !ok || graph.NodeID(i) != gotNodeID {
			t.Errorf("NodeID of name %v not match, expect NodeID %v but got %v (exist in map %v)",
				v, graph.NodeID(i), gotNodeID, ok)
		}
	}
}
