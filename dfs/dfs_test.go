package dfs

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/graphsample2"
)

func forestEqual(p []dfsTree, q []dfsTree) bool {
	if len(p) != len(q) {
		return false
	}

	if (p == nil) != (q == nil) {
		return false
	}

	for i := range p {
		if p[i] != q[i] {
			return false
		}
	}

	return true
}

func TestDfsOnGraphSample2(t *testing.T) {

	want := []dfsTree{dfsTree{graph.NodeID(0)}, dfsTree{graph.NodeID(2)}}

	// adjacency list based graph
	dList, err := NewDfs(graphsample2.AdjacencyListGraphSample)
	if err != nil {
		t.Errorf("DFS on adjacency list based graph failed, err %v", err)
	}
	if !forestEqual(dList.Forest, want) {
		t.Errorf("DFS on graphsample2.AdjacencyListGraphSample, got forest %v, want %v", dList.Forest, want)
	}

	// adjacency matrix based graph
	dMatrix, err := NewDfs(graphsample2.AdjacencyMatrixGraphSample)
	if err != nil {
		t.Errorf("DFS on adjacency matrix based graph failed, err %v", err)
	}
	if !forestEqual(dMatrix.Forest, want) {
		t.Errorf("DFS on graphsample2.AdjacencyMatrixGraphSample, got forest %v, want %v", dList.Forest, want)
	}

}
