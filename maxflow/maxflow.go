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

type flowStorage map[graph.EdgeID]flownetwork.EdgeFlowUnit

// residualNetwork has same structure as flownetwork, but will have reverse edges
type residualNetwork struct {
	adjGraph           graph.AdjacencyListGraph
	residualCapacities flownetwork.CapacityStorage
}

// calculateResidualNetwork will calculate residual network with current flow based on flow network
func calculateResidualNetwork(fn *flownetwork.FlowNetwork, flow flowStorage) *residualNetwork {
	rn := &residualNetwork{graph.AdjacencyListGraph{}, flownetwork.CapacityStorage{}}

	fn.Graph().IterateAllNodes(func(u graph.NodeID) {
		rn.adjGraph = append(rn.adjGraph, []graph.NodeID{})
	})

	fn.Graph().IterateAllNodes(func(u graph.NodeID) {
		fn.Graph().IterateAdjacencyNodes(u, func(v graph.NodeID) {
			edge := graph.EdgeID{From: u, To: v}

			edgeFlow := flow[edge]
			if edgeFlow > 0 {
				if fn.Capacity(edge) > edgeFlow {
					rn.adjGraph[u] = append(rn.adjGraph[u], v)
					rn.residualCapacities[edge] = fn.Capacity(edge) - edgeFlow
				}

				//reverse edge for residual network graph
				rn.adjGraph[v] = append(rn.adjGraph[v], u)
				rn.residualCapacities[edge.Reverse()] = edgeFlow
			} else {
				rn.adjGraph[u] = append(rn.adjGraph[u], v)
				rn.residualCapacities[edge] = fn.Capacity(edge)
			}
		})
	})

	return rn
}

func newFlow(fn *flownetwork.FlowNetwork) flowStorage {
	flow := flowStorage{}

	fn.Graph().IterateAllNodes(func(u graph.NodeID) {
		fn.Graph().IterateAdjacencyNodes(u, func(v graph.NodeID) {
			edge := graph.EdgeID{From: u, To: v}
			flow[edge] = 0
		})
	})

	return flow
}

func newAugmentingFlow(rn *residualNetwork, augmentingPath graph.Path) flowStorage {

	// calculate min edge flow
	var minEdgeFlow flownetwork.EdgeFlowUnit
	for i := 0; i < len(augmentingPath)-1; i++ {
		edge := graph.EdgeID{From: augmentingPath[i], To: augmentingPath[i+1]}

		if minEdgeFlow == 0 {
			minEdgeFlow = rn.residualCapacities[edge]
		} else {
			if minEdgeFlow > rn.residualCapacities[edge] {
				minEdgeFlow = rn.residualCapacities[edge]
			}
		}
	}

	// construct augmenting flow
	augmentingFlow := flowStorage{}
	for i := 0; i < len(augmentingPath)-1; i++ {
		edge := graph.EdgeID{From: augmentingPath[i], To: augmentingPath[i+1]}
		augmentingFlow[edge] = minEdgeFlow
	}

	return augmentingFlow
}

func (f flowStorage) Value(source graph.NodeID) flownetwork.EdgeFlowUnit {
	var maximum flownetwork.EdgeFlowUnit

	for k, v := range f {
		if k.From == source {
			maximum += v
		}
	}

	return maximum
}

func (f flowStorage) minus(augmentingFlow flowStorage) {

	for k, v := range augmentingFlow {
		if _, ok := f[k]; ok {
			f[k] += v
		} else {
			f[k.Reverse()] -= v
		}
	}
}

func (f flowStorage) print() {
	for k, v := range f {
		fmt.Printf("edge %v flow %v, ", k, v)
	}
	fmt.Println()
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
			bfs, err := bfs.NewBfs(rn.adjGraph, f.Source(), nil)
			if err != nil {
				fmt.Println(err)
				break // bfs failed
			}
			_, path, err := bfs.Query(f.Target())
			if err != nil {
				fmt.Println(err)
				break // no more agumenting path can be found
			}
			augmentingPath = path
		} else {
			dfsSearchedContext, _ := dfs.NewDfs(rn.adjGraph, dfs.Recurse)
			path, err := dfsSearchedContext.RetrievePath(f.Source(), f.Target())
			if err != nil {
				fmt.Println(err)
				break // no more agumenting path can be found
			}
			augmentingPath = path
		}
		//fmt.Println(augmentingPath)

		// phase 3, generate augmenting flow from augmenting path
		augmentingFlow := newAugmentingFlow(rn, augmentingPath)

		// phase 4, current flow - augmenting flow
		currFlow.minus(augmentingFlow)

	}

	return currFlow.Value(f.Source())
}
