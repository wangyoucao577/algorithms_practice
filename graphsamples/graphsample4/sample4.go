// Package graphsample4 defined a directed graph comes from
//  "Introduction to Algorithms - Third Edition" 22.5
package graphsample4

import "github.com/wangyoucao577/algorithms_practice/graph"

/* This sample directed graph comes from
  "Introduction to Algorithms - Third Edition" 22.5 strongly connected component

	V = 8 (node count)
	E = 13 (edge count)
	define directed graph G(V,E) as below:

	a(0) -> b(1)  →  c(2) ← → d(3)
	  ↑  ↙   ↓        ↓        ↓
	e(4) →  f(5) ← → g(6)  →  h(7) ->
								↑    ↓
								  <-

    NOTE: node `h` has a spin edge pointer to itself.
*/

const (
	nodeCount     = 8
	directedGraph = true
)

func initializeGraphEdges(g graph.Graph) graph.Graph {

	g.AddEdge(0, 1)

	g.AddEdge(1, 2)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)

	g.AddEdge(2, 3)
	g.AddEdge(2, 6)

	g.AddEdge(3, 2)
	g.AddEdge(3, 7)

	g.AddEdge(4, 0)
	g.AddEdge(4, 5)

	g.AddEdge(5, 6)

	g.AddEdge(6, 5)
	g.AddEdge(6, 7)

	g.AddEdge(7, 7)
	return g
}

type nodeIDNameConverter struct {
	orderedNodesName []string
	nodeNameToIDMap  map[string]graph.NodeID
}

// define fixed nodes order in the graph, then we use the `index` as nodeID for search,
// will be easier to implement by code.
// node name only for print.
var nodeConverter = nodeIDNameConverter{
	[]string{"a", "b", "c", "d", "e", "f", "g", "h"},
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
	sample := graph.NewAdjacencyListGraph(nodeCount, directedGraph)

	return initializeGraphEdges(sample)
}

// AdjacencyMatrixGraphSample return the adjacency matrix based graph sample instance
func AdjacencyMatrixGraphSample() graph.Graph {
	sample := graph.NewAdjacencyMatrixGraph(nodeCount, directedGraph)

	return initializeGraphEdges(sample)
}
