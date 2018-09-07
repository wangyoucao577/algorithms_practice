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

const (
	nodeCount = 6
)

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
	g.AddEdge(0, 3)

	g.AddEdge(1, 4)

	g.AddEdge(2, 4)
	g.AddEdge(2, 5)

	g.AddEdge(3, 1)

	g.AddEdge(4, 3)

	g.AddEdge(5, 5)

	return g
}
