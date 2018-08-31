package main

import (
	"fmt"
)

func main() {

	fmt.Println("Adjacency List Based Graph")
	b1, err := newBFS(adjListGraph, nodeName("s").ID())
	if err != nil {
		return
	}
	fmt.Println(b1) // TODO: implement `bfs.String()`

	b1.Query(nodeName("v").ID())
	b1.Query(nodeName("x").ID())
	b1.Query(nodeName("y").ID())
	b1.Query(nodeName("u").ID())

	fmt.Println("Adjacency Matrix Based Graph")
	b2, err := newBFS(adjMatrixGraph, nodeName("s").ID())
	if err != nil {
		return
	}
	fmt.Println(b2) // TODO: implement `bfs.String()`

	b2.Query(nodeName("v").ID())
	b2.Query(nodeName("x").ID())
	b2.Query(nodeName("y").ID())
	b2.Query(nodeName("u").ID())

}
