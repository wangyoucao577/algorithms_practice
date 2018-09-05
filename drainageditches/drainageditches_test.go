package drainageditches

import (
	"strings"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/networkflowgraph"
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
	want := 50

	g, err := networkflowgraph.ConstructNetworkFlowGraph(nodeCount, edgeCount, strings.NewReader(inputScanContents))
	if err != nil {
		t.Error(err)
	}

	maxFlow := FordFulkerson(g)
	if maxFlow != want {
		t.Errorf("got maximum flow %v, want %v", maxFlow, want)
	}
}
