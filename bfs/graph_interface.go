package main

type NodeName string
type NodeID uint // represent each node by `unsigned int` in the code, start from 0

const (
	// InvalidNodeID define max uint
	InvalidNodeID = NodeID(^uint(0))
)

type IterateAction func(NodeID)

type Graph interface {
	NodeCount() int
	IsNodeValid(NodeID) bool

	IterateAllNodes(IterateAction)
	IterateAdjacencyNodes(NodeID, IterateAction)
}

/************************* Adjacency  List  Based Graph Representation *****************************/

type AdjacencyListGraph [][]NodeID

func (g AdjacencyListGraph) NodeCount() int {
	return len(g)
}

func (g AdjacencyListGraph) IsNodeValid(currNode NodeID) bool {
	if int(currNode) >= len(g) {
		return false
	}
	return true
}

func (g AdjacencyListGraph) IterateAllNodes(action IterateAction) {
	for i := range g {
		action(NodeID(i))
	}
}

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

type AdjacencyMatrixGraph [][]bool

func (g AdjacencyMatrixGraph) NodeCount() int {
	return len(g)
}

func (g AdjacencyMatrixGraph) IsNodeValid(currNode NodeID) bool {
	if int(currNode) >= len(g) {
		return false
	}
	return true
}

func (g AdjacencyMatrixGraph) IterateAllNodes(action IterateAction) {
	for i := range g {
		action(NodeID(i))
	}
}

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
