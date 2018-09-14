package weightedgraph

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/graph"
)

// Weight represent a numberical value for each edge
type Weight int

// WeightedGraph represent a graph, each edge has a weight
type WeightedGraph struct {
	graph.Graph

	weights map[graph.EdgeID]Weight
}

// NewWeightedGraph create a new empty weighted graph
func NewWeightedGraph(nodeCount int, directed bool, newGraphFunc func(int, bool) graph.Graph) *WeightedGraph {
	g := &WeightedGraph{newGraphFunc(nodeCount, directed),
		map[graph.EdgeID]Weight{}}
	return g
}

// AddEdge add a weighted edge into the weighted graph
func (g *WeightedGraph) AddEdge(from, to graph.NodeID, w Weight) error {

	err := g.Graph.AddEdge(from, to)
	if err != nil {
		return err
	}

	g.weights[graph.EdgeID{From: from, To: to}] = w

	return nil
}

// Weight return the weight of the edge
func (g *WeightedGraph) Weight(from, to graph.NodeID) (Weight, error) {

	if !g.IsEdgeValid(from, to) {
		return 0, fmt.Errorf("Edge %v-%v is invalid", from, to)
	}

	edge := graph.EdgeID{From: from, To: to}

	fromToWeight, ok := g.weights[edge]
	if ok {
		return fromToWeight, nil
	}

	if g.Directed() {
		return 0, fmt.Errorf("no weight found for edge %v", edge)
	}

	//for undirected graph, still need to check reverse
	toFromWeight, ok := g.weights[edge.Reverse()]
	if ok {
		return toFromWeight, nil
	}

	// for undirected graph, both from-to and to-from weight not found
	return 0, fmt.Errorf("no weight found for edge %v", edge)
}

// Validate check whether a valid weighted graph
// i.e. all edges have a weight
func (g *WeightedGraph) Validate() bool {

	if g.EdgeCount() != len(g.weights) {
		return false
	}

	directedGraph := g.Directed()
	valid := true

	g.ControllableIterateAllNodes(func(u graph.NodeID) graph.IterateControl {
		g.ControllableIterateAdjacencyNodes(u, func(v graph.NodeID) graph.IterateControl {
			edge := graph.EdgeID{From: u, To: v}

			_, ok := g.weights[edge]
			if ok {
				return graph.ContinueIterate
			}

			if !directedGraph { //undirected graph, also check reverse
				_, reverseOk := g.weights[edge.Reverse()]
				if reverseOk {
					return graph.ContinueIterate
				}
			}

			valid = false
			return graph.BreakIterate
		})

		if !valid {
			return graph.BreakIterate
		}

		return graph.ContinueIterate
	})

	return valid
}
