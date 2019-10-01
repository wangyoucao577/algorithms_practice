package dfs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/graphsamples/graphsample2"
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

func (d *Dfs) ValidateEdges(g graph.Graph, t *testing.T) {

	g.IterateAllNodes(func(k graph.NodeID) {
		g.IterateAdjacencyNodes(k, func(v graph.NodeID) {
			edge := graph.EdgeID{From: k, To: v}
			kD := d.nodesAttr[k].timestampD
			kF := d.nodesAttr[k].timestampF
			vD := d.nodesAttr[v].timestampD
			vF := d.nodesAttr[v].timestampF

			switch d.edgesAttr[edge].t {
			case branch:
				// expect `k.D < v.D < v.F < k.F`
				if !((kD < vD) && (vD < vF) && (vF < kF)) {
					t.Errorf("edge %v type %v, expect %d < %d < %d < %d", edge, branch, kD, vD, vF, kF)
				}
			case forward:
				// expect `k.D < v.D < v.F < k.F`
				if !((kD < vD) && (vD < vF) && (vF < kF)) {
					t.Errorf("edge %v type %v, expect %d < %d < %d < %d", edge, forward, kD, vD, vF, kF)
				}
			case backward:
				// expect `v.D <= k.D < k.F <= v.F`
				if !((vD <= kD) && (kD < kF) && (kF <= vF)) {
					t.Errorf("edge %v type %v, expect %d <= %d < %d <= %d", edge, backward, vD, kD, kF, vF)
				}
			case cross:
				// expect `v.D < v.F < k.D < k.F`
				if !((vD < vF) && (vF < kD) && (kD < kF)) {
					t.Errorf("edge %v type %v, expect %d <= %d < %d <= %d", edge, cross, vD, vF, kD, kF)
				}
			}
		})
	})

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
		dList, err := NewDfsForest(graphsample2.AdjacencyListGraphSample(), v.m)
		if err != nil {
			t.Errorf("DFS on adjacency list based graph failed, err %v", err)
		}
		if !reflect.DeepEqual(dList.forest, v.want) {
			t.Errorf("DFS on graphsample2.AdjacencyListGraphSample, got forest %v, want %v", dList.forest, v.want)
		}
		dList.ValidateEdges(graphsample2.AdjacencyListGraphSample(), t)
		//dList.printEdges(graphsample2.AdjacencyListGraphSample())
		if latestDfs != nil {
			if !reflect.DeepEqual(latestDfs, dList) {
				t.Errorf("DFS structure not equal to latest one, got forest %v, want %v", dList, latestDfs)
			}
		} else {
			latestDfs = dList
		}

		// adjacency matrix based graph
		dMatrix, err := NewDfsForest(graphsample2.AdjacencyMatrixGraphSample(), v.m)
		if err != nil {
			t.Errorf("DFS on adjacency matrix based graph failed, err %v", err)
		}
		if !reflect.DeepEqual(dMatrix.forest, v.want) {
			t.Errorf("DFS on graphsample2.AdjacencyMatrixGraphSample, got forest %v, want %v", dList.forest, v.want)
		}
		dMatrix.ValidateEdges(graphsample2.AdjacencyMatrixGraphSample(), t)
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
