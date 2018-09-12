// Package graph defined graph representation types and interfaces.
// support both Adjacency List and Adjacency Matrix.
package graph

import "fmt"

// IterateControl will control all iterate functions' behavior, go on or break
type IterateControl int

const (
	// ContinueIterate will let the iterate func go on
	ContinueIterate IterateControl = iota

	// BreakIterate will let the iterate func break immdiately
	BreakIterate
)

// IterateAction will be called in each Iterate functions
type IterateAction func(NodeID)

// IterateActionWithControl will be called in each Iterate functions,
// if it returned with `BreakIterate` then stop the iterate function immediately
type IterateActionWithControl func(NodeID) IterateControl

// Graph defined common interfaces of a graph,
// whatever Adjacency List or Adjacency Matrix based graph
type Graph interface {

	// Directed return true if it's a directed graph, otherwise it's a undirected graph.
	Directed() bool

	// AddEdge add en edge between FromNode and ToNode.
	// make sure both FromNode and ToNode are already inside the graph.
	// NOTE: for a undirected graph, this func will add both from-to and to-from into the graph.
	// Invoker should make sure the edge is not in the graph before call this func.
	AddEdge(NodeID, NodeID) error

	// DeleteEdge delete en edge between FromNode and ToNode from the graph.
	// make sure both FromNode and ToNode are already inside the graph.
	// NOTE: for a undirected graph, this func will delete both from-to and to-from from the graph.
	DeleteEdge(NodeID, NodeID) error

	// IsEdgeValid check whether a edge in or not in the graph
	// NOTE: for a undirected graph, it will check both from-to and to-from.
	IsEdgeValid(NodeID, NodeID) bool

	// NodeCount return how many nodes in the graph
	NodeCount() int

	// EdgeCount return how many edges in the graph
	// NOTE: for a undirected graph, from-to and to-from count as same edge.
	EdgeCount() int

	// IsNodeValid check whether a node in or not in the graph
	IsNodeValid(NodeID) bool

	// IterateAllNodes for/range on all nodes of the graph,
	// call IterateAction for each iterated node
	IterateAllNodes(IterateAction)

	// ControllableIterateAllNodes for/range on all nodes of the graph,
	// call IterateAction for each iterated node.
	// break the loop immdiately if action return `BreakIterate`
	ControllableIterateAllNodes(IterateActionWithControl)

	// IterateAdjacencyNodes for/range on all adjacency nodes of current node,
	// call IterateAction for each iterated node
	IterateAdjacencyNodes(NodeID, IterateAction)

	// ReverseIterateAdjacencyNodes reverse for/range on all adjacency nodes of current node,
	// call IterateAction for each iterated node
	ReverseIterateAdjacencyNodes(NodeID, IterateAction)

	// ControllableIterateAdjacencyNodes for/range on all adjacency nodes of current node,
	// call IterateAction for each iterated node.
	// break the loop immdiately if action return `BreakIterate`
	ControllableIterateAdjacencyNodes(NodeID, IterateActionWithControl)
}

// Transpose generate a new transposed graph from current graph.
// NOTE: only directed graph can have transpose graph, since it's only reverse all edges.
func Transpose(g Graph, newGraphFunc func(int, bool) Graph) (Graph, error) {

	if !g.Directed() {
		return nil, fmt.Errorf("input graph is not a directed graph")
	}

	newGraph := newGraphFunc(g.NodeCount(), g.Directed())

	g.IterateAllNodes(func(u NodeID) {
		g.IterateAdjacencyNodes(u, func(v NodeID) {
			// means u->v exist in current graph
			newGraph.AddEdge(v, u)
		})
	})

	return newGraph, nil
}
