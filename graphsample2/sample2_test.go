package graphsample2

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

func TestAdjacencyListGraphSample(t *testing.T) {

	want := [][]graph.NodeID{
		[]graph.NodeID{1, 3},
		[]graph.NodeID{4},
		[]graph.NodeID{4, 5},
		[]graph.NodeID{1},
		[]graph.NodeID{3},
		[]graph.NodeID{5},
	}

	g := AdjacencyListGraphSample()

	if g.Directed() != directedGraph {
		t.Errorf("new adjacency list graph sample, got directed %v, want %v", g.Directed(), directedGraph)
	}

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

	var iteratedEdgeCount int
	g.IterateEdges(func(edge graph.EdgeID) {
		iteratedEdgeCount++
	})
	if iteratedEdgeCount != wantEdgeCount {
		t.Errorf("new adjacency list graph sample, got iterated edge count %d, want %d", iteratedEdgeCount, wantEdgeCount)
	}
}

// AdjacencyMatrixGraphSample adjacency Matrix based graph sample 2
/*
    For directed graph, the matrix will be asymmetric.

	  u v w x y z

  u   0 1 0 1 0 0
  v   0 0 0 0 1 0
  w   0 0 0 0 1 1
  x   0 1 0 0 0 0
  y   0 0 0 1 0 0
  z   0 0 0 0 0 1
*/

func TestAdjacencyMatrixGraphSample(t *testing.T) {

	want := [][]bool{
		{false, true, false, true, false, false},
		{false, false, false, false, true, false},
		{false, false, false, false, true, true},
		{false, true, false, false, false, false},
		{false, false, false, true, false, false},
		{false, false, false, false, false, true},
	}

	g := AdjacencyMatrixGraphSample()

	if g.Directed() != directedGraph {
		t.Errorf("new adjacency matrix graph sample, got directed %v, want %v", g.Directed(), directedGraph)
	}

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

	var iteratedEdgeCount int
	g.IterateEdges(func(edge graph.EdgeID) {
		iteratedEdgeCount++
	})
	if iteratedEdgeCount != wantEdgeCount {
		t.Errorf("new adjacency matrix graph sample, got iterated edge count %d, want %d", iteratedEdgeCount, wantEdgeCount)
	}
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
