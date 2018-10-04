package minspanningtree

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

const (
	maxKey weightedgraph.Weight = weightedgraph.Weight((^uint(0)) >> 1)
)

type heapNode struct {
	self   graph.NodeID
	parent graph.NodeID

	key weightedgraph.Weight

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}
