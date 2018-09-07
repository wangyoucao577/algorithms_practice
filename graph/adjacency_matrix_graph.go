package graph

import "fmt"

/************************* Adjacency Matrix Based Graph Representation *****************************/

// adjacencyMatrixGraph represent a graph by Adjacency Matrix
type adjacencyMatrixGraph [][]bool

// NewAdjacencyMatrixGraph create a adjacency list based graph with nodes
func NewAdjacencyMatrixGraph(nodeCount int) Graph {
	g := adjacencyMatrixGraph{}
	for i := 0; i < nodeCount; i++ {
		g = append(g, make([]bool, nodeCount, nodeCount))
	}
	return g
}

// AddEdge add en edge between FromNode and ToNode.
// make sure both FromNode and ToNode are already inside the graph.
func (g adjacencyMatrixGraph) AddEdge(from, to NodeID) error {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return fmt.Errorf("From node %v or To node %v invalid", from, to)
	}

	g[from][to] = true

	return nil
}

// DeleteEdge delete en edge between FromNode and ToNode from the graph.
// make sure both FromNode and ToNode are already inside the graph.
func (g adjacencyMatrixGraph) DeleteEdge(from, to NodeID) error {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return fmt.Errorf("From node %v or To node %v invalid", from, to)
	}

	g.ControllableIterateAdjacencyNodes(from, func(v NodeID) IterateControl {
		if v == to {
			g[from][to] = false
			return BreakIterate
		}
		return ContinueIterate
	})

	return nil
}

// NodeCount return how many nodes in the graph
func (g adjacencyMatrixGraph) NodeCount() int {
	return len(g)
}

// EdgeCount return how many edges in the graph
func (g adjacencyMatrixGraph) EdgeCount() int {
	var count int
	g.IterateAllNodes(func(u NodeID) {
		g.IterateAdjacencyNodes(u, func(_ NodeID) {
			count++
		})
	})
	return count
}

// IsNodeValid check whether a node in or not in the graph
func (g adjacencyMatrixGraph) IsNodeValid(currNode NodeID) bool {
	if int(currNode) >= len(g) {
		return false
	}
	return true
}

// IterateAllNodes for/range on all nodes of the graph,
// call IterateAction for each iterated node
func (g adjacencyMatrixGraph) IterateAllNodes(action IterateAction) {
	for i := range g {
		action(NodeID(i))
	}
}

// ControllableIterateAllNodes for/range on all nodes of the graph,
// call IterateAction for each iterated node.
// break the loop immdiately if action return `BreakIterate`
func (g adjacencyMatrixGraph) ControllableIterateAllNodes(action IterateActionWithControl) {
	for i := range g {
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

	for i, v := range g[currNode] {
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

	adjRow := g[currNode]
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

	for i, v := range g[currNode] {
		if v {
			condition := action(NodeID(i))
			if condition == BreakIterate {
				break
			}
		}
	}
}

/************************* Adjacency Matrix Based Graph Representation *****************************/
