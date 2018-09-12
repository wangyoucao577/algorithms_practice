package dfs

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

type dfsTopologicalSort struct {
	Dfs
	sorted  []graph.NodeID
	acyclic bool
}

// NewTopologicalSort execute the DFS search on a directed acyclic graph,
// traversing all nodes to get topological sort
func NewTopologicalSort(g graph.Graph) ([]graph.NodeID, error) {

	if !g.Directed() {
		return nil, fmt.Errorf("It's not a Directed Graph")
	}

	// Initialize
	dfsContext := dfsTopologicalSort{
		Dfs{0, []dfsTree{}, nodeAttrArray{}, edgeAttrArray{}},
		[]graph.NodeID{}, false}
	dfsContext.initialize(g)

	// DFS
	g.ControllableIterateAllNodes(func(k graph.NodeID) graph.IterateControl {
		if dfsContext.acyclic {
			return graph.BreakIterate
		}

		if dfsContext.nodesAttr[k].nodeColor == white {
			dfsContext.forest = append(dfsContext.forest, dfsTree{k}) //record a tree's root

			// execute one tree search
			dfsContext.stackBasedVisit(g, k)
		}

		return graph.ContinueIterate
	})

	if dfsContext.acyclic {
		return nil, fmt.Errorf("It's not a Directed Acyclic Graph")
	}

	//reversing in-place
	for i, j := 0, len(dfsContext.sorted)-1; i < j; i, j = i+1, j-1 {
		dfsContext.sorted[i], dfsContext.sorted[j] = dfsContext.sorted[j], dfsContext.sorted[i]
	}

	return dfsContext.sorted, nil
}

func (d *dfsTopologicalSort) stackBasedVisit(g graph.Graph, root graph.NodeID) {
	d.time++
	d.nodesAttr[root].nodeColor = gray
	d.nodesAttr[root].timestampD = d.time

	var stack = []graph.NodeID{}
	stack = append(stack, root)

	for len(stack) > 0 {
		if d.acyclic {
			return //found acyclic in the graph, no need to continue search
		}

		currNode := stack[len(stack)-1]

		newWhiteNodeFound := false
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
			} else if d.nodesAttr[v].nodeColor == gray {
				// backward edge, means it's not a directed acyclic graph
				// so that we can exit the search since we'll not be able to find the topological sort
				d.acyclic = true
				return graph.BreakIterate
			}

			return graph.ContinueIterate
		})

		if !newWhiteNodeFound {
			d.time++
			d.nodesAttr[currNode].nodeColor = black
			d.nodesAttr[currNode].timestampF = d.time

			d.sorted = append(d.sorted, currNode)

			stack = stack[:len(stack)-1] // pop from stack
		}
	}
}
