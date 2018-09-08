package graphsample3

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

func TestAdjacencyListGraphSample(t *testing.T) {

	want := [][]graph.NodeID{
		[]graph.NodeID{2, 3},
		[]graph.NodeID{3},
		[]graph.NodeID{3, 5},
		[]graph.NodeID{},
		[]graph.NodeID{5, 6},
		[]graph.NodeID{7},
		[]graph.NodeID{7},
		[]graph.NodeID{},
		[]graph.NodeID{},
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
