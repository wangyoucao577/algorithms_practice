package main

import (
	"fmt"
)

func main() {

	fmt.Println("Adjacency List Based Graph")
	b1, err := newBFS(adjListGraph, nodeID("s"))
	if err != nil {
		return
	}
	fmt.Println(b1) // TODO: implement `bfs.String()`

	b1.Query(nodeID("v"))
	b1.Query(nodeID("x"))
	b1.Query(nodeID("y"))
	b1.Query(nodeID("u"))

	fmt.Println("Adjacency Matrix Based Graph")
	b2, err := newBFS(adjMatrixGraph, nodeID("s"))
	if err != nil {
		return
	}
	fmt.Println(b2) // TODO: implement `bfs.String()`

	b2.Query(nodeID("v"))
	b2.Query(nodeID("x"))
	b2.Query(nodeID("y"))
	b2.Query(nodeID("u"))

}
