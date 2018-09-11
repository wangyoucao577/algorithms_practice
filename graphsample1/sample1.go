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
	 |       |  /   |     |
	v(4)   w(5) - x(6) - y(7)
*/

const (
	nodeCount = 8
)

type nodeIDNameConverter struct {
	orderedNodesName []string
	nodeNameToIDMap  map[string]graph.NodeID
}

// define fixed nodes order in the graph, then we use the `index` as nodeID for search,
// will be easier to implement by code.
// node name only for print.
var nodeConverter = nodeIDNameConverter{
	[]string{"r", "s", "t", "u", "v", "w", "x", "y"},
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

// AdjacencyListGraphSample return the adjacency list based graph sample instance
func AdjacencyListGraphSample() graph.Graph {
	sample := graph.NewAdjacencyListGraph(nodeCount)

	return initializeGraphEdges(sample)
}

// AdjacencyMatrixGraphSample return the adjacency matrix based graph sample instance
func AdjacencyMatrixGraphSample() graph.Graph {
	sample := graph.NewAdjacencyMatrixGraph(nodeCount)

	return initializeGraphEdges(sample)
}

func initializeGraphEdges(g graph.Graph) graph.Graph {

	g.AddEdge(0, 1)
	g.AddEdge(0, 4)

	g.AddEdge(1, 0)
	g.AddEdge(1, 5)

	g.AddEdge(2, 3)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)

	g.AddEdge(3, 2)
	g.AddEdge(3, 7)

	g.AddEdge(4, 0)

	g.AddEdge(5, 1)
	g.AddEdge(5, 2)
	g.AddEdge(5, 6)

	g.AddEdge(6, 2)
	g.AddEdge(6, 5)
	g.AddEdge(6, 7)

	g.AddEdge(7, 3)
	g.AddEdge(7, 6)

	return g
}
