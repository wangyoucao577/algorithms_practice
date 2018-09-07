package bfs

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/levelgraph"
)

// NewLevelGraph create a new level graph by BFS
func NewLevelGraph(g graph.Graph, source graph.NodeID, control SearchControl) (*levelgraph.LevelGraph, error) {
	if g.NodeCount() < 2 {
		return nil, fmt.Errorf("Invalid Graph len %d, at least 2 nodes should in the graph", g.NodeCount())
	}
	if !g.IsNodeValid(source) {
		return nil, fmt.Errorf("Invalid Source %v not in the graph", source)
	}

	lg := levelgraph.NewLevelGraph(g)

	// Initialize
	bfsContext := &Bfs{source, nodeAttrArray{}}
	g.IterateAllNodes(func(k graph.NodeID) {
		bfsContext.nodesAttr[k] = &nodeAttr{0, white, graph.InvalidNodeID} // create node attr for each node
	})
	bfsContext.nodesAttr[source].depth = 0
	bfsContext.nodesAttr[source].nodeColor = gray
	bfsContext.nodesAttr[source].parent = graph.InvalidNodeID
	lg.SetNodeLevel(source, 0)

	var queue []graph.NodeID // next search queue
	queue = append(queue, source)

	for len(queue) > 0 {
		// pop the first element
		u := queue[0]
		queue = queue[1:]
		if control != nil {
			if control(u) == Break {
				//fmt.Printf("break at node %v\n", u)
				break // break the search in advance by the process control
			}
		}

		currDepth := bfsContext.nodesAttr[u].depth
		g.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			if bfsContext.nodesAttr[v].nodeColor == white {
				bfsContext.nodesAttr[v].depth = currDepth + 1
				bfsContext.nodesAttr[v].parent = u
				bfsContext.nodesAttr[v].nodeColor = gray
				queue = append(queue, v)

				// construct level graph
				lg.SetNodeLevel(v, levelgraph.Level(currDepth))
				lg.AddEdge(u, v)
			}
		})

		bfsContext.nodesAttr[u].nodeColor = black
	}

	return lg, nil
}
