package minspanningtree

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// MinSpanningTree return the minimum spanning tree from input graph
type MinSpanningTree struct {
	edges []graph.EdgeID

	g weightedgraph.WeightedGraph
}

// Weight return the weight value of the minimum spanning tree
func (m MinSpanningTree) Weight() (weightedgraph.Weight, error) {

	var totalWeight weightedgraph.Weight
	for _, v := range m.edges {
		w, err := m.g.Weight(v.From, v.To)
		if err != nil {
			return 0, err
		}
		totalWeight += w
	}
	return totalWeight, nil
}
