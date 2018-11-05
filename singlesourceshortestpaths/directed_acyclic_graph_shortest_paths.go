package singlesourceshortestpaths

import (
	"github.com/wangyoucao577/algorithms_practice/dfs"
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// DirectedAcyclicGraphShortestPaths implements the shortest paths search based on topo-sort
func DirectedAcyclicGraphShortestPaths(g *weightedgraph.WeightedGraph, s graph.NodeID) (*ShortestPaths, error) {

	sorted, err := dfs.NewTopologicalSort(g.Graph)
	if err != nil { // not a directed acyclic graph
		return nil, err
	}

	sp := &ShortestPaths{}
	sp.initializeSingleSource(g, s)

	sourceFound := false
	for _, u := range sorted {

		// start relax from the excepted source node
		if !sourceFound {
			if u != s {
				continue
			}
			sourceFound = true
		}

		g.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			w, _ := g.Weight(u, v)
			sp.relax(u, v, w)
		})
	}

	return sp, nil
}
