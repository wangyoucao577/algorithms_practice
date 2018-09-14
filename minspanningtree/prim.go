package minspanningtree

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Prim algorithm calculate minimum spanning tree on the input undirected graph
func Prim(g weightedgraph.WeightedGraph) (MinSpanningTree, error) {

	mst := MinSpanningTree{[]graph.EdgeID{}, g}

	//TODO: implement

	return mst, nil
}
