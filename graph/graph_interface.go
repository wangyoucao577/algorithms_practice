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

	// NodeCount return how many nodes in the graph
	NodeCount() int

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

	// ControllableIterateAdjacencyNodes for/range on all adjacency nodes of current node,
	// call IterateAction for each iterated node.
	// break the loop immdiately if action return `BreakIterate`
	ControllableIterateAdjacencyNodes(NodeID, IterateActionWithControl)
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
