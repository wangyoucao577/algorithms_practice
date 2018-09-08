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

type color int

const (
	white color = iota // Not find
	gray               // Have found but not scan adjacency list
	black              // Have scaned adjacency list
)

type nodeAttr struct {
	timestampD int          // when this node first found (i.e. when set it to gray)
	timestampF int          // when all subnodes of this node have been searched (i.e. when set it to black)
	nodeColor  color        // record node status during search
	parent     graph.NodeID // remember parent node, InvalidNodeID means no parent
}
type nodeAttrArray map[graph.NodeID]*nodeAttr // nodeID is not a int start from 0, so we use `map` instread of `array`

type dfsTree struct {
	root graph.NodeID // DFS start node, i.e. root of a tree
}

type edgeType int

// below 4 types of edges defined in "Introduction to Algorithms" ch22.3
const (
	unknown  edgeType = iota // unknown before DFS
	branch                   // the edge is one of branches of a DFS tree
	forward                  // it's not a branch but also connected to a descendant
	backward                 // connected to a ancestor (include spin edge)
	cross                    // others, i.e. not above three types
)

type edgeAttr struct {
	t edgeType
}

func (t edgeType) String() string {
	switch t {
	case branch:
		return fmt.Sprintf("branch")
	case forward:
		return fmt.Sprintf("forward")
	case backward:
		return fmt.Sprintf("backward")
	case cross:
		return fmt.Sprintf("cross")
	default:
		return fmt.Sprintf("unknown")
	}
}

type edgeAttrArray map[graph.EdgeID]*edgeAttr

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
	g.IterateAllNodes(func(k graph.NodeID) {
		dfsContext.nodesAttr[k] = &nodeAttr{0, 0, white, graph.InvalidNodeID} // create node attr for each node

		g.IterateAdjacencyNodes(k, func(v graph.NodeID) {
			edge := graph.EdgeID{From: k, To: v}
			dfsContext.edgesAttr[edge] = &edgeAttr{unknown}
		})
	})

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
	g.IterateAllNodes(func(k graph.NodeID) {
		dfsContext.nodesAttr[k] = &nodeAttr{0, 0, white, graph.InvalidNodeID} // create node attr for each node

		g.IterateAdjacencyNodes(k, func(v graph.NodeID) {
			edge := graph.EdgeID{From: k, To: v}
			dfsContext.edgesAttr[edge] = &edgeAttr{unknown}
		})
	})

	dfsContext.forest = append(dfsContext.forest, dfsTree{root}) //record a tree's root

	// execute one tree search
	dfsContext.dfsStackBasedVisit(g, root, control)

	return dfsContext, nil
}

// NewDfsForest execute the DFS search on a graph, traversing all nodes
func NewDfsForest(g graph.Graph, m implementMethod) (*Dfs, error) {

	// Initialize
	dfsContext := &Dfs{0, []dfsTree{}, nodeAttrArray{}, edgeAttrArray{}}
	g.IterateAllNodes(func(k graph.NodeID) {
		dfsContext.nodesAttr[k] = &nodeAttr{0, 0, white, graph.InvalidNodeID} // create node attr for each node

		g.IterateAdjacencyNodes(k, func(v graph.NodeID) {
			edge := graph.EdgeID{From: k, To: v}
			dfsContext.edgesAttr[edge] = &edgeAttr{unknown}
		})
	})

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

func (d *Dfs) dfsStackBasedVisit(g graph.Graph, root graph.NodeID, control SearchControl) {
	d.time++
	d.nodesAttr[root].nodeColor = gray
	d.nodesAttr[root].timestampD = d.time

	var stack = []graph.NodeID{}
	stack = append(stack, root)

	exitNow := false // whether exit now or go on traversing

	for len(stack) > 0 {

		currNode := stack[len(stack)-1]
		if control != nil {
			if control(currNode) == Break {
				//fmt.Printf("break at node %v\n", u)
				exitNow = true
			}
		}

		newWhiteNodeFound := false

		// if match exit condition, will not try new nodes
		// but we still hope to finish nodes which have already been viewed
		if !exitNow {
			g.ControllableIterateAdjacencyNodes(currNode, func(v graph.NodeID) graph.IterateControl {
				edge := graph.EdgeID{From: currNode, To: v}
				if d.nodesAttr[v].nodeColor == white {
					newWhiteNodeFound = true

					d.nodesAttr[v].parent = currNode

					d.time++
					d.nodesAttr[v].nodeColor = gray
					d.nodesAttr[v].timestampD = d.time

					stack = append(stack, v) // push stack: push to the end

					// set attr for edge
					d.edgesAttr[edge].t = branch
					return graph.BreakIterate
				} else if d.nodesAttr[v].nodeColor == black {
					// stack based implementation will see same edge more than 1 time,
					// let's igore not-first see
					if d.edgesAttr[edge].t == unknown {
						if d.nodesAttr[currNode].timestampD < d.nodesAttr[v].timestampD {
							d.edgesAttr[edge].t = forward
						} else {
							d.edgesAttr[edge].t = cross
						}
					}
				} else if d.nodesAttr[v].nodeColor == gray {
					d.edgesAttr[edge].t = backward
				}

				return graph.ContinueIterate
			})
		}

		if !newWhiteNodeFound {
			d.time++
			d.nodesAttr[currNode].nodeColor = black
			d.nodesAttr[currNode].timestampF = d.time

			stack = stack[:len(stack)-1] // pop from stack
		}
	}
}

func (d *Dfs) dfsRecurseVisit(g graph.Graph, currNode graph.NodeID) {
	d.time++
	d.nodesAttr[currNode].nodeColor = gray
	d.nodesAttr[currNode].timestampD = d.time

	g.IterateAdjacencyNodes(currNode, func(v graph.NodeID) {
		edge := graph.EdgeID{From: currNode, To: v}
		if d.nodesAttr[v].nodeColor == white {
			d.nodesAttr[v].parent = currNode
			d.dfsRecurseVisit(g, v)

			// set attr for edge
			d.edgesAttr[edge].t = branch
		} else if d.nodesAttr[v].nodeColor == gray {
			d.edgesAttr[edge].t = backward
		} else if d.nodesAttr[v].nodeColor == black {
			if d.nodesAttr[currNode].timestampD < d.nodesAttr[v].timestampD {
				d.edgesAttr[edge].t = forward
			} else {
				d.edgesAttr[edge].t = cross
			}
		}
	})

	d.time++
	d.nodesAttr[currNode].nodeColor = black
	d.nodesAttr[currNode].timestampF = d.time
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
