package graph

import "fmt"

/************************* Adjacency Matrix Based Graph Representation *****************************/

// adjacencyMatrixGraph represent a graph by Adjacency Matrix
type adjacencyMatrixGraph struct {
	adjMatrix [][]bool
	directed  bool //directed graph or undirected graph
}

// NewAdjacencyMatrixGraph create a adjacency list based graph with nodes
func NewAdjacencyMatrixGraph(nodeCount int, directed bool) Graph {

	g := adjacencyMatrixGraph{[][]bool{}, directed}
	for i := 0; i < nodeCount; i++ {
		g.adjMatrix = append(g.adjMatrix, make([]bool, nodeCount, nodeCount))
	}
	return g
}

// Directed return true if it's a directed graph, otherwise it's a undirected graph.
func (g adjacencyMatrixGraph) Directed() bool {
	return g.directed
}

// AddEdge add en edge between FromNode and ToNode.
// make sure both FromNode and ToNode are already inside the graph.
// NOTE: for a undirected graph, this func will add both from-to and to-from into the graph.
func (g adjacencyMatrixGraph) AddEdge(from, to NodeID) error {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return fmt.Errorf("From node %v or To node %v invalid", from, to)
	}

	g.adjMatrix[from][to] = true
	if !g.directed {
		// for undirected graph, add both from-to and to-from into the graph
		g.adjMatrix[to][from] = true
	}

	return nil
}

// DeleteEdge delete en edge between FromNode and ToNode from the graph.
// make sure both FromNode and ToNode are already inside the graph.
// NOTE: for a undirected graph, this func will delete both from-to and to-from from the graph.
func (g adjacencyMatrixGraph) DeleteEdge(from, to NodeID) error {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return fmt.Errorf("From node %v or To node %v invalid", from, to)
	}

	g.adjMatrix[from][to] = false

	if !g.directed {
		// for undirected graph, delete both from-to and to-from from the graph
		g.adjMatrix[to][from] = false
	}

	return nil
}

// IsEdgeValid check whether a edge in or not in the graph
// NOTE: for a undirected graph, it will check both from-to and to-from.
func (g adjacencyMatrixGraph) IsEdgeValid(from, to NodeID) bool {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return false
	}

	var fromToValid, toFromValid bool
	fromToValid = g.adjMatrix[from][to]
	if g.directed {
		return fromToValid
	}

	// for undirected graph, check both from-to and to-from
	toFromValid = g.adjMatrix[to][from]
	return fromToValid && toFromValid
}

// NodeCount return how many nodes in the graph
func (g adjacencyMatrixGraph) NodeCount() int {
	return len(g.adjMatrix)
}

// EdgeCount return how many edges in the graph
// NOTE: for a undirected graph, from-to and to-from count as same edge.
func (g adjacencyMatrixGraph) EdgeCount() int {
	var count int
	g.IterateAllNodes(func(u NodeID) {
		g.IterateAdjacencyNodes(u, func(_ NodeID) {
			count++
		})
	})

	if !g.directed {
		return count / 2
	}
	return count
}

// IsNodeValid check whether a node in or not in the graph
func (g adjacencyMatrixGraph) IsNodeValid(currNode NodeID) bool {
	if int(currNode) >= len(g.adjMatrix) {
		return false
	}
	return true
}

// IterateAllNodes for/range on all nodes of the graph,
// call IterateAction for each iterated node
func (g adjacencyMatrixGraph) IterateAllNodes(action IterateAction) {
	for i := range g.adjMatrix {
		action(NodeID(i))
	}
}

// ControllableIterateAllNodes for/range on all nodes of the graph,
// call IterateAction for each iterated node.
// break the loop immdiately if action return `BreakIterate`
func (g adjacencyMatrixGraph) ControllableIterateAllNodes(action IterateActionWithControl) {
	for i := range g.adjMatrix {
		condition := action(NodeID(i))
		if condition == BreakIterate {
			break
		}
	}
}

// IterateAdjacencyNodes for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node
func (g adjacencyMatrixGraph) IterateAdjacencyNodes(currNode NodeID, action IterateAction) {
	if !g.IsNodeValid(currNode) {
		return
	}

	for i, v := range g.adjMatrix[currNode] {
		if v {
			action(NodeID(i))
		}
	}
}

// ReverseIterateAdjacencyNodes reverse for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node
func (g adjacencyMatrixGraph) ReverseIterateAdjacencyNodes(currNode NodeID, action IterateAction) {
	if !g.IsNodeValid(currNode) {
		return
	}

	adjRow := g.adjMatrix[currNode]
	for i := len(adjRow) - 1; i >= 0; i-- {
		if adjRow[i] {
			action(NodeID(i))
		}
	}
}

// ControllableIterateAdjacencyNodes for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node.
// break the loop immdiately if action return `BreakIterate`
func (g adjacencyMatrixGraph) ControllableIterateAdjacencyNodes(currNode NodeID, action IterateActionWithControl) {
	if !g.IsNodeValid(currNode) {
		return
	}

	for i, v := range g.adjMatrix[currNode] {
		if v {
			condition := action(NodeID(i))
			if condition == BreakIterate {
				break
			}
		}
	}
}

/************************* Adjacency Matrix Based Graph Representation *****************************/
