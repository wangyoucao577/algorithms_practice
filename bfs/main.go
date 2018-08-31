package main

import (
	"fmt"
)

func main() {

	fmt.Println("Adjacency List Based Graph")
	b1, err := NewBfs(adjListGraph, NodeName("s").ID())
	if err != nil {
		return
	}
	fmt.Println(b1) // TODO: implement `bfs.String()`

	b1.Query(NodeName("v").ID())
	b1.Query(NodeName("x").ID())
	b1.Query(NodeName("y").ID())
	b1.Query(NodeName("u").ID())

	fmt.Println("Adjacency Matrix Based Graph")
	b2, err := NewBfs(adjMatrixGraph, NodeName("s").ID())
	if err != nil {
		return
	}
	fmt.Println(b2) // TODO: implement `bfs.String()`

	b2.Query(NodeName("v").ID())
	b2.Query(NodeName("x").ID())
	b2.Query(NodeName("y").ID())
	b2.Query(NodeName("u").ID())

}
