// Package shortestpath implmenets single-source shortest paths algorithms
package shortestpath

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

type shortestPath struct {
	nodesMap map[graph.NodeID]*nodeAttr
}

func (sp *shortestPath) initializeSingleSource(g *weightedgraph.WeightedGraph, s graph.NodeID) {

	sp.nodesMap = map[graph.NodeID]*nodeAttr{}
	g.IterateAllNodes(func(u graph.NodeID) {
		sp.nodesMap[u] = &nodeAttr{infinitelyWeight, graph.InvalidNodeID}
	})

	sp.nodesMap[s].d = 0
}

func (sp *shortestPath) relax(u, v graph.NodeID, w weightedgraph.Weight) {
	if sp.nodesMap[v].d > sp.nodesMap[u].d+w { // can handle infinitelyWeight since we set it's too max
		sp.nodesMap[v].d = sp.nodesMap[u].d + w
		sp.nodesMap[v].parent = u
	}
}

func (sp shortestPath) relaxable(u, v graph.NodeID, w weightedgraph.Weight) bool {
	if sp.nodesMap[v].d > sp.nodesMap[u].d+w { // can handle infinitelyWeight since we set it's too max
		return true
	}
	return false
}
