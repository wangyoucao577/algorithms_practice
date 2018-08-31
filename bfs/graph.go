package main

/* This sample undirected graph comes from
  "Introduction to Algorithms - Third Edition" 22.2 BFS

  V = 8 (node count)
  E = 9 (edge count)
  define undirected graph G(V,E) as below (`s` is the source node):

r - s   t - u
|   | / |   |
v   w - x - y
*/

type nodeID string // represent each node by `string` in the code

type rangeAction func(nodeID)

type graphOperator interface {
	NodeCount() int
	IsNodeValid(nodeID) bool

	IterateAllNodes(rangeAction)
	IterateAdjacencyNodes(nodeID, rangeAction)
}

/************************* Adjacency  List  Based Graph Representation *****************************/

// since we represent node by `string`,
// we have to use `map` instead of `array` to represent the Adjacency List Based Graph
type adjacencyListGraph map[nodeID][]nodeID

func (g adjacencyListGraph) NodeCount() int {
	return len(g)
}

func (g adjacencyListGraph) IsNodeValid(currNode nodeID) bool {
	_, ok := g[currNode]
	return ok
}

func (g adjacencyListGraph) IterateAllNodes(action rangeAction) {
	for k := range g {
		action(k)
	}
}

func (g adjacencyListGraph) IterateAdjacencyNodes(currNode nodeID, action rangeAction) {
	for _, v := range g[currNode] {
		action(v)
	}
}

var adjListGraph = adjacencyListGraph{
	"r": {"s", "v"},
	"s": {"r", "w"},
	"t": {"u", "w", "x"},
	"u": {"t", "y"},
	"v": {"r"},
	"w": {"s", "t", "x"},
	"x": {"t", "w", "y"},
	"y": {"u", "x"},
}

/************************* Adjacency  List  Based Graph Representation *****************************/

/************************* Adjacency Matrix Based Graph Representation *****************************/

type adjacencyMatrixGraph struct {
	NodesOrder map[nodeID]int
	Matrix     [][]bool
}

func (g adjacencyMatrixGraph) NodeCount() int {
	return len(g.NodesOrder)
}

func (g adjacencyMatrixGraph) IsNodeValid(currNode nodeID) bool {
	_, ok := g.NodesOrder[currNode]
	return ok
}

func (g adjacencyMatrixGraph) IterateAllNodes(action rangeAction) {
	for k := range g.NodesOrder {
		action(k)
	}
}

func (g adjacencyMatrixGraph) IterateAdjacencyNodes(currNode nodeID, action rangeAction) {
	row, ok := g.NodesOrder[currNode]
	if !ok {
		return
	}

	for k, v := range g.NodesOrder {
		if g.Matrix[row][v] {
			action(k)
		}
	}
}

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
var adjMatrixGraph = adjacencyMatrixGraph{
	NodesOrder: map[nodeID]int{"r": 0, "s": 1, "t": 2, "u": 3, "v": 4, "w": 5, "x": 6, "y": 7},
	Matrix: [][]bool{
		{false, true, false, false, true, false, false, false},
		{true, false, false, false, false, true, false, false},
		{false, false, false, true, false, true, true, false},
		{false, false, true, false, false, false, false, true},
		{true, false, false, false, false, false, false, false},
		{false, true, true, false, false, false, true, false},
		{false, false, true, false, false, true, false, true},
		{false, false, false, true, false, false, true, false},
	},
}

/************************* Adjacency Matrix Based Graph Representation *****************************/

// If we not use `string` to represent node,
// we should define the `String()` for the new type.
func (n nodeID) String() string {
	return string(n)
}
