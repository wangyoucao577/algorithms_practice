// Package shortestpaths implmenets single-source shortest paths algorithms
package shortestpaths

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

const (
	infinitelyWeight weightedgraph.Weight = weightedgraph.Weight((^uint(0)) >> 1)
)

type nodeAttr struct {
	d      weightedgraph.Weight // weight from source to current node
	parent graph.NodeID         // remember parent node, InvalidNodeID means no parent
}

// ShortestPaths defined a structure to deal with single-source shortest paths problem
type ShortestPaths struct {
	g *weightedgraph.WeightedGraph // directed weighted graph
	s graph.NodeID                 // source node

	nodesMap map[graph.NodeID]*nodeAttr // extend information during calculation
}

func (sp *ShortestPaths) initializeSingleSource(g *weightedgraph.WeightedGraph, s graph.NodeID) {
	sp.g = g
	sp.s = s
	sp.nodesMap = map[graph.NodeID]*nodeAttr{}
	g.IterateAllNodes(func(u graph.NodeID) {
		sp.nodesMap[u] = &nodeAttr{infinitelyWeight, graph.InvalidNodeID}
	})

	sp.nodesMap[s].d = 0
}

// relax return the new weight if updated
func (sp *ShortestPaths) relax(u, v graph.NodeID, w weightedgraph.Weight) weightedgraph.Weight {
	if sp.nodesMap[v].d > sp.nodesMap[u].d+w { // can handle infinitelyWeight since we set it's too max
		sp.nodesMap[v].d = sp.nodesMap[u].d + w
		sp.nodesMap[v].parent = u

		return sp.nodesMap[v].d
	}
	return 0
}

func (sp ShortestPaths) relaxable(u, v graph.NodeID, w weightedgraph.Weight) bool {
	if sp.nodesMap[v].d > sp.nodesMap[u].d+w { // can handle infinitelyWeight since we set it's too max
		return true
	}
	return false
}

// RetrievePath to retrieve the shortest path and it's weight from s to v
func (sp ShortestPaths) RetrievePath(v graph.NodeID) (graph.Path, weightedgraph.Weight) {

	path := graph.Path{}
	weight := sp.nodesMap[v].d

	if weight == infinitelyWeight { // no valid path
		return path, weight
	}

	path = append(path, v)
	for v != sp.s {
		path = append(path, sp.nodesMap[v].parent)
		v = sp.nodesMap[v].parent
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, weight
}
