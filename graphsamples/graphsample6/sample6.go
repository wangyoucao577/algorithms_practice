// Package graphsample6 defined a directed weighted graph comes from
//  "Introduction to Algorithms - Third Edition" 24.1 Bellman-Ford
package graphsample6

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

/* This sample directed weighted graph comes from
  "Introduction to Algorithms - Third Edition" 24.1 Bellman-Ford

	V = 5 (node count)
	E = 10 (edge count)
	define directed weighted graph G(V,E) as below:

	// too complex to draw as plain text...
*/

const (
	nodeCount     = 5
	directedGraph = true
)

type nodeIDNameConverter struct {
	orderedNodesName []string
	nodeNameToIDMap  map[string]graph.NodeID
}

// define fixed nodes order in the graph, then we use the `index` as nodeID for search,
// will be easier to implement by code.
// node name only for print.
var nodeConverter = nodeIDNameConverter{
	[]string{"s", "t", "x", "y", "z"},
	map[string]graph.NodeID{}, // will be inited during import
}

// initialization during package import
func init() {
	for i, v := range nodeConverter.orderedNodesName {
		nodeConverter.nodeNameToIDMap[v] = graph.NodeID(i)
	}
}

// IDToName convert NodeID to human readable name
func IDToName(i graph.NodeID) string {
	if i == graph.InvalidNodeID {
		return "InvalidNodeID"
	}
	return nodeConverter.orderedNodesName[i]
}

// NameToID convert node human readable name to NodeID
func NameToID(name string) graph.NodeID {
	return nodeConverter.nodeNameToIDMap[name]
}

// GraphSample return the weighted graph sample instance
func GraphSample() *weightedgraph.WeightedGraph {
	wg := weightedgraph.NewWeightedGraph(nodeCount, directedGraph, graph.NewAdjacencyListGraph)

	wg.AddEdge(0, 1, 6)
	wg.AddEdge(0, 3, 7)

	wg.AddEdge(1, 2, 5)
	wg.AddEdge(1, 3, 8)
	wg.AddEdge(1, 4, -4)

	wg.AddEdge(2, 1, -2)

	wg.AddEdge(3, 2, -3)
	wg.AddEdge(3, 4, 9)

	wg.AddEdge(4, 0, 2)
	wg.AddEdge(4, 2, 7)

	return wg
}
