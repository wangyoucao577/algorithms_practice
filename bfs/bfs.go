// Breadth First Search
package main

import "fmt"

type nodeAttr struct {
	Depth  int    // record depth for each node during search
	Viewed bool   // false is WHITE i.e. not viewed, true is BLACK i.e. has been viewed
	Parent nodeID // remember parent node, "" means no parent
}
type nodeAttrArray map[nodeID]*nodeAttr // nodeID is not a int start from 0, so we use `map` instread of `array`

type bfs struct {
	Source    nodeID        // BFS start point
	NodesAttr nodeAttrArray // store depth/parent/viewed during BFS
}

func newBFS(graph graphOperator, source nodeID) (*bfs, error) {
	if graph.NodeCount() < 2 {
		return nil, fmt.Errorf("Invalid Graph len %d, at least 2 nodes should in the graph", graph.NodeCount())
	}
	if !graph.IsNodeValid(source) {
		return nil, fmt.Errorf("Invalid Source %s not in the graph", source)
	}

	// Initialize
	bfsContext := &bfs{source, nodeAttrArray{}}
	graph.IterateAllNodes(func(k nodeID) {
		bfsContext.NodesAttr[k] = &nodeAttr{0, false, ""} // create node attr for each node
	})
	bfsContext.NodesAttr[source].Depth = 0
	bfsContext.NodesAttr[source].Viewed = true
	bfsContext.NodesAttr[source].Parent = ""

	fmt.Printf("BFS Process: ")

	var queue []nodeID // next search queue
	queue = append(queue, source)

	for len(queue) > 0 {
		// pop the first element
		u := queue[0]
		queue = queue[1:]

		currDepth := bfsContext.NodesAttr[u].Depth
		graph.IterateAdjacencyNodes(u, func(v nodeID) {
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

func (b *bfs) Query(target nodeID) error {

	currNodeAttr, targetInNodes := b.NodesAttr[target]
	if !targetInNodes {
		return fmt.Errorf("Invalid Nodes Attr Array: target %s not in the array", target)
	}
	fmt.Printf("%s -> %s depth %d\n", b.Source, target, currNodeAttr.Depth)

	shortestPath := []nodeID{}
	shortestPath = append(shortestPath, target)
	for currNodeAttr.Parent != "" {
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
