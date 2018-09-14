package graph

import (
	"fmt"
)

/************************* Adjacency  List  Based Graph Representation *****************************/

// adjacencyListGraph represent a graph by Adjacency List
type adjacencyListGraph struct {
	adjList  [][]NodeID
	directed bool //directed graph or undirected graph
}

// NewAdjacencyListGraph create a adjacency list based graph with nodes
func NewAdjacencyListGraph(nodeCount int, directed bool) Graph {

	g := adjacencyListGraph{[][]NodeID{}, directed}
	for i := 0; i < nodeCount; i++ {
		g.adjList = append(g.adjList, []NodeID{})
	}
	return g
}

// Directed return true if it's a directed graph, otherwise it's a undirected graph.
func (g adjacencyListGraph) Directed() bool {
	return g.directed
}

// AddEdge add en edge between FromNode and ToNode.
// make sure both FromNode and ToNode are already inside the graph.
// NOTE: for a undirected graph, this func will add both from-to and to-from into the graph.
// Invoker should make sure the edge is not in the graph before call this func.
func (g adjacencyListGraph) AddEdge(from, to NodeID) error {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return fmt.Errorf("From node %v or To node %v invalid", from, to)
	}

	g.adjList[from] = append(g.adjList[from], to)
	if !g.directed {
		// for undirected graph, add both from-to and to-from into the graph
		g.adjList[to] = append(g.adjList[to], from)
	}

	return nil
}

// DeleteEdge delete en edge between FromNode and ToNode from the graph.
// make sure both FromNode and ToNode are already inside the graph.
// NOTE: for a undirected graph, this func will delete both from-to and to-from from the graph.
func (g adjacencyListGraph) DeleteEdge(from, to NodeID) error {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return fmt.Errorf("From node %v or To node %v invalid", from, to)
	}

	for i := 0; i < len(g.adjList[from]); i++ {
		if g.adjList[from][i] == to {
			g.adjList[from] = append(g.adjList[from][0:i], g.adjList[from][i+1:]...)
			break
		}
	}

	if !g.directed {
		// for undirected graph, delete both from-to and to-from from the graph
		for i := 0; i < len(g.adjList[to]); i++ {
			if g.adjList[to][i] == from {
				g.adjList[to] = append(g.adjList[to][0:i], g.adjList[to][i+1:]...)
				break
			}
		}
	}

	return nil
}

// IsEdgeValid check whether a edge in or not in the graph
// NOTE: for a undirected graph, it will check both from-to and to-from.
func (g adjacencyListGraph) IsEdgeValid(from, to NodeID) bool {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return false
	}

	var fromToValid, toFromValid bool
	for i := 0; i < len(g.adjList[from]); i++ {
		if g.adjList[from][i] == to {
			fromToValid = true
			break
		}
	}
	if g.directed {
		return fromToValid
	}

	// for undirected graph, check both from-to and to-from
	for i := 0; i < len(g.adjList[to]); i++ {
		if g.adjList[to][i] == from {
			toFromValid = true
			break
		}
	}
	return fromToValid && toFromValid
}

// NodeCount return how many nodes in the graph
func (g adjacencyListGraph) NodeCount() int {
	return len(g.adjList)
}

// EdgeCount return how many edges in the graph
// NOTE: for a undirected graph, from-to and to-from count as same edge.
func (g adjacencyListGraph) EdgeCount() int {
	var count int
	g.IterateAllNodes(func(u NodeID) {
		count += len(g.adjList[u])
	})

	if !g.directed {
		return count / 2
	}
	return count
}

// IsNodeValid check whether a node in or not in the graph
func (g adjacencyListGraph) IsNodeValid(currNode NodeID) bool {
	if int(currNode) >= len(g.adjList) {
		return false
	}
	return true
}

// IterateAllNodes for/range on all nodes of the graph,
// call IterateAction for each iterated node
func (g adjacencyListGraph) IterateAllNodes(action IterateAction) {
	for i := range g.adjList {
		action(NodeID(i))
	}
}

// ControllableIterateAllNodes for/range on all nodes of the graph,
// call IterateAction for each iterated node.
// break the loop immdiately if action return `BreakIterate`
func (g adjacencyListGraph) ControllableIterateAllNodes(action IterateActionWithControl) {
	for i := range g.adjList {
		condition := action(NodeID(i))
		if condition == BreakIterate {
			break
		}
	}
}

// IterateAdjacencyNodes for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node
func (g adjacencyListGraph) IterateAdjacencyNodes(currNode NodeID, action IterateAction) {
	if !g.IsNodeValid(currNode) {
		return
	}

	for _, v := range g.adjList[currNode] {
		action(v)
	}
}

// ReverseIterateAdjacencyNodes reverse for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node
func (g adjacencyListGraph) ReverseIterateAdjacencyNodes(currNode NodeID, action IterateAction) {
	if !g.IsNodeValid(currNode) {
		return
	}

	adjList := g.adjList[currNode]
	for i := len(adjList) - 1; i >= 0; i-- {
		action(adjList[i])
	}
}

// ControllableIterateAdjacencyNodes for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node.
// break the loop immdiately if action return `BreakIterate`
func (g adjacencyListGraph) ControllableIterateAdjacencyNodes(currNode NodeID, action IterateActionWithControl) {
	if !g.IsNodeValid(currNode) {
		return
	}

	for _, v := range g.adjList[currNode] {
		condition := action(v)
		if condition == BreakIterate {
			break
		}
	}
}

// IterateEdges for/range on all edges of the graph,
// call ActionOnEdge for each iterated edge
// NOTE: for undirected graph, will only iterate each edge once
func (g adjacencyListGraph) IterateEdges(action ActionOnEdge) {

	directedGraph := g.Directed()

	set := map[EdgeID]struct{}{} //used to filter setteled edge
	g.IterateAllNodes(func(u NodeID) {
		g.IterateAdjacencyNodes(u, func(v NodeID) {
			edge := EdgeID{From: u, To: v}
			if directedGraph { //for directed graph, each edge is valid
				action(edge)
				return
			}

			//for undirected graph, we have to filter same edge(i.e. from-to , to-from)
			_, okFromTo := set[edge]
			_, okToFrom := set[edge.Reverse()]
			if !okFromTo && !okToFrom {
				// not touch before
				action(edge)
				set[edge] = struct{}{}
			}
		})
	})
}

/************************* Adjacency  List  Based Graph Representation *****************************/
