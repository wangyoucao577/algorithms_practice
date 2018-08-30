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

// If we not use `string` to represent node,
// we should define the `String()` for the new type.
func (n nodeID) String() string {
	return string(n)
}
