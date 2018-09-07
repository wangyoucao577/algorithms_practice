// Package graph defined graph representation types and interfaces.
// support both Adjacency List and Adjacency Matrix.
package graph

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

	// AddEdge add en edge between FromNode and ToNode.
	// make sure both FromNode and ToNode are already inside the graph.
	AddEdge(NodeID, NodeID) error

	// DeleteEdge delete en edge between FromNode and ToNode from the graph.
	// make sure both FromNode and ToNode are already inside the graph.
	DeleteEdge(NodeID, NodeID) error

	// NodeCount return how many nodes in the graph
	NodeCount() int

	// EdgeCount return how many edges in the graph
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
