package maxflow

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/bfs"
	"github.com/wangyoucao577/algorithms_practice/dfs"
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/networkflowgraph"
)

type edgeFlow struct {
	f int
}

type graphFlow map[graph.EdgeID]*edgeFlow

type augmentingPath struct {
	path    graph.Path
	minFlow edgeFlow
}

type edgeAttr struct {
	capacity int
}
type edgeAttrArray map[graph.EdgeID]*edgeAttr

// residualNetworkGraph has same structure as NetworkFlowGraph, but will have reverse edges
type residualNetworkGraph struct {
	adjGraph  graph.AdjacencyListGraph
	edgesAttr edgeAttrArray
}

// calculateResidualNetwork will calculate residual network with current flow based on network flow graph
func calculateResidualNetwork(g *networkflowgraph.NetworkFlowGraph, f graphFlow) *residualNetworkGraph {
	r := &residualNetworkGraph{graph.AdjacencyListGraph{}, edgeAttrArray{}}

	g.AdjGraph.IterateAllNodes(func(u graph.NodeID) {
		r.adjGraph = append(r.adjGraph, []graph.NodeID{})
	})

	g.AdjGraph.IterateAllNodes(func(u graph.NodeID) {
		g.AdjGraph.IterateAdjacencyNodes(u, func(v graph.NodeID) {
			edge := graph.EdgeID{From: u, To: v}

			edgeFlow := f[edge]
			if edgeFlow.f > 0 {
				if g.Capacity(edge) > edgeFlow.f {
					r.adjGraph[u] = append(r.adjGraph[u], v)
					r.edgesAttr[edge] = &edgeAttr{g.Capacity(edge) - edgeFlow.f}
				}

				//reverse edge for residual network graph
				r.adjGraph[v] = append(r.adjGraph[v], u)
				r.edgesAttr[edge.Reverse()] = &edgeAttr{edgeFlow.f}
			} else {
				r.adjGraph[u] = append(r.adjGraph[u], v)
				r.edgesAttr[edge] = &edgeAttr{g.Capacity(edge)}
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
			newFlow[edge] = &edgeFlow{0}
		})
	})

	return newFlow
}

func (f graphFlow) maximumFlow(source graph.NodeID) int {
	var maximum int

	for k, v := range f {
		if k.From == source {
			maximum += v.f
		}
	}

	return maximum
}

func (f graphFlow) minus(a augmentingPath) {

	for i := 0; i < len(a.path)-1; i++ {
		edge := graph.EdgeID{From: a.path[i], To: a.path[i+1]}
		if _, ok := f[edge]; ok {
			f[edge].f += a.minFlow.f
		} else {
			f[edge.Reverse()].f -= a.minFlow.f
		}

	}
}

func (f graphFlow) print() {
	for k, v := range f {
		fmt.Printf("edge %v flow %v, ", k, v.f)
	}
	fmt.Println()
}

func (r *residualNetworkGraph) calculateResidualCapacity(path graph.Path) augmentingPath {
	a := augmentingPath{path, edgeFlow{0}}

	for i := 0; i < len(a.path)-1; i++ {
		edge := graph.EdgeID{From: a.path[i], To: a.path[i+1]}

		if a.minFlow.f == 0 {
			a.minFlow.f = r.edgesAttr[edge].capacity
		} else {
			if a.minFlow.f > r.edgesAttr[edge].capacity {
				a.minFlow.f = r.edgesAttr[edge].capacity
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
		if newAugmentingPath.minFlow.f <= 0 {
			fmt.Println("no new valid augmenting path")
			break // no more agumenting path can be found
		}

		// flow - augmentingPathFlow
		currGraphFlow.minus(newAugmentingPath)

	}

	return currGraphFlow.maximumFlow(g.Source)
}
