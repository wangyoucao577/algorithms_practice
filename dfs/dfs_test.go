package dfs

import (
	"reflect"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/graphsample2"
)

func TestDfsOnGraphSample2(t *testing.T) {

	var tests = []struct {
		m    implementMethod
		want []dfsTree
	}{
		{Recurse, []dfsTree{dfsTree{graph.NodeID(0)}, dfsTree{graph.NodeID(2)}}},
		{StackBased, []dfsTree{dfsTree{graph.NodeID(0)}, dfsTree{graph.NodeID(2)}}},
	}

	for _, v := range tests {
		// adjacency list based graph
		dList, err := NewDfs(graphsample2.AdjacencyListGraphSample, v.m)
		if err != nil {
			t.Errorf("DFS on adjacency list based graph failed, err %v", err)
		}
		if !reflect.DeepEqual(dList.forest, v.want) {
			t.Errorf("DFS on graphsample2.AdjacencyListGraphSample, got forest %v, want %v", dList.forest, v.want)
		}

		// adjacency matrix based graph
		dMatrix, err := NewDfs(graphsample2.AdjacencyMatrixGraphSample, v.m)
		if err != nil {
			t.Errorf("DFS on adjacency matrix based graph failed, err %v", err)
		}
		if !reflect.DeepEqual(dMatrix.forest, v.want) {
			t.Errorf("DFS on graphsample2.AdjacencyMatrixGraphSample, got forest %v, want %v", dList.forest, v.want)
		}
	}

}
