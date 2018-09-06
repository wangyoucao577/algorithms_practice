package bfs

import (
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/graphsample1"
)

func TestBfsOnGraphSample1(t *testing.T) {
	source := graph.NodeID(1)

	type queryWant struct {
		Depth        int
		ShortestPath graph.Path
	}
	var tests = []struct {
		target graph.NodeID
		want   queryWant
	}{
		{source, queryWant{0, graph.Path{source}}},
		{graph.NodeID(0), queryWant{1, graph.Path{source, graph.NodeID(0)}}},
		{graph.NodeID(4), queryWant{2, graph.Path{source, graph.NodeID(0), graph.NodeID(4)}}},
		{graph.NodeID(5), queryWant{1, graph.Path{source, graph.NodeID(5)}}},
		{graph.NodeID(6), queryWant{2, graph.Path{source, graph.NodeID(5), graph.NodeID(6)}}},
		{graph.NodeID(2), queryWant{2, graph.Path{source, graph.NodeID(5), graph.NodeID(2)}}},
		{graph.NodeID(3), queryWant{3, graph.Path{source, graph.NodeID(5), graph.NodeID(2), graph.NodeID(3)}}},
		{graph.NodeID(7), queryWant{3, graph.Path{source, graph.NodeID(5), graph.NodeID(6), graph.NodeID(7)}}},
	}

	// adjacency list based graph
	bList, err := NewBfs(graphsample1.AdjacencyListGraphSample(), source, nil)
	if err != nil {
		t.Errorf("BFS on adjacency list based graph failed, source %v", source)
	}
	for _, v := range tests {
		depth, path, _ := bList.Query(v.target)
		if depth != v.want.Depth {
			t.Errorf("Query for %v, got depth %d, want %d", v.target, depth, v.want.Depth)
		}

		if !path.Equal(v.want.ShortestPath) {
			t.Errorf("Query for %v, got shortest path %v, want %v", v.target, path, v.want.ShortestPath)
		}
	}

	// adjacency matrix based graph
	bMatrix, err := NewBfs(graphsample1.AdjacencyMatrixGraphSample(), source, nil)
	if err != nil {
		t.Errorf("BFS on adjacency Matrix based graph failed, source %v", source)
	}
	for _, v := range tests {
		depth, path, _ := bMatrix.Query(v.target)
		if depth != v.want.Depth {
			t.Errorf("Query for %v, got depth %d, want %d", v.target, depth, v.want.Depth)
		}

		if !path.Equal(v.want.ShortestPath) {
			t.Errorf("Query for %v, got shortest path %v, want %v", v.target, path, v.want.ShortestPath)
		}
	}

}
