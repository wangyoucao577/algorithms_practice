package main

/* This sample undirected graph comes from
  "Introduction to Algorithms - Third Edition" 22.2 BFS

	V = 8 (node count)
	E = 9 (edge count)
	define undirected graph G(V,E) as below (`s` is the source node):

	r(0) - s(1)   t(2) - u(3)
	|     |   /   |     |
	v(4)   w(5) - x(6) - y(7)
*/

func (n nodeName) String() string {
	return string(n)
}

func (n nodeID) String() string {
	return n.Name().String()
}

// define fixed nodes order in the graph, then we use the `index` as nodeID for search,
// will be easier to implement by code.
// node name only for print.
var orderedNodesName = [...]nodeName{"r", "s", "t", "u", "v", "w", "x", "y"}

func (n nodeID) Name() nodeName {
	return orderedNodesName[n]
}

var nodeNameToIDMap = map[nodeName]nodeID{"r": 0, "s": 1, "t": 2, "u": 3, "v": 4, "w": 5, "x": 6, "y": 7}

func (n nodeName) ID() nodeID {
	return nodeNameToIDMap[n]
}

/************************* Adjacency  List  Based Graph Representation *****************************/
var adjListGraph = adjacencyListGraph{
	[]nodeID{1, 4},
	[]nodeID{0, 5},
	[]nodeID{3, 5, 6},
	[]nodeID{2, 7},
	[]nodeID{0},
	[]nodeID{1, 2, 6},
	[]nodeID{2, 5, 7},
	[]nodeID{3, 6},
}

/************************* Adjacency  List  Based Graph Representation *****************************/

/************************* Adjacency Matrix Based Graph Representation *****************************/
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
	{false, true, false, false, true, false, false, false},
	{true, false, false, false, false, true, false, false},
	{false, false, false, true, false, true, true, false},
	{false, false, true, false, false, false, false, true},
	{true, false, false, false, false, false, false, false},
	{false, true, true, false, false, false, true, false},
	{false, false, true, false, false, true, false, true},
	{false, false, false, true, false, false, true, false},
}

/************************* Adjacency Matrix Based Graph Representation *****************************/
