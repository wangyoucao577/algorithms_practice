// Package bfs - Breadth First Search
package bfs

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

type nodeAttr struct {
	Depth  int          // record depth for each node during search
	Viewed bool         // false is WHITE i.e. not viewed, true is BLACK i.e. has been viewed
	Parent graph.NodeID // remember parent node, InvalidNodeID means no parent
}
type nodeAttrArray map[graph.NodeID]*nodeAttr // nodeID is not a int start from 0, so we use `map` instread of `array`

// Bfs defined a structure to store result after BFS search
type Bfs struct {
	Source    graph.NodeID  // BFS start point
	NodesAttr nodeAttrArray // store depth/parent/viewed during BFS
}

// SearchMonitor defined a func to monitor the BFS process
// it will be called if one of the two paramenters changed
type SearchMonitor func(queue []graph.NodeID, currNode graph.NodeID)

// NewBfs execute the BFS search for a specified source on a graph
func NewBfs(g graph.Graph, source graph.NodeID, monitor SearchMonitor) (*Bfs, error) {
	if g.NodeCount() < 2 {
		return nil, fmt.Errorf("Invalid Graph len %d, at least 2 nodes should in the graph", g.NodeCount())
	}
	if !g.IsNodeValid(source) {
		return nil, fmt.Errorf("Invalid Source %v not in the graph", source)
	}

	// Initialize
	bfsContext := &Bfs{source, nodeAttrArray{}}
	g.IterateAllNodes(func(k graph.NodeID) {
		bfsContext.NodesAttr[k] = &nodeAttr{0, false, graph.InvalidNodeID} // create node attr for each node
	})
	bfsContext.NodesAttr[source].Depth = 0
	bfsContext.NodesAttr[source].Viewed = true
	bfsContext.NodesAttr[source].Parent = graph.InvalidNodeID

	var queue []graph.NodeID // next search queue
	queue = append(queue, source)
	if monitor != nil {
		monitor(queue, graph.InvalidNodeID)
	}

	for len(queue) > 0 {
		// pop the first element
		u := queue[0]
		queue = queue[1:]
		if monitor != nil {
			monitor(queue, u)
		}

		currDepth := bfsContext.NodesAttr[u].Depth
		g.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			if !bfsContext.NodesAttr[v].Viewed {
				bfsContext.NodesAttr[v].Depth = currDepth + 1
				bfsContext.NodesAttr[v].Parent = u
				bfsContext.NodesAttr[v].Viewed = true
				queue = append(queue, v)
				if monitor != nil {
					monitor(queue, u)
				}
			}
		})

		bfsContext.NodesAttr[u].Viewed = true
	}

	return bfsContext, nil
}

// Query find shortest path between source and target
// on the BFS searched graph.
// return depth and path if succeed.
func (b *Bfs) Query(target graph.NodeID) (int, graph.Path) {

	currNodeAttr, targetInNodes := b.NodesAttr[target]
	if !targetInNodes {
		panic(fmt.Errorf("target node %v not in the graph", target))
	}

	depth := currNodeAttr.Depth
	shortestPath := []graph.NodeID{}
	shortestPath = append(shortestPath, target)
	for currNodeAttr.Parent != graph.InvalidNodeID {
		currNode := currNodeAttr.Parent
		currNodeAttr = b.NodesAttr[currNode]
		shortestPath = append(shortestPath, currNode)
	}
	for i, j := 0, len(shortestPath)-1; i < j; i, j = i+1, j-1 {
		shortestPath[i], shortestPath[j] = shortestPath[j], shortestPath[i]
	}

	return depth, shortestPath
}
