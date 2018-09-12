package maxflow

import (
	"strings"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/flownetwork"
)

func TestDrainageDitches(t *testing.T) {
	/*
		HDU1532 Drainage Ditches
		i.e. [POJ1273 Drainage Ditches](http://poj.org/problem?id=1273)

			Sample Input

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

	if f.Directed() != true {
		t.Errorf("expect directed graph, but got not")
	}

	maxFlow := FordFulkerson(f, false)
	if maxFlow != want {
		t.Errorf("[FordFulkerson] got maximum flow %v, want %v", maxFlow, want)
	}

	edmondsKarpMaxFlow := FordFulkerson(f, true)
	if edmondsKarpMaxFlow != want {
		t.Errorf("[EdmondsKarp] got maximum flow %v, want %v", edmondsKarpMaxFlow, want)
	}

	dinicMaxFlow := Dinic(f)
	if dinicMaxFlow != want {
		t.Errorf("[Dinic] got maximum flow %v, want %v", dinicMaxFlow, want)
	}

}

func TestFlowProblem(t *testing.T) {
	/*
		[HDU3549 Flow Problem](https://blog.csdn.net/sr_19930829/article/details/39525605)

		Sample Input
			2
			3 2
			1 2 1
			2 3 1
			3 3
			1 2 1
			2 3 1
			1 3 1
	*/

	tests := []struct {
		nodeCount        int
		edgeCount        int
		capacityContents string
		want             flownetwork.EdgeFlowUnit
	}{
		{3, 2, "1 2 1\n2 3 1", 1},
		{3, 3, "1 2 1\n2 3 1\n1 3 1", 2},
	}

	for _, v := range tests {
		f, err := flownetwork.ConstructFlowNetwork(v.nodeCount, v.edgeCount,
			strings.NewReader(v.capacityContents))
		if err != nil {
			t.Error(err)
		}

		if f.Directed() != true {
			t.Errorf("expect directed graph, but got not")
		}

		maxFlow := FordFulkerson(f, false)
		if maxFlow != v.want {
			t.Errorf("[FordFulkerson] got maximum flow %v, want %v", maxFlow, v.want)
		}

		edmondsKarpMaxFlow := FordFulkerson(f, true)
		if edmondsKarpMaxFlow != v.want {
			t.Errorf("[EdmondsKarp] got maximum flow %v, want %v", edmondsKarpMaxFlow, v.want)
		}

		dinicMaxFlow := Dinic(f)
		if dinicMaxFlow != v.want {
			t.Errorf("[Dinic] got maximum flow %v, want %v", dinicMaxFlow, v.want)
		}

	}

}
