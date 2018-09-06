package flownetwork

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

	f, err := ConstructFlowNetwork(nodeCount, edgeCount, strings.NewReader(inputScanContents))
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(g)
	if f.Graph().NodeCount() != nodeCount {
		t.Errorf("node count got %d, want %d", f.Graph().NodeCount(), nodeCount)
	}
	if f.Graph().EdgeCount() != edgeCount {
		t.Errorf("edge count got %d, want %d", f.Graph().EdgeCount(), edgeCount)
	}
	if len(f.capacities) != edgeCount {
		t.Errorf("edge attr count got %d, want %d", len(f.capacities), edgeCount)
	}

	f.Graph().IterateAllNodes(func(u graph.NodeID) {
		f.Graph().IterateAdjacencyNodes(u, func(v graph.NodeID) {
			// edge u->v must exist
			edge := graph.EdgeID{From: u, To: v}
			if _, ok := f.capacities[edge]; !ok {
				t.Errorf("edge %v no capacity exist", edge)
			}

			// edge v->u must NOT exist
			if _, ok := f.capacities[edge.Reverse()]; ok {
				t.Errorf("edge %v capacity exist, but expect not exist", edge.Reverse())
			}
		})
	})
}
