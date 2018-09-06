package networkflowgraph

import (
	"strings"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

func TestConstructDrainageDitchesSampleGraph(t *testing.T) {

	/*
		Sample Input [POJ1273](http://poj.org/problem?id=1273)

			5 4
			1 2 40
			1 4 20
			2 4 20
			2 3 30
			3 4 10
	*/
	edgeCount, nodeCount := 5, 4
	inputScanContents := "1 2 40\n1 4 20\n2 4 20\n2 3 30\n3 4 10"

	g, err := ConstructNetworkFlowGraph(nodeCount, edgeCount, strings.NewReader(inputScanContents))
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(g)
	if g.AdjGraph.NodeCount() != nodeCount {
		t.Errorf("node count got %d, want %d", g.AdjGraph.NodeCount(), nodeCount)
	}
	if g.AdjGraph.EdgeCount() != edgeCount {
		t.Errorf("edge count got %d, want %d", g.AdjGraph.EdgeCount(), edgeCount)
	}
	if len(g.capacities) != edgeCount {
		t.Errorf("edge attr count got %d, want %d", len(g.capacities), edgeCount)
	}

	g.AdjGraph.IterateAllNodes(func(u graph.NodeID) {
		g.AdjGraph.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			// edge u->v must exist
			edge := graph.EdgeID{From: u, To: v}
			if _, ok := g.capacities[edge]; !ok {
				t.Errorf("edge %v no capacity exist", edge)
			}

			// edge v->u must NOT exist
			if _, ok := g.capacities[edge.Reverse()]; ok {
				t.Errorf("edge %v capacity exist, but expect not exist", edge.Reverse())
			}
		})
	})
}
