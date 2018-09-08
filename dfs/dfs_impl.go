package dfs

import "github.com/wangyoucao577/algorithms_practice/graph"

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

func (d *Dfs) initialize(g graph.Graph) {
	d.time = 0
	d.forest = []dfsTree{}

	g.IterateAllNodes(func(k graph.NodeID) {
		d.nodesAttr[k] = &nodeAttr{0, 0, white, graph.InvalidNodeID} // create node attr for each node

		g.IterateAdjacencyNodes(k, func(v graph.NodeID) {
			edge := graph.EdgeID{From: k, To: v}
			d.edgesAttr[edge] = &edgeAttr{unknown}
		})
	})
}
