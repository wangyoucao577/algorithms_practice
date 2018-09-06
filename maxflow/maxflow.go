// Package maxflow implemented alogrithms to solve maximum flow problem,
// e.g. FordFulkerson, EmondsKarp, Dinic, etc.
package maxflow

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/bfs"
	"github.com/wangyoucao577/algorithms_practice/dfs"
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/networkflowgraph"
)

type graphFlow map[graph.EdgeID]networkflowgraph.FlowUnit

type augmentingPath struct {
	path    graph.Path
	minFlow networkflowgraph.FlowUnit
}

type flowStroage map[graph.EdgeID]networkflowgraph.FlowUnit

// residualNetworkGraph has same structure as NetworkFlowGraph, but will have reverse edges
type residualNetworkGraph struct {
	adjGraph graph.AdjacencyListGraph
	flows    flowStroage
}

// calculateResidualNetwork will calculate residual network with current flow based on network flow graph
func calculateResidualNetwork(g *networkflowgraph.NetworkFlowGraph, f graphFlow) *residualNetworkGraph {
	r := &residualNetworkGraph{graph.AdjacencyListGraph{}, flowStroage{}}

	g.AdjGraph.IterateAllNodes(func(u graph.NodeID) {
		r.adjGraph = append(r.adjGraph, []graph.NodeID{})
	})

	g.AdjGraph.IterateAllNodes(func(u graph.NodeID) {
		g.AdjGraph.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			edge := graph.EdgeID{From: u, To: v}

			edgeFlow := f[edge]
			if edgeFlow > 0 {
				if g.Capacity(edge) > edgeFlow {
					r.adjGraph[u] = append(r.adjGraph[u], v)
					r.flows[edge] = g.Capacity(edge) - edgeFlow
				}

				//reverse edge for residual network graph
				r.adjGraph[v] = append(r.adjGraph[v], u)
				r.flows[edge.Reverse()] = edgeFlow
			} else {
				r.adjGraph[u] = append(r.adjGraph[u], v)
				r.flows[edge] = g.Capacity(edge)
			}
		})
	})

	return r
}

func newGraphFlow(g *networkflowgraph.NetworkFlowGraph) graphFlow {
	newFlow := graphFlow{}

	g.AdjGraph.IterateAllNodes(func(u graph.NodeID) {
		g.AdjGraph.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			edge := graph.EdgeID{From: u, To: v}
			newFlow[edge] = 0
		})
	})

	return newFlow
}

func (f graphFlow) maximumFlow(source graph.NodeID) networkflowgraph.FlowUnit {
	var maximum networkflowgraph.FlowUnit

	for k, v := range f {
		if k.From == source {
			maximum += v
		}
	}

	return maximum
}

func (f graphFlow) minus(a augmentingPath) {

	for i := 0; i < len(a.path)-1; i++ {
		edge := graph.EdgeID{From: a.path[i], To: a.path[i+1]}
		if _, ok := f[edge]; ok {
			f[edge] += a.minFlow
		} else {
			f[edge.Reverse()] -= a.minFlow
		}

	}
}

func (f graphFlow) print() {
	for k, v := range f {
		fmt.Printf("edge %v flow %v, ", k, v)
	}
	fmt.Println()
}

func (r *residualNetworkGraph) calculateResidualCapacity(path graph.Path) augmentingPath {
	a := augmentingPath{path, 0}

	for i := 0; i < len(a.path)-1; i++ {
		edge := graph.EdgeID{From: a.path[i], To: a.path[i+1]}

		if a.minFlow == 0 {
			a.minFlow = r.flows[edge]
		} else {
			if a.minFlow > r.flows[edge] {
				a.minFlow = r.flows[edge]
			}
		}
	}

	return a
}

// FordFulkerson algorithm for maximum flow problem
func FordFulkerson(g *networkflowgraph.NetworkFlowGraph, edmondsKarp bool) int {

	//initialize flow
	currGraphFlow := newGraphFlow(g)

	for {
		currGraphFlow.print()

		// construct new residual network graph
		residualGraph := calculateResidualNetwork(g, currGraphFlow)
		//fmt.Println(residualGraph)

		// try to find augmenting path in the residual network graph
		var newPath graph.Path
		if edmondsKarp {
			bfs, err := bfs.NewBfs(residualGraph.adjGraph, g.Source, nil)
			if err != nil {
				fmt.Println(err)
				break // bfs failed
			}
			_, path := bfs.Query(g.Target)
			if len(path) == 0 {
				fmt.Println("no new valid path by BFS")
				break // no more agumenting path can be found
			}
			newPath = path
		} else {
			dfsSearchedContext, _ := dfs.NewDfs(residualGraph.adjGraph, dfs.Recurse)
			path, err := dfsSearchedContext.RetrievePath(g.Source, g.Target)
			if err != nil {
				fmt.Println(err)
				break // no more agumenting path can be found
			}
			newPath = path
		}
		fmt.Println(newPath)

		newAugmentingPath := residualGraph.calculateResidualCapacity(newPath)
		if newAugmentingPath.minFlow <= 0 {
			fmt.Println("no new valid augmenting path")
			break // no more agumenting path can be found
		}

		// flow - augmentingPathFlow
		currGraphFlow.minus(newAugmentingPath)

	}

	return int(currGraphFlow.maximumFlow(g.Source))
}
