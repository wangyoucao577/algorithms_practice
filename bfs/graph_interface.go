package main

type nodeName string
type nodeID uint // represent each node by `unsigned int` in the code, start from 0

const (
	// InvalidNodeID define max uint
	InvalidNodeID = nodeID(^uint(0))
)

type rangeAction func(nodeID)

type graphOperator interface {
	NodeCount() int
	IsNodeValid(nodeID) bool

	IterateAllNodes(rangeAction)
	IterateAdjacencyNodes(nodeID, rangeAction)
}

/************************* Adjacency  List  Based Graph Representation *****************************/

type adjacencyListGraph [][]nodeID

func (g adjacencyListGraph) NodeCount() int {
	return len(g)
}

func (g adjacencyListGraph) IsNodeValid(currNode nodeID) bool {
	if int(currNode) >= len(g) {
		return false
	}
	return true
}

func (g adjacencyListGraph) IterateAllNodes(action rangeAction) {
	for i := range g {
		action(nodeID(i))
	}
}

func (g adjacencyListGraph) IterateAdjacencyNodes(currNode nodeID, action rangeAction) {
	if !g.IsNodeValid(currNode) {
		return
	}

	for _, v := range g[currNode] {
		action(v)
	}
}

/************************* Adjacency  List  Based Graph Representation *****************************/

/************************* Adjacency Matrix Based Graph Representation *****************************/

type adjacencyMatrixGraph [][]bool

func (g adjacencyMatrixGraph) NodeCount() int {
	return len(g)
}

func (g adjacencyMatrixGraph) IsNodeValid(currNode nodeID) bool {
	if int(currNode) >= len(g) {
		return false
	}
	return true
}

func (g adjacencyMatrixGraph) IterateAllNodes(action rangeAction) {
	for i := range g {
		action(nodeID(i))
	}
}

func (g adjacencyMatrixGraph) IterateAdjacencyNodes(currNode nodeID, action rangeAction) {
	if !g.IsNodeValid(currNode) {
		return
	}

	for i, v := range g[currNode] {
		if v {
			action(nodeID(i))
		}
	}
}

/************************* Adjacency Matrix Based Graph Representation *****************************/
