package shortestpaths

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// BellmanFord implements bellman-ford algorithm for single-source shortest path problem
func BellmanFord(g *weightedgraph.WeightedGraph, s graph.NodeID) (bool, *ShortestPaths) {

	sp := &ShortestPaths{}
	sp.initializeSingleSource(g, s)

	v := g.NodeCount()
	for i := 0; i < v-1; i++ { // iterate v-1 times
		g.IterateEdges(func(e graph.EdgeID) {
			w, _ := g.Weight(e.From, e.To)
			sp.relax(e.From, e.To, w)
		})
	}

	negativeCycle := false
	g.IterateEdges(func(e graph.EdgeID) {
		w, _ := g.Weight(e.From, e.To)
		if sp.relaxable(e.From, e.To, w) {

			//TODO: improvable - if negative cycle found, then can return directly.
			negativeCycle = true
		}
	})

	if negativeCycle {
		return false, nil
	}
	return true, sp
}
