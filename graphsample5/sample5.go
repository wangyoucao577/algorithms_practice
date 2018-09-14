// Package graphsample5 defined a undirected weighted graph comes from
//  "Introduction to Algorithms - Third Edition" 23.2 Kruskal & Prim
package graphsample5

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

/* This sample undirected weighted graph comes from
  "Introduction to Algorithms - Third Edition" 23.2 Kruskal & Prim

	V = 9 (node count)
	E = 14 (edge count)
	define undirected weighted graph G(V,E) as below:

             8      7
		b(1) - c(2) - d(3)
	 4/  |    2/   \    |  \9
  a(0) 11|  i(8)   4\   |14  e(4)
     8\  | 7/ \6     \  |  /10
		h(7) - g(6) - f(5)
			 1      2
*/

const (
	nodeCount     = 9
	directedGraph = false
)

type nodeIDNameConverter struct {
	orderedNodesName []string
	nodeNameToIDMap  map[string]graph.NodeID
}

// define fixed nodes order in the graph, then we use the `index` as nodeID for search,
// will be easier to implement by code.
// node name only for print.
var nodeConverter = nodeIDNameConverter{
	[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
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

	wg.AddEdge(0, 1, 4)
	wg.AddEdge(0, 7, 8)

	wg.AddEdge(1, 2, 8)
	wg.AddEdge(1, 7, 11)

	wg.AddEdge(2, 3, 7)
	wg.AddEdge(2, 5, 4)
	wg.AddEdge(2, 8, 2)

	wg.AddEdge(3, 4, 9)
	wg.AddEdge(3, 5, 14)

	wg.AddEdge(4, 5, 10)

	wg.AddEdge(5, 6, 2)

	wg.AddEdge(6, 7, 1)
	wg.AddEdge(6, 8, 6)

	wg.AddEdge(7, 8, 7)

	return wg
}
