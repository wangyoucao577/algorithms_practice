// Package bfs - Breadth First Search
package bfs

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

type color int

const (
	white color = iota // Not find
	gray               // Have found but not scan adjacency list
	black              // Have scaned adjacency list
)

type nodeAttr struct {
	depth     int          // record depth for each node during search
	nodeColor color        // record node status during search
	parent    graph.NodeID // remember parent node, InvalidNodeID means no parent
}
type nodeAttrArray map[graph.NodeID]*nodeAttr // nodeID is not a int start from 0, so we use `map` instread of `array`

// Bfs defined a structure to store result after BFS search
type Bfs struct {
	Source    graph.NodeID  // BFS start point
	nodesAttr nodeAttrArray // store depth/parent/viewed during BFS
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

// NewBfs execute the BFS search for a specified source on a graph
func NewBfs(g graph.Graph, source graph.NodeID, control SearchControl) (*Bfs, error) {
	if g.NodeCount() < 2 {
		return nil, fmt.Errorf("Invalid Graph len %d, at least 2 nodes should in the graph", g.NodeCount())
	}
	if !g.IsNodeValid(source) {
		return nil, fmt.Errorf("Invalid Source %v not in the graph", source)
	}

	// Initialize
	bfsContext := &Bfs{source, nodeAttrArray{}}
	g.IterateAllNodes(func(k graph.NodeID) {
		bfsContext.nodesAttr[k] = &nodeAttr{0, white, graph.InvalidNodeID} // create node attr for each node
	})
	bfsContext.nodesAttr[source].depth = 0
	bfsContext.nodesAttr[source].nodeColor = gray
	bfsContext.nodesAttr[source].parent = graph.InvalidNodeID

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
			}
		})

		bfsContext.nodesAttr[u].nodeColor = black
	}

	return bfsContext, nil
}

// Query find shortest path between source and target
// on the BFS searched graph.
// return depth and path if succeed.
func (b *Bfs) Query(target graph.NodeID) (int, graph.Path, error) {

	currNodeAttr, targetInNodes := b.nodesAttr[target]
	if !targetInNodes {
		panic(fmt.Errorf("target node %v not in the graph", target))
	}

	if target != b.Source && currNodeAttr.parent == graph.InvalidNodeID {
		return 0, nil, fmt.Errorf("no valid path from %v to %v", b.Source, target)
	}

	depth := currNodeAttr.depth
	shortestPath := []graph.NodeID{}
	shortestPath = append(shortestPath, target)
	for currNodeAttr.parent != graph.InvalidNodeID {
		currNode := currNodeAttr.parent
		currNodeAttr = b.nodesAttr[currNode]
		shortestPath = append(shortestPath, currNode)
	}
	for i, j := 0, len(shortestPath)-1; i < j; i, j = i+1, j-1 {
		shortestPath[i], shortestPath[j] = shortestPath[j], shortestPath[i]
	}

	return depth, shortestPath, nil
}
