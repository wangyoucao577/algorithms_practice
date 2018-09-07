// Package maxflow implemented alogrithms to solve maximum flow problem,
// e.g. FordFulkerson, EmondsKarp, Dinic, etc.
package maxflow

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/bfs"
	"github.com/wangyoucao577/algorithms_practice/dfs"
	"github.com/wangyoucao577/algorithms_practice/flownetwork"
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/levelgraph"
)

type flowStorage map[graph.EdgeID]flownetwork.EdgeFlowUnit

func (flow flowStorage) empty() bool {
	return len(flow) == 0
}

// residualNetwork has same structure as flownetwork, but will have reverse edges
type residualNetwork struct {
	graph.Graph
	residualCapacities flownetwork.CapacityStorage
}

// levelNetwork grouped level graph with capacities
type levelNetwork struct {
	levelgraph.LevelGraph
	residualCapacities flownetwork.CapacityStorage
}

func newLevelNetwork(lg *levelgraph.LevelGraph, rn *residualNetwork) *levelNetwork {
	ln := &levelNetwork{*lg, flownetwork.CapacityStorage{}}

	ln.IterateAllNodes(func(u graph.NodeID) {
		ln.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			edge := graph.EdgeID{From: u, To: v}

			ln.residualCapacities[edge] = rn.residualCapacities[edge]
		})
	})

	return ln
}

func (ln *levelNetwork) retrieveAndRemoveFlow(path graph.Path) flowStorage {

	flow := newAugmentingFlow(ln.residualCapacities, path)

	for k, v := range flow {
		ln.DeleteEdge(k.From, k.To)
		ln.residualCapacities[k] -= v
	}

	return flow
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

func newFlow(fn *flownetwork.FlowNetwork) flowStorage {
	flow := flowStorage{}

	fn.IterateAllNodes(func(u graph.NodeID) {
		fn.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			edge := graph.EdgeID{From: u, To: v}
			flow[edge] = 0
		})
	})

	return flow
}

func newAugmentingFlow(capacities flownetwork.CapacityStorage, augmentingPath graph.Path) flowStorage {

	// calculate min edge flow
	var minEdgeFlow flownetwork.EdgeFlowUnit
	for i := 0; i < len(augmentingPath)-1; i++ {
		edge := graph.EdgeID{From: augmentingPath[i], To: augmentingPath[i+1]}

		if minEdgeFlow == 0 {
			minEdgeFlow = capacities[edge]
		} else {
			if minEdgeFlow > capacities[edge] {
				minEdgeFlow = capacities[edge]
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

func (f flowStorage) plus(augmentingFlow flowStorage) {

	for k, v := range augmentingFlow {
		if _, ok := f[k]; ok {
			f[k] += v
		} else {
			f[k] = v
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

	return currFlow.Value(f.Source())
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

	return currFlow.Value(f.Source())

}
