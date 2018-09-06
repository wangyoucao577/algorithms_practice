package main

import (
	"fmt"
	"os"

	"github.com/wangyoucao577/algorithms_practice/flownetwork"
	"github.com/wangyoucao577/algorithms_practice/maxflow"
)

func main() {

	var nodeCount, edgeCount int
	fmt.Fscanf(os.Stdin, "%d%d", &edgeCount, &nodeCount)
	if nodeCount < 2 || edgeCount < 0 {
		fmt.Printf("Invalid node count %d edge count %d\n", nodeCount, edgeCount)
		return
	}
	f, err := flownetwork.ConstructFlowNetwork(nodeCount, edgeCount, os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(g)

	flowValue := maxflow.FordFulkerson(f, true)
	fmt.Println(flowValue)
}
