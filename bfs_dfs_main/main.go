package main

import (
	"fmt"
	"io"
	"os"

	"github.com/wangyoucao577/algorithms_practice/bfs"
	"github.com/wangyoucao577/algorithms_practice/graph"
)

func printPath(w io.Writer, s graph.NodeID, t graph.NodeID, depth int, path graph.Path, idToName graph.NodeIDToName) {
	fmt.Fprintf(w, "%s -> %s shortest path (depth %d) : %s",
		idToName.IDToName(s), idToName.IDToName(t), depth, idToName.IDToName(s))
	for _, v := range path {
		if v != s {
			fmt.Fprintf(w, " -> %s", idToName.IDToName(v))
		}
	}
	fmt.Fprintln(w)
}

func main() {

	fmt.Println("Adjacency List Based Graph")
	source := nodeConverter.NameToID("s")
	var target graph.NodeID
	var depth int
	var path graph.Path

	b1, err := bfs.NewBfs(adjListGraph, source, os.Stdout, nodeConverter)
	if err != nil {
		return
	}
	//fmt.Println(b1) // TODO: implement `bfs.String()`

	target = nodeConverter.NameToID("v")
	depth, path = b1.Query(target)
	printPath(os.Stdout, source, target, depth, path, nodeConverter)

	target = nodeConverter.NameToID("x")
	depth, path = b1.Query(target)
	printPath(os.Stdout, source, target, depth, path, nodeConverter)

	target = nodeConverter.NameToID("y")
	depth, path = b1.Query(target)
	printPath(os.Stdout, source, target, depth, path, nodeConverter)

	target = nodeConverter.NameToID("u")
	depth, path = b1.Query(target)
	printPath(os.Stdout, source, target, depth, path, nodeConverter)

	fmt.Println("\nAdjacency Matrix Based Graph")
	b2, err := bfs.NewBfs(adjMatrixGraph, source, os.Stdout, nodeConverter)
	if err != nil {
		return
	}
	//fmt.Println(b2) // TODO: implement `bfs.String()`

	target = nodeConverter.NameToID("v")
	depth, path = b2.Query(target)
	printPath(os.Stdout, source, target, depth, path, nodeConverter)

	target = nodeConverter.NameToID("x")
	depth, path = b2.Query(target)
	printPath(os.Stdout, source, target, depth, path, nodeConverter)

	target = nodeConverter.NameToID("y")
	depth, path = b2.Query(target)
	printPath(os.Stdout, source, target, depth, path, nodeConverter)

	target = nodeConverter.NameToID("u")
	depth, path = b2.Query(target)
	printPath(os.Stdout, source, target, depth, path, nodeConverter)

}
