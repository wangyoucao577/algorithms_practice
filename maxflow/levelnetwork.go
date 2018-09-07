package maxflow

import (
	"github.com/wangyoucao577/algorithms_practice/flownetwork"
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/levelgraph"
)

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
