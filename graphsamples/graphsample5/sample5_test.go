package graphsample5

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

func TestUndirectedWeightedGraphSample(t *testing.T) {

	g := GraphSample()

	if g.Directed() != directedGraph {
		t.Errorf("new weighted graph sample, got directed %v, want %v", g.Directed(), directedGraph)
	}

	if g.NodeCount() != nodeCount {
		t.Errorf("new weighted graph sample, got node count %d, want %d", g.NodeCount(), nodeCount)
	}

	wantEdgeCount := 14 // undirected edge
	if g.EdgeCount() != wantEdgeCount {
		t.Errorf("new weighted graph sample, got edge count %d, want %d", g.EdgeCount(), wantEdgeCount)
	}

	var iteratedEdgeCount int
	g.IterateEdges(func(edge graph.EdgeID) {
		iteratedEdgeCount++
	})
	if iteratedEdgeCount != wantEdgeCount {
		t.Errorf("new adjacency list graph sample, got iterated edge count %d, want %d", iteratedEdgeCount, wantEdgeCount)
	}

	if !g.Validate() {
		t.Errorf("new weighted graph sample is not a valid weighted graph")
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
