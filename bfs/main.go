// Breadth First Search
package main

/* This sample undirected graph comes from
  "Introduction to Algorithms - Third Edition" 22.2 BFS

  V = 8 (node count)
  E = 9 (edge count)
  define undirected graph G(V,E) as below (`s` is the source node):

r - s   t - u
|   | / |   |
v   w - x - y
*/

import (
	"fmt"
)

type nodeAttr struct {
	Depth  int    //
	Viewed bool   // false is WHITE i.e. not viewed, true is BLACK i.e. has been viewed
	Parent string // remember parent node, "" means no parent
}
type nodeArray map[string]*nodeAttr
type adjacencyList map[string][]string

type bfs struct {
	Source    string
	NodesAttr map[string]*nodeAttr
}

func printNodeArray(nodes nodeArray) {
	for k, v := range nodes {
		fmt.Printf("node %s -- %v\n", k, v)
	}
}

func newBfs(graph adjacencyList, source string) (*bfs, error) {
	if len(graph) < 2 {
		return nil, fmt.Errorf("Invalid Graph len %d, at least 2 nodes should in the graph", len(graph))
	}
	if _, ok := graph[source]; !ok {
		return nil, fmt.Errorf("Invalid Source %s not in the graph", source)
	}

	// Initialize
	bfsContext := &bfs{source, map[string]*nodeAttr{}}
	for k := range graph { // create node attr for each node
		bfsContext.NodesAttr[k] = &nodeAttr{0, false, ""}
	}
	bfsContext.NodesAttr[source].Depth = 0
	bfsContext.NodesAttr[source].Viewed = true
	bfsContext.NodesAttr[source].Parent = ""

	fmt.Printf("BFS Process: ")

	var queue []string // next search queue
	queue = append(queue, source)

	for len(queue) > 0 {
		// pop the first element
		u := queue[0]
		queue = queue[1:]

		currDepth := bfsContext.NodesAttr[u].Depth
		for _, v := range graph[u] {
			if !bfsContext.NodesAttr[v].Viewed {
				bfsContext.NodesAttr[v].Depth = currDepth + 1
				bfsContext.NodesAttr[v].Parent = u
				bfsContext.NodesAttr[v].Viewed = true
				queue = append(queue, v)
			}
		}

		if u != source {
			fmt.Printf(" -> ")
		}
		fmt.Printf("%s", u)
		bfsContext.NodesAttr[u].Viewed = true
	}

	fmt.Println("")
	return bfsContext, nil
}

func (b *bfs) Query(target string) error {

	currNodeAttr, targetInNodes := b.NodesAttr[target]
	if !targetInNodes {
		return fmt.Errorf("Invalid Nodes Attr Array: target %s not in the array", target)
	}
	fmt.Printf("%s -> %s depth %d\n", b.Source, target, currNodeAttr.Depth)

	shortestPath := []string{}
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

var adjListGraph = adjacencyList{
	"r": {"s", "v"},
	"s": {"r", "w"},
	"t": {"u", "w", "x"},
	"u": {"t", "y"},
	"v": {"r"},
	"w": {"s", "t", "x"},
	"x": {"t", "w", "y"},
	"y": {"u", "x"},
}

func main() {

	b, err := newBfs(adjListGraph, "s")
	if err != nil {
		return
	}
	printNodeArray(b.NodesAttr)

	b.Query("v")
	b.Query("x")
	b.Query("y")
	b.Query("u")
}