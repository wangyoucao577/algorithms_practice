// Breadth First Search
package main

import "fmt"

type NodeAttr struct {
	Depth  int    // record depth for each node during search
	Viewed bool   // false is WHITE i.e. not viewed, true is BLACK i.e. has been viewed
	Parent NodeID // remember parent node, InvalidNodeID means no parent
}
type NodeAttrArray map[NodeID]*NodeAttr // nodeID is not a int start from 0, so we use `map` instread of `array`

type Bfs struct {
	Source    NodeID        // BFS start point
	NodesAttr NodeAttrArray // store depth/parent/viewed during BFS
}

func NewBfs(graph Graph, source NodeID) (*Bfs, error) {
	if graph.NodeCount() < 2 {
		return nil, fmt.Errorf("Invalid Graph len %d, at least 2 nodes should in the graph", graph.NodeCount())
	}
	if !graph.IsNodeValid(source) {
		return nil, fmt.Errorf("Invalid Source %s not in the graph", source)
	}

	// Initialize
	bfsContext := &Bfs{source, NodeAttrArray{}}
	graph.IterateAllNodes(func(k NodeID) {
		bfsContext.NodesAttr[k] = &NodeAttr{0, false, InvalidNodeID} // create node attr for each node
	})
	bfsContext.NodesAttr[source].Depth = 0
	bfsContext.NodesAttr[source].Viewed = true
	bfsContext.NodesAttr[source].Parent = InvalidNodeID

	fmt.Printf("BFS Process: ")

	var queue []NodeID // next search queue
	queue = append(queue, source)

	for len(queue) > 0 {
		// pop the first element
		u := queue[0]
		queue = queue[1:]

		currDepth := bfsContext.NodesAttr[u].Depth
		graph.IterateAdjacencyNodes(u, func(v NodeID) {
			if !bfsContext.NodesAttr[v].Viewed {
				bfsContext.NodesAttr[v].Depth = currDepth + 1
				bfsContext.NodesAttr[v].Parent = u
				bfsContext.NodesAttr[v].Viewed = true
				queue = append(queue, v)
			}
		})

		if u != source {
			fmt.Printf(" -> ")
		}
		fmt.Printf("%s", u)
		bfsContext.NodesAttr[u].Viewed = true
	}

	fmt.Println("")
	return bfsContext, nil
}

func (b *Bfs) Query(target NodeID) error {

	currNodeAttr, targetInNodes := b.NodesAttr[target]
	if !targetInNodes {
		return fmt.Errorf("Invalid Nodes Attr Array: target %s not in the array", target)
	}
	fmt.Printf("%s -> %s depth %d\n", b.Source, target, currNodeAttr.Depth)

	shortestPath := []NodeID{}
	shortestPath = append(shortestPath, target)
	for currNodeAttr.Parent != InvalidNodeID {
		currNode := currNodeAttr.Parent
		currNodeAttr = b.NodesAttr[currNode]
		shortestPath = append(shortestPath, currNode)
	}
	for i, j := 0, len(shortestPath)-1; i < j; i, j = i+1, j-1 {
		shortestPath[i], shortestPath[j] = shortestPath[j], shortestPath[i]
	}
	fmt.Printf("  shortest path: ")
	for _, n := range shortestPath {
		fmt.Printf("%s ", n)
	}
	fmt.Println()

	return nil
}
