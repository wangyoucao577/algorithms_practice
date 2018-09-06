package main

import (
	"fmt"
	"os"

	"github.com/wangyoucao577/algorithms_practice/maxflow"
	"github.com/wangyoucao577/algorithms_practice/networkflowgraph"
)

func main() {

	var nodeCount, edgeCount int
	fmt.Fscanf(os.Stdin, "%d%d", &edgeCount, &nodeCount)
	if nodeCount < 2 || edgeCount < 0 {
		fmt.Printf("Invalid node count %d edge count %d\n", nodeCount, edgeCount)
		return
	}
	g, err := networkflowgraph.ConstructNetworkFlowGraph(nodeCount, edgeCount, os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(g)

	f := maxflow.FordFulkerson(g, true)
	fmt.Println(f)
}
