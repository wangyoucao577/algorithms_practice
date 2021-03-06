package graphsample4

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

func TestAdjacencyListGraphSample(t *testing.T) {

	want := [][]graph.NodeID{
		[]graph.NodeID{1},
		[]graph.NodeID{2, 4, 5},
		[]graph.NodeID{3, 6},
		[]graph.NodeID{2, 7},
		[]graph.NodeID{0, 5},
		[]graph.NodeID{6},
		[]graph.NodeID{5, 7},
		[]graph.NodeID{7},
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
