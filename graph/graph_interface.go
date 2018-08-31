// Package graph defined graph representation types and interfaces.
// support both Adjacency List and Adjacency Matrix.
package graph

// NodeID represent each node by `unsigned int`, start from 0
type NodeID uint

const (
	// InvalidNodeID defined invalid value of NodeID
	InvalidNodeID = NodeID(^uint(0))
)

// Path represented by nodes
type Path []NodeID

// IterateAction will be called in each Iterate functions
type IterateAction func(NodeID)

// Graph defined common interfaces of a graph,
// whatever Adjacency List or Adjacency Matrix based graph
type Graph interface {

	// NodeCount return how many nodes in the graph
	NodeCount() int

	// IsNodeValid check whether a node in or not in the graph
	IsNodeValid(NodeID) bool

	// IterateAllNodes for/range on all nodes of the graph,
	// call IterateAction for each iterated node
	IterateAllNodes(IterateAction)

	// IterateAdjacencyNodes for/range on all adjacency nodes of current node,
	// call IterateAction for each iterated node
	IterateAdjacencyNodes(NodeID, IterateAction)
}

/************************* Adjacency  List  Based Graph Representation *****************************/

// AdjacencyListGraph represent a graph by Adjacency List
type AdjacencyListGraph [][]NodeID

// NodeCount return how many nodes in the graph
func (g AdjacencyListGraph) NodeCount() int {
	return len(g)
}

// IsNodeValid check whether a node in or not in the graph
func (g AdjacencyListGraph) IsNodeValid(currNode NodeID) bool {
	if int(currNode) >= len(g) {
		return false
	}
	return true
}

// IterateAllNodes for/range on all nodes of the graph,
// call IterateAction for each iterated node
func (g AdjacencyListGraph) IterateAllNodes(action IterateAction) {
	for i := range g {
		action(NodeID(i))
	}
}

// IterateAdjacencyNodes for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node
func (g AdjacencyListGraph) IterateAdjacencyNodes(currNode NodeID, action IterateAction) {
	if !g.IsNodeValid(currNode) {
		return
	}

	for _, v := range g[currNode] {
		action(v)
	}
}

/************************* Adjacency  List  Based Graph Representation *****************************/

/************************* Adjacency Matrix Based Graph Representation *****************************/

// AdjacencyMatrixGraph represent a graph by Adjacency Matrix
type AdjacencyMatrixGraph [][]bool

// NodeCount return how many nodes in the graph
func (g AdjacencyMatrixGraph) NodeCount() int {
	return len(g)
}

// IsNodeValid check whether a node in or not in the graph
func (g AdjacencyMatrixGraph) IsNodeValid(currNode NodeID) bool {
	if int(currNode) >= len(g) {
		return false
	}
	return true
}

// IterateAllNodes for/range on all nodes of the graph,
// call IterateAction for each iterated node
func (g AdjacencyMatrixGraph) IterateAllNodes(action IterateAction) {
	for i := range g {
		action(NodeID(i))
	}
}

// IterateAdjacencyNodes for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node
func (g AdjacencyMatrixGraph) IterateAdjacencyNodes(currNode NodeID, action IterateAction) {
	if !g.IsNodeValid(currNode) {
		return
	}

	for i, v := range g[currNode] {
		if v {
			action(NodeID(i))
		}
	}
}

/************************* Adjacency Matrix Based Graph Representation *****************************/

//Equal to compare whether current Path equal to another one
func (p Path) Equal(q Path) bool {
	if len(p) != len(q) {
		return false
	}

	if (p == nil) != (q == nil) {
		return false
	}

	for i := range p {
		if p[i] != q[i] {
			return false
		}
	}

	return true
}
