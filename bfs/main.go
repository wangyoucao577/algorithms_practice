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

func printNodeArray(nodes nodeArray) {
	for k, v := range nodes {
		fmt.Printf("node %s -- %v\n", k, v)
	}
}

func (graph adjacencyList) BreadthFirstSearch(source string) (nodeArray, error) {
	if len(graph) < 2 {
		return nil, fmt.Errorf("Invalid Graph len %d, at least 2 nodes should in the graph", len(graph))
	}
	if _, ok := graph[source]; !ok {
		return nil, fmt.Errorf("Invalid Source %s not in the graph", source)
	}

	// Initialize
	nodes := nodeArray{}
	for k := range graph { // create node attr for each node
		nodes[k] = &nodeAttr{0, false, ""}
	}
	nodes[source].Depth = 0
	nodes[source].Viewed = true
	nodes[source].Parent = ""

	fmt.Printf("BFS Process: ")

	var queue []string // next search queue
	queue = append(queue, source)

	for len(queue) > 0 {
		// pop the first element
		u := queue[0]
		queue = queue[1:]

		currDepth := nodes[u].Depth
		for _, v := range graph[u] {
			if !nodes[v].Viewed {
				nodes[v].Depth = currDepth + 1
				nodes[v].Parent = u
				nodes[v].Viewed = true
				queue = append(queue, v)
			}
		}

		if u != source {
			fmt.Printf(" -> ")
		}
		fmt.Printf("%s", u)
		nodes[u].Viewed = true
	}

	fmt.Println("")
	return nodes, nil
}

func (graph adjacencyList) Query(source string, target string, nodes nodeArray) error {
	if len(graph) < 2 {
		return fmt.Errorf("Invalid Graph len %d, at least 2 nodes should in the graph", len(graph))
	}
	if _, ok := graph[source]; !ok {
		return fmt.Errorf("Invalid Source %s not in the graph", source)
	}
	if _, ok := graph[target]; !ok {
		return fmt.Errorf("Invalid Target %s not in the graph", target)
	}
	//TODO: check whether `nodes` valid

	currNodeAttr, targetInNodes := nodes[target]
	if !targetInNodes {
		return fmt.Errorf("Invalid Nodes Attr Array: target %s not in the array", target)
	}
	fmt.Printf("%s -> %s depth %d\n", source, target, currNodeAttr.Depth)

	shortestPath := []string{}
	shortestPath = append(shortestPath, target)
	for currNodeAttr.Parent != "" {
		currNode := currNodeAttr.Parent
		currNodeAttr = nodes[currNode]
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

	source := "s"

	nodes, err := adjListGraph.BreadthFirstSearch(source)
	if err != nil {
		return
	}
	printNodeArray(nodes)

	adjListGraph.Query(source, "v", nodes)
	adjListGraph.Query(source, "x", nodes)
	adjListGraph.Query(source, "y", nodes)
	adjListGraph.Query(source, "u", nodes)

}
