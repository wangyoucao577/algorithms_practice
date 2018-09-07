// Package maxflow implemented alogrithms to solve maximum flow problem,
// e.g. FordFulkerson, EmondsKarp, Dinic, etc.
package maxflow

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/bfs"
	"github.com/wangyoucao577/algorithms_practice/dfs"
	"github.com/wangyoucao577/algorithms_practice/flownetwork"
	"github.com/wangyoucao577/algorithms_practice/graph"
)

// residualNetwork has same structure as flownetwork, but will have reverse edges
type residualNetwork struct {
	graph.Graph
	residualCapacities flownetwork.CapacityStorage
}

// calculateResidualNetwork will calculate residual network with current flow based on flow network
func calculateResidualNetwork(fn *flownetwork.FlowNetwork, flow flowStorage) *residualNetwork {
	rn := &residualNetwork{graph.NewAdjacencyListGraph(fn.NodeCount()), flownetwork.CapacityStorage{}}

	fn.IterateAllNodes(func(u graph.NodeID) {
		fn.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			edge := graph.EdgeID{From: u, To: v}

			edgeFlow := flow[edge]
			if edgeFlow > 0 {
				if fn.Capacity(edge) > edgeFlow {
					rn.AddEdge(u, v)
					rn.residualCapacities[edge] = fn.Capacity(edge) - edgeFlow
				}

				//reverse edge for residual network graph
				rn.AddEdge(v, u)
				rn.residualCapacities[edge.Reverse()] = edgeFlow
			} else {
				rn.AddEdge(u, v)
				rn.residualCapacities[edge] = fn.Capacity(edge)
			}
		})
	})

	return rn
}

// FordFulkerson algorithm for maximum flow problem, with/without EdmondKarp improvement
func FordFulkerson(f *flownetwork.FlowNetwork, edmondsKarp bool) flownetwork.EdgeFlowUnit {
	fmt.Printf("FordFulkerson, EmondsKarp %v\n", edmondsKarp)

	//initialize flow
	currFlow := newFlow(f)

	for {
		//currFlow.print()

		// phase 1, construct residual network based on original flow network and current flow
		rn := calculateResidualNetwork(f, currFlow)

		// pahse 2, try to find augmenting path in the residual network graph
		var augmentingPath graph.Path
		if edmondsKarp { //EdmondsKarp use BFS to find a path, better effectiveness
			bfs, err := bfs.NewBfs(rn.Graph, f.Source(), nil)
			if err != nil {
				//fmt.Println(err)
				break // bfs failed
			}
			_, path, err := bfs.Query(f.Target())
			if err != nil {
				//fmt.Println(err)
				break // no more agumenting path can be found
			}
			augmentingPath = path
		} else {
			dfs, _ := dfs.NewDfs(rn.Graph, f.Source(), dfs.Recurse)
			path, err := dfs.Query(f.Source(), f.Target())
			if err != nil {
				//fmt.Println(err)
				break // no more agumenting path can be found
			}
			augmentingPath = path
		}
		//fmt.Println(augmentingPath)

		// phase 3, generate augmenting flow from augmenting path
		augmentingFlow := newAugmentingFlow(rn.residualCapacities, augmentingPath)

		// phase 4, current flow - augmenting flow
		currFlow.minus(augmentingFlow)

	}

	return currFlow.value(f.Source())
}

// Dinic algorithm for maximum flow problem
func Dinic(f *flownetwork.FlowNetwork) flownetwork.EdgeFlowUnit {
	fmt.Println("Dinic")

	//initialize flow
	currFlow := newFlow(f)

	for {
		//currFlow.print()

		// phase 1, construct residual network based on original flow network and current flow
		rn := calculateResidualNetwork(f, currFlow)

		// phase 2, construct level graph from residual network
		lg, err := bfs.NewLevelGraph(rn, f.Source())
		if err != nil {
			fmt.Println(err)
			break // can not construct a new level graph by BFS
		}
		ln := newLevelNetwork(lg, rn)

		// phase 3, try to find blocking flow on the level graph
		//  here we use DFS as typical method
		blockingFlow := flowStorage{}
		for {
			dfs, _ := dfs.NewDfs(lg, f.Source(), dfs.Recurse)
			path, err := dfs.Query(f.Source(), f.Target())
			if err != nil {
				fmt.Println(err)
				break // no more agumenting path can be found
			}

			augmentingFlow := ln.retrieveAndRemoveFlow(path)
			blockingFlow.plus(augmentingFlow)
		}
		if blockingFlow.empty() {
			break // no more blocking flow can be found
		}

		// phase 4, current flow - augmenting flow (blocking flow)
		currFlow.minus(blockingFlow)

	}

	return currFlow.value(f.Source())

}
