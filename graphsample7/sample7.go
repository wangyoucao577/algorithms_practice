// Package graphsample7 defined a directed acyclic weighted graph comes from
//  "Introduction to Algorithms - Third Edition" 24.2 Directed Acyclic Graph Shortest Paths
package graphsample7

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

/* This sample directed weighted graph comes from
  "Introduction to Algorithms - Third Edition" 24.2 Directed Acyclic Graph Shortest Paths

	V = 6 (node count)
	E = 10 (edge count)
	define directed weighted graph G(V,E) as below:

	  5    2    7    -1   -2
	r -> s -> t -> x -> y -> z
   (0)  (1)  (2)  (3)  (4)  (5)

		3          4
	r ------> t ------> y

	         6          1
		 s ------> x ------> z

			          2
		      t -----------> z
*/

const (
	nodeCount     = 6
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
	[]string{"r", "s", "t", "x", "y", "z"},
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

	wg.AddEdge(0, 1, 5)
	wg.AddEdge(0, 2, 3)

	wg.AddEdge(1, 2, 2)
	wg.AddEdge(1, 3, 6)

	wg.AddEdge(2, 3, 7)
	wg.AddEdge(2, 4, 4)
	wg.AddEdge(2, 5, 2)

	wg.AddEdge(3, 4, -1)
	wg.AddEdge(3, 5, 1)

	wg.AddEdge(4, 5, -2)

	return wg
}
