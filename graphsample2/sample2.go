// Package graphsample2 defined a directed graph comes from
//  "Introduction to Algorithms - Third Edition" 22.1
package graphsample2

import "github.com/wangyoucao577/algorithms_practice/graph"

/* This sample directed graph comes from
  "Introduction to Algorithms - Third Edition" 22.1

	V = 6 (node count)
	E = 8 (edge count)
	define directed graph G(V,E) as below:

	u(0) -> v(1)    w(2)
	  ↓  ↗   ↓   ↙   ↓
	x(3) <- y(4)    z(5) ->
					  ↑    ↓
						<-

    NOTE: node `z` has a spin edge pointer to itself.
*/

type nodeIDNameConverter struct {
	orderedNodesName []string
	nodeNameToIDMap  map[string]graph.NodeID
}

// define fixed nodes order in the graph, then we use the `index` as nodeID for search,
// will be easier to implement by code.
// node name only for print.
var nodeConverter = nodeIDNameConverter{
	[]string{"u", "v", "w", "x", "y", "z"},
	map[string]graph.NodeID{"u": 0, "v": 1, "w": 2, "x": 3, "y": 4, "z": 5},
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

// AdjacencyListGraphSample adjacency list based graph sample 2
var adjacencyListGraphSample = graph.AdjacencyListGraph{
	[]graph.NodeID{1, 3},
	[]graph.NodeID{4},
	[]graph.NodeID{4, 5},
	[]graph.NodeID{1},
	[]graph.NodeID{3},
	[]graph.NodeID{5},
}

// AdjacencyMatrixGraphSample adjacency Matrix based graph sample 2
/*
    For directed graph, the matrix will be asymmetric.

	  u v w x y z

  u   0 1 0 1 0 0
  v   0 0 0 0 1 0
  w   0 0 0 0 1 1
  x   0 1 0 0 0 0
  y   0 0 0 1 0 0
  z   0 0 0 0 0 1
*/
var adjacencyMatrixGraphSample = graph.AdjacencyMatrixGraph{
	{false, true, false, true, false, false},
	{false, false, false, false, true, false},
	{false, false, false, false, true, true},
	{false, true, false, false, false, false},
	{false, false, false, true, false, false},
	{false, false, false, false, false, true},
}

// AdjacencyListGraphSample return the adjacency list based graph sample instance
func AdjacencyListGraphSample() graph.Graph {
	return adjacencyListGraphSample
}

// AdjacencyMatrixGraphSample return the adjacency matrix based graph sample instance
func AdjacencyMatrixGraphSample() graph.Graph {
	return adjacencyMatrixGraphSample
}
