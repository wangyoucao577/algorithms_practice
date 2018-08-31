// Package graphsample1 defined a undirected graph comes from
//  "Introduction to Algorithms - Third Edition" 22.2 BFS
package graphsample1

import "github.com/wangyoucao577/algorithms_practice/graph"

/* This sample undirected graph comes from
  "Introduction to Algorithms - Third Edition" 22.2 BFS

	V = 8 (node count)
	E = 9 (edge count)
	define undirected graph G(V,E) as below (`s` is the source node):

	r(0) - s(1)   t(2) - u(3)
	|     |   /   |     |
	v(4)   w(5) - x(6) - y(7)
*/

type nodeIDNameConverter struct {
	orderedNodesName []string
	nodeNameToIDMap  map[string]graph.NodeID
}

// define fixed nodes order in the graph, then we use the `index` as nodeID for search,
// will be easier to implement by code.
// node name only for print.
var nodeConverter = nodeIDNameConverter{
	[]string{"r", "s", "t", "u", "v", "w", "x", "y"},
	map[string]graph.NodeID{"r": 0, "s": 1, "t": 2, "u": 3, "v": 4, "w": 5, "x": 6, "y": 7},
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

// AdjacencyListGraphSample1 adjacency list based graph sample 1
var AdjacencyListGraphSample1 = graph.AdjacencyListGraph{
	[]graph.NodeID{1, 4},
	[]graph.NodeID{0, 5},
	[]graph.NodeID{3, 5, 6},
	[]graph.NodeID{2, 7},
	[]graph.NodeID{0},
	[]graph.NodeID{1, 2, 6},
	[]graph.NodeID{2, 5, 7},
	[]graph.NodeID{3, 6},
}

// AdjacencyMatrixGraphSample1 adjacency Matrix based graph sample 1
/*
  For this undirected graph, we can only store half of the matrix to save storage if needed

	  r s t u v w x y

  r   0 1 0 0 1 0 0 0
  s   1 0 0 0 0 1 0 0
  t   0 0 0 1 0 1 1 0
  u   0 0 1 0 0 0 0 1
  v   1 0 0 0 0 0 0 0
  w   0 1 1 0 0 0 1 0
  x   0 0 1 0 0 1 0 1
  y   0 0 0 1 0 0 1 0
*/
var AdjacencyMatrixGraphSample1 = graph.AdjacencyMatrixGraph{
	{false, true, false, false, true, false, false, false},
	{true, false, false, false, false, true, false, false},
	{false, false, false, true, false, true, true, false},
	{false, false, true, false, false, false, false, true},
	{true, false, false, false, false, false, false, false},
	{false, true, true, false, false, false, true, false},
	{false, false, true, false, false, true, false, true},
	{false, false, false, true, false, false, true, false},
}
