package main

import (
	"fmt"
	"io"
	"os"

	"github.com/wangyoucao577/algorithms_practice/bfs"
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/graphsamples/graphsample1"
)

func printPath(w io.Writer, s graph.NodeID, t graph.NodeID, depth int, path graph.Path) {
	fmt.Fprintf(w, "%s -> %s shortest path (depth %d) : %s",
		graphsample1.IDToName(s), graphsample1.IDToName(t), depth, graphsample1.IDToName(s))
	for _, v := range path {
		if v != s {
			fmt.Fprintf(w, " -> %s", graphsample1.IDToName(v))
		}
	}
	fmt.Fprintln(w)
}

func main() {

	source := graphsample1.NameToID("s")
	var target graph.NodeID
	var depth int
	var path graph.Path

	fmt.Println()
	fmt.Printf("Run BFS on Adjacency List Based Graph, source %v\n", graphsample1.IDToName(source))
	b1, err := bfs.NewBfs(graphsample1.AdjacencyListGraphSample(), source, nil)
	if err != nil {
		return
	}
	//fmt.Println(b1) // TODO: implement `bfs.String()`

	target = graphsample1.NameToID("v")
	depth, path, _ = b1.Query(target)
	printPath(os.Stdout, source, target, depth, path)

	target = graphsample1.NameToID("x")
	depth, path, _ = b1.Query(target)
	printPath(os.Stdout, source, target, depth, path)

	target = graphsample1.NameToID("y")
	depth, path, _ = b1.Query(target)
	printPath(os.Stdout, source, target, depth, path)

	target = graphsample1.NameToID("u")
	depth, path, _ = b1.Query(target)
	printPath(os.Stdout, source, target, depth, path)

	fmt.Println()
	fmt.Printf("Run BFS on Adjacency Matrix Based Graph, source %v\n", graphsample1.IDToName(source))
	b2, err := bfs.NewBfs(graphsample1.AdjacencyMatrixGraphSample(), source, nil)
	if err != nil {
		return
	}
	//fmt.Println(b2) // TODO: implement `bfs.String()`

	target = graphsample1.NameToID("v")
	depth, path, _ = b2.Query(target)
	printPath(os.Stdout, source, target, depth, path)

	target = graphsample1.NameToID("x")
	depth, path, _ = b2.Query(target)
	printPath(os.Stdout, source, target, depth, path)

	target = graphsample1.NameToID("y")
	depth, path, _ = b2.Query(target)
	printPath(os.Stdout, source, target, depth, path)

	target = graphsample1.NameToID("u")
	depth, path, _ = b2.Query(target)
	printPath(os.Stdout, source, target, depth, path)

}
