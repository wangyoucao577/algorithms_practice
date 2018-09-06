package maxflow

import (
	"strings"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/flownetwork"
)

func TestDrainageDitches(t *testing.T) {
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
	want := flownetwork.EdgeFlowUnit(50)

	f, err := flownetwork.ConstructFlowNetwork(nodeCount, edgeCount, strings.NewReader(inputScanContents))
	if err != nil {
		t.Error(err)
	}

	maxFlow := FordFulkerson(f, false)
	if maxFlow != want {
		t.Errorf("[FordFulkerson] got maximum flow %v, want %v", maxFlow, want)
	}

	edmondsKarpMaxFlow := FordFulkerson(f, true)
	if edmondsKarpMaxFlow != want {
		t.Errorf("[EdmondsKarp] got maximum flow %v, want %v", edmondsKarpMaxFlow, want)
	}

}
