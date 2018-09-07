// Package graph defined graph representation types and interfaces.
// support both Adjacency List and Adjacency Matrix.
package graph

import (
	"fmt"
)

// NodeID represent each node by `unsigned int`, start from 0
type NodeID uint

const (
	// InvalidNodeID defined invalid value of NodeID
	InvalidNodeID = NodeID(^uint(0))
)

// EdgeID represent an edge between two nodes.
// If it's an undirected edge, From and To can be swapped.
type EdgeID struct {
	From NodeID
	To   NodeID
}

// Path represented by nodes
type Path []NodeID

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

/************************* Adjacency Matrix Based Graph Representation *****************************/

// AdjacencyMatrixGraph represent a graph by Adjacency Matrix
type AdjacencyMatrixGraph [][]bool

// NewAdjacencyMatrixGraph create a adjacency list based graph with nodes
func NewAdjacencyMatrixGraph(nodeCount int) Graph {
	g := AdjacencyMatrixGraph{}
	for i := 0; i < nodeCount; i++ {
		g = append(g, make([]bool, nodeCount, nodeCount))
	}
	return g
}

// AddEdge add en edge between FromNode and ToNode.
// make sure both FromNode and ToNode are already inside the graph.
func (g AdjacencyMatrixGraph) AddEdge(from, to NodeID) error {
	if !g.IsNodeValid(from) || !g.IsNodeValid(to) {
		return fmt.Errorf("From node %v or To node %v invalid", from, to)
	}

	g[from][to] = true

	return nil
}

// NodeCount return how many nodes in the graph
func (g AdjacencyMatrixGraph) NodeCount() int {
	return len(g)
}

// EdgeCount return how many edges in the graph
func (g AdjacencyMatrixGraph) EdgeCount() int {
	var count int
	g.IterateAllNodes(func(u NodeID) {
		g.IterateAdjacencyNodes(u, func(_ NodeID) {
			count++
		})
	})
	return count
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

// ControllableIterateAllNodes for/range on all nodes of the graph,
// call IterateAction for each iterated node.
// break the loop immdiately if action return `BreakIterate`
func (g AdjacencyMatrixGraph) ControllableIterateAllNodes(action IterateActionWithControl) {
	for i := range g {
		condition := action(NodeID(i))
		if condition == BreakIterate {
			break
		}
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

// ReverseIterateAdjacencyNodes reverse for/range on all adjacency nodes of current node,
// call IterateAction for each iterated node
func (g AdjacencyMatrixGraph) ReverseIterateAdjacencyNodes(currNode NodeID, action IterateAction) {
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
func (g AdjacencyMatrixGraph) ControllableIterateAdjacencyNodes(currNode NodeID, action IterateActionWithControl) {
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

//Equal to compare whether current Path equal to another one
func (p Path) Equal(q Path) bool {
	//NOTE: could use reflect.DeepEqual() instead.
	// but we implement it manually to avoid `import reflect`.
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

//UndirectedEqual to check whether two undirected edges are equal
func (e EdgeID) UndirectedEqual(f EdgeID) bool {
	if e == f {
		return true
	}
	if e.From == f.To && e.To == f.From {
		return true
	}
	return false
}

//Reverse return reverse edge of current one
func (e EdgeID) Reverse() EdgeID {
	return EdgeID{e.To, e.From}
}

//IsValid return whether the edgeID is valid
//if From or To is InvalidNodeID, the edgeID is invalid
func (e EdgeID) IsValid() bool {
	return e.From != InvalidNodeID && e.To != InvalidNodeID
}
