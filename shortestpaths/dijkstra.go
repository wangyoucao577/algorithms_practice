package shortestpaths

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Dijkstra implements the dijkstra shortest path algorithm
func Dijkstra(g *weightedgraph.WeightedGraph, s graph.NodeID) *ShortestPaths {

	sp := &ShortestPaths{}
	sp.initializeSingleSource(g, s)

	q := []graph.NodeID{}
	g.IterateAllNodes(func(u graph.NodeID) {
		q = append(q, u)
	})

	for len(q) > 0 {
		u, newQ := sp.extractMin(q)
		q = newQ

		g.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			w, _ := g.Weight(u, v)
			sp.relax(u, v, w)
		})
	}

	return sp
}

func (sp *ShortestPaths) extractMin(q []graph.NodeID) (graph.NodeID, []graph.NodeID) {
	i := 0
	min := sp.nodesMap[q[i]].d

	for j := 1; j < len(q); j++ {
		if sp.nodesMap[q[j]].d < min {
			i = j
			min = sp.nodesMap[q[j]].d
		}
	}
	u := q[i]

	if i == 0 {
		q = q[1:]
	} else if i == len(q)-1 {
		q = q[:i]
	} else {
		q = append(q[:i], q[i+1:]...)
	}

	return u, q
}
