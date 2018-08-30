package main

import (
	"fmt"
)

func main() {

	b, err := newBfs(adjListGraph, nodeID("s"))
	if err != nil {
		return
	}
	fmt.Println(b) // TODO: implement `bfs.String()`

	b.Query(nodeID("v"))
	b.Query(nodeID("x"))
	b.Query(nodeID("y"))
	b.Query(nodeID("u"))
}
