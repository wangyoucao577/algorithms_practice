package dfs

import (
	"reflect"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/graphsample3"
)

func TestTopologicalSort(t *testing.T) {

	alternativesWant := [][]graph.NodeID{
		{1, 0, 2, 3, 8, 4, 5, 6, 7},
		{8, 4, 6, 1, 0, 2, 5, 7, 3},
	}

	sorted, err := NewTopologicalSort(graphsample3.AdjacencyListGraphSample())
	if err != nil {
		t.Errorf("DFS based Topological Sort on adjacency list based graph failed, err %v", err)
	}

	matchedWant := false
	for _, want := range alternativesWant {
		if reflect.DeepEqual(sorted, want) {
			matchedWant = true
			break
		}
	}
	if !matchedWant {
		t.Errorf("Topological Sorted not match any want, got %v (human readable: %v)",
			sorted, graphsample3.IDSToNames(sorted))
	}
}
