package graph

import "fmt"

/************************* Adjacency  List  Based Graph Representation *****************************/

// AdjacencyListGraph represent a graph by Adjacency List
type AdjacencyListGraph [][]NodeID

// NewAdjacencyListGraph create a adjacency list based graph with nodes
func NewAdjacencyListGraph(nodeCount int) Graph {
	g := AdjacencyListGraph{}
	for i := 0; i < nodeCount; i++ {
		g = append(g, []NodeID{})
	}
	return g
}

// AddEdge add en edge between FromNode and ToNode.
// make sure both FromNode and ToNode are already inside the graph.
func (g AdjacencyListGraph) AddEdge(from, to NodeID) error {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return fmt.Errorf("From node %v or To node %v invalid", from, to)
	}

	g[from] = append(g[from], to)

	return nil
}

// DeleteEdge delete en edge between FromNode and ToNode from the graph.
// make sure both FromNode and ToNode are already inside the graph.
func (g AdjacencyListGraph) DeleteEdge(from, to NodeID) error {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return fmt.Errorf("From node %v or To node %v invalid", from, to)
	}

	for i := 0; i < len(g[from]); i++ {
		if g[from][i] == to {
			g[from] = append(g[from][0:i], g[from][i+1:]...)
			break
		}
	}

	return nil
}

// NodeCount return how many nodes in the graph
func (g AdjacencyListGraph) NodeCount() int {
	return len(g)
}

// EdgeCount return how many edges in the graph
func (g AdjacencyListGraph) EdgeCount() int {
	var count int
	g.IterateAllNodes(func(u NodeID) {
		count += len(g[u])
	})
	return count
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

// ControllableIterateAllNodes for/range on all nodes of the graph,
// call IterateAction for each iterated node.
// break the loop immdiately if action return `BreakIterate`
func (g AdjacencyListGraph) ControllableIterateAllNodes(action IterateActionWithControl) {
	for i := range g {
		condition := action(NodeID(i))
		if condition == BreakIterate {
			break
		}
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

// ReverseIterateAdjacencyNodes reverse for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node
func (g AdjacencyListGraph) ReverseIterateAdjacencyNodes(currNode NodeID, action IterateAction) {
	if !g.IsNodeValid(currNode) {
		return
	}

	adjList := g[currNode]
	for i := len(adjList) - 1; i >= 0; i-- {
		action(adjList[i])
	}
}

// ControllableIterateAdjacencyNodes for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node.
// break the loop immdiately if action return `BreakIterate`
func (g AdjacencyListGraph) ControllableIterateAdjacencyNodes(currNode NodeID, action IterateActionWithControl) {
	if !g.IsNodeValid(currNode) {
		return
	}

	for _, v := range g[currNode] {
		condition := action(v)
		if condition == BreakIterate {
			break
		}
	}
}

/************************* Adjacency  List  Based Graph Representation *****************************/
