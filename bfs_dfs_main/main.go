package main

import (
	"fmt"
	"os"

	"github.com/wangyoucao577/algorithms_practice/bfs"
)

func main() {

	fmt.Println("Adjacency List Based Graph")
	b1, err := bfs.NewBfs(adjListGraph, nodeConverter.NameToID("s"), os.Stdout, nodeConverter)
	if err != nil {
		return
	}
	//fmt.Println(b1) // TODO: implement `bfs.String()`

	b1.Query(nodeConverter.NameToID("v"))
	b1.Query(nodeConverter.NameToID("x"))
	b1.Query(nodeConverter.NameToID("y"))
	b1.Query(nodeConverter.NameToID("u"))

	fmt.Println("\nAdjacency Matrix Based Graph")
	b2, err := bfs.NewBfs(adjMatrixGraph, nodeConverter.NameToID("s"), os.Stdout, nodeConverter)
	if err != nil {
		return
	}
	//fmt.Println(b2) // TODO: implement `bfs.String()`

	b2.Query(nodeConverter.NameToID("v"))
	b2.Query(nodeConverter.NameToID("x"))
	b2.Query(nodeConverter.NameToID("y"))
	b2.Query(nodeConverter.NameToID("u"))

}
