// Package graphsample3 defined a Directed Acyclic Graph comes from
//  "Introduction to Algorithms - Third Edition" 22.4
package graphsample3

import "github.com/wangyoucao577/algorithms_practice/graph"

/* This sample directed acyclic graph comes from
  "Introduction to Algorithms - Third Edition" 22.4

	V = 9 (node count)
	E = 9 (edge count)
	define directed graph G(V,E) as below:

	0   1   8
	↓ ↘ ↓
	2 → 3
	↓
	5 ← 4
		↓
	 ↘	6
		↓
		7

	0 - 内裤
	1 - 袜子
	2 - 裤子
	3 - 鞋
	4 - 衬衣
	5 - 腰带
	6 - 领带
	7 - 夹克
	8 - 手表

*/

const (
	nodeCount = 9
)

type nodeIDNameConverter struct {
	orderedNodesName []string
	nodeNameToIDMap  map[string]graph.NodeID
}

// define fixed nodes order in the graph, then we use the `index` as nodeID for search,
// will be easier to implement by code.
// node name only for print.
var nodeConverter = nodeIDNameConverter{
	[]string{"内裤", "袜子", "裤子", "鞋", "衬衣", "腰带", "领带", "夹克", "手表"},
	map[string]graph.NodeID{"内裤": 0, "袜子": 1, "裤子": 2, "鞋": 3, "衬衣": 4, "腰带": 5, "领带": 6, "夹克": 7, "手表": 8},
}

// IDToName convert NodeID to human readable name
func IDToName(i graph.NodeID) string {
	if i == graph.InvalidNodeID {
		return "InvalidNodeID"
	}
	return nodeConverter.orderedNodesName[i]
}

// IDSToNames convert NodeIDs to human readable names
func IDSToNames(ids []graph.NodeID) string {
	s := ""
	for _, v := range ids {
		s += IDToName(v) + " "
	}
	return s
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

	g.AddEdge(0, 2)
	g.AddEdge(0, 3)

	g.AddEdge(1, 3)

	g.AddEdge(2, 3)
	g.AddEdge(2, 5)

	g.AddEdge(4, 5)
	g.AddEdge(4, 6)

	g.AddEdge(5, 7)

	g.AddEdge(6, 7)
	return g
}
