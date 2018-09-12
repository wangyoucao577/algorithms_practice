// Package dfs - Depth First Search
package dfs

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

type implementMethod int

const (
	// Recurse implement DFS by recurse
	Recurse implementMethod = iota

	// StackBased implement DFS by stack and loop
	StackBased
)

type dfsTree struct {
	root graph.NodeID // DFS start node, i.e. root of a tree
}

// Dfs defined a structure to store result after DFS search
type Dfs struct {
	time      int           // as timestamp during DFS, should be a global var during DFS
	forest    []dfsTree     // generated forest by DFS
	nodesAttr nodeAttrArray // store nodes' attr during DFS
	edgesAttr edgeAttrArray // store edges' attr during DFS
}

// SearchControlCondition will control all search functions' behavior, continue or break
type SearchControlCondition int

const (
	// Break will let the search func break immdiately
	Break SearchControlCondition = iota

	// Continue will let the search func go on
	Continue
)

// SearchControl will control all search functions' behavior, continue or break
type SearchControl func(graph.NodeID) SearchControlCondition

// NewDfs execute the DFS search on a graph, start from the root node
func NewDfs(g graph.Graph, root graph.NodeID, m implementMethod) (*Dfs, error) {
	// Initialize
	dfsContext := &Dfs{0, []dfsTree{}, nodeAttrArray{}, edgeAttrArray{}}
	dfsContext.initialize(g)

	dfsContext.forest = append(dfsContext.forest, dfsTree{root}) //record a tree's root

	// execute one tree search
	switch m {
	case Recurse:
		dfsContext.dfsRecurseVisit(g, root)
	case StackBased:
		dfsContext.dfsStackBasedVisit(g, root, nil)
	}

	return dfsContext, nil
}

// NewControllableDfs execute the DFS search on a graph, start from the root node
// can be exit by control condition
// only stack based implemetation
func NewControllableDfs(g graph.Graph, root graph.NodeID, control SearchControl) (*Dfs, error) {
	// Initialize
	dfsContext := &Dfs{0, []dfsTree{}, nodeAttrArray{}, edgeAttrArray{}}
	dfsContext.initialize(g)

	dfsContext.forest = append(dfsContext.forest, dfsTree{root}) //record a tree's root

	// execute one tree search
	dfsContext.dfsStackBasedVisit(g, root, control)

	return dfsContext, nil
}

// NewDfsForest execute the DFS search on a graph, traversing all nodes
func NewDfsForest(g graph.Graph, m implementMethod) (*Dfs, error) {

	// Initialize
	dfsContext := &Dfs{0, []dfsTree{}, nodeAttrArray{}, edgeAttrArray{}}
	dfsContext.initialize(g)

	// DFS
	g.IterateAllNodes(func(k graph.NodeID) {
		if dfsContext.nodesAttr[k].nodeColor == white {
			dfsContext.forest = append(dfsContext.forest, dfsTree{k}) //record a tree's root

			// execute one tree search
			switch m {
			case Recurse:
				dfsContext.dfsRecurseVisit(g, k)
			case StackBased:
				dfsContext.dfsStackBasedVisit(g, k, nil)
			}
		}
	})
	return dfsContext, nil
}

// Query retrieve path from source to target based on a dfs tree/forest
func (d *Dfs) Query(source, target graph.NodeID) (graph.Path, error) {
	path := graph.Path{}

	curr := target
	for {
		path = append(path, curr)

		parent := d.nodesAttr[curr].parent
		if parent == source {
			path = append(path, source)
			break
		} else if parent == graph.InvalidNodeID {
			return nil, fmt.Errorf("no valid path from %v to %v", source, target)
		}

		curr = parent
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path, nil
}
