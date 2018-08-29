// Breadth first Search
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

func adjacencyListBasedBFS(graph adjacencyList, source string) (nodeArray, error) {
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

func adjacencyListBasedQuery() {

}

func main() {
	adjListGraph := adjacencyList{}
	adjListGraph["r"] = []string{"s", "v"}
	adjListGraph["s"] = []string{"r", "w"}
	adjListGraph["t"] = []string{"u", "w", "x"}
	adjListGraph["u"] = []string{"t", "y"}
	adjListGraph["v"] = []string{"r"}
	adjListGraph["w"] = []string{"s", "t", "x"}
	adjListGraph["x"] = []string{"t", "w", "y"}
	adjListGraph["y"] = []string{"u", "x"}

	if nodes, err := adjacencyListBasedBFS(adjListGraph, "s"); err == nil {
		printNodeArray(nodes)
	}

}
