package dfs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/graphsample2"
)

func (d *Dfs) printEdges(g graph.Graph) {
	g.IterateAllNodes(func(k graph.NodeID) {
		g.IterateAdjacencyNodes(k, func(v graph.NodeID) {
			edge := graph.EdgeID{From: k, To: v}
			fmt.Printf("edge %v, type %v\n", edge, d.edgesAttr[edge].t)
		})
	})
	fmt.Println()
}
func TestDfsOnGraphSample2(t *testing.T) {

	var tests = []struct {
		m    implementMethod
		want []dfsTree
	}{
		{Recurse, []dfsTree{dfsTree{graph.NodeID(0)}, dfsTree{graph.NodeID(2)}}},
		{StackBased, []dfsTree{dfsTree{graph.NodeID(0)}, dfsTree{graph.NodeID(2)}}},
	}

	var latestDfs *Dfs = nil

	for _, v := range tests {
		// adjacency list based graph
		dList, err := NewDfs(graphsample2.AdjacencyListGraphSample(), v.m)
		if err != nil {
			t.Errorf("DFS on adjacency list based graph failed, err %v", err)
		}
		if !reflect.DeepEqual(dList.forest, v.want) {
			t.Errorf("DFS on graphsample2.AdjacencyListGraphSample, got forest %v, want %v", dList.forest, v.want)
		}
		//dList.printEdges(graphsample2.AdjacencyListGraphSample())
		if latestDfs != nil {
			if !reflect.DeepEqual(latestDfs, dList) {
				t.Errorf("DFS structure not equal to latest one, got forest %v, want %v", dList, latestDfs)
			}
		} else {
			latestDfs = dList
		}

		// adjacency matrix based graph
		dMatrix, err := NewDfs(graphsample2.AdjacencyMatrixGraphSample(), v.m)
		if err != nil {
			t.Errorf("DFS on adjacency matrix based graph failed, err %v", err)
		}
		if !reflect.DeepEqual(dMatrix.forest, v.want) {
			t.Errorf("DFS on graphsample2.AdjacencyMatrixGraphSample, got forest %v, want %v", dList.forest, v.want)
		}
		//dMatrix.printEdges(graphsample2.AdjacencyMatrixGraphSample())
		if latestDfs != nil {
			if !reflect.DeepEqual(latestDfs, dMatrix) {
				t.Errorf("DFS structure not equal to latest one, got forest %v, want %v", dMatrix, latestDfs)
			}
		} else {
			latestDfs = dMatrix
		}

	}

}
