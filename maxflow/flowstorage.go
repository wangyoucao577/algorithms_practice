package maxflow

import (
	"fmt"

	"github.com/wangyoucao577/algorithms_practice/flownetwork"
	"github.com/wangyoucao577/algorithms_practice/graph"
)

type flowStorage map[graph.EdgeID]flownetwork.EdgeFlowUnit

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

func (f flowStorage) empty() bool {
	return len(f) == 0
}

func (f flowStorage) value(source graph.NodeID) flownetwork.EdgeFlowUnit {
	var maximum flownetwork.EdgeFlowUnit

	for k, v := range f {
		if k.From == source {
			maximum += v
		}
	}

	return maximum
}

func (f flowStorage) print() {
	for k, v := range f {
		fmt.Printf("edge %v flow %v, ", k, v)
	}
	fmt.Println()
}
