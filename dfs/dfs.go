// Package dfs - Depth First Search
package dfs

import "github.com/wangyoucao577/algorithms_practice/graph"

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

// Dfs defined a structure to store result after DFS search
type Dfs struct {
	time      int           //as timestamp during DFS, should be a global var during DFS
	forest    []dfsTree     //generated forest by DFS
	nodesAttr nodeAttrArray // store depth/parent/timestamp during DFS
}

// NewDfs execute the DFS search on a graph
func NewDfs(g graph.Graph, m implementMethod) (*Dfs, error) {

	// Initialize
	dfsContext := &Dfs{0, []dfsTree{}, nodeAttrArray{}}
	g.IterateAllNodes(func(k graph.NodeID) {
		dfsContext.nodesAttr[k] = &nodeAttr{0, 0, white, graph.InvalidNodeID} // create node attr for each node
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
				dfsContext.dfsStackBasedVisit(g, k)
			}
		}
	})
	return dfsContext, nil
}

func (d *Dfs) dfsStackBasedVisit(g graph.Graph, root graph.NodeID) {
	d.time++
	d.nodesAttr[root].nodeColor = gray
	d.nodesAttr[root].timestampD = d.time

	var stack = []graph.NodeID{}
	stack = append(stack, root)

	for len(stack) > 0 {

		currNode := stack[len(stack)-1]

		newWhiteNodeFound := false
		g.ControllableIterateAdjacencyNodes(currNode, func(v graph.NodeID) graph.IterateControl {
			if d.nodesAttr[v].nodeColor == white {
				newWhiteNodeFound = true

				d.nodesAttr[v].parent = currNode

				d.time++
				d.nodesAttr[v].nodeColor = gray
				d.nodesAttr[v].timestampD = d.time

				stack = append(stack, v) // push stack: push to the end
				return graph.BreakIterate
			}
			return graph.ContinueIterate
		})

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
		if d.nodesAttr[v].nodeColor == white {
			d.nodesAttr[v].parent = currNode
			d.dfsRecurseVisit(g, v)
		}
	})

	d.time++
	d.nodesAttr[currNode].nodeColor = black
	d.nodesAttr[currNode].timestampF = d.time
}
