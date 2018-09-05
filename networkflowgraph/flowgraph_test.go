package networkflowgraph

import (
	"strings"
	"testing"
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
	if g.baseGraph.NodeCount() != nodeCount {
		t.Errorf("node count got %d, want %d", g.baseGraph.NodeCount(), nodeCount)
	}
	if g.baseGraph.EdgeCount() != edgeCount {
		t.Errorf("edge count got %d, want %d", g.baseGraph.EdgeCount(), edgeCount)
	}
	if len(g.edgesAttr) != edgeCount {
		t.Errorf("edge attr count got %d, want %d", len(g.edgesAttr), edgeCount)
	}

}
