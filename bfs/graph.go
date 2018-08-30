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

const (
	nodeCount = 8 // V= 8
)

/************************* Adjacency  List  Based Graph Representation *****************************/

// since we represent node by `string`,
// we have to use `map` instead of `array` to represent the Adjacency List Based Graph
type adjacencyListGraph map[nodeID][]nodeID

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

// since we represent node by `string`,
// we have to use `map` instead of `array` to represent the Adjacency Matrix Based Graph
type adjacencyMatrixGraph map[nodeID][nodeCount]bool

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
	"r": {false, true, false, false, true, false, false, false},
	"s": {true, false, false, false, false, true, false, false},
	"t": {false, false, false, true, false, true, true, false},
	"u": {false, false, true, false, false, false, false, true},
	"v": {true, false, false, false, false, false, false, false},
	"w": {false, true, true, false, false, false, true, false},
	"x": {false, false, true, false, false, true, false, true},
	"y": {false, false, false, true, false, false, true, false},
}

/************************* Adjacency Matrix Based Graph Representation *****************************/

// If we not use `string` to represent node,
// we should define the `String()` for the new type.
func (n nodeID) String() string {
	return string(n)
}
