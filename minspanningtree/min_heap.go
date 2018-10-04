package minspanningtree

import (
	"container/heap"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Construct MinHeap based on "container/heap"

type minHeap []*heapNode // priority queue with minimum key at the top

// Below 5 functions implement interfaces of Heap
func (mh minHeap) Len() int { return len(mh) }
func (mh minHeap) Less(i, j int) bool {
	return mh[i].key < mh[j].key
}
func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].index = i
	mh[j].index = j
}
func (mh *minHeap) Push(x interface{}) {
	n := len(*mh)
	item := x.(*heapNode)
	item.index = n
	*mh = append(*mh, item)
}
func (mh *minHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety in find
	*mh = old[0 : n-1]
	return item
}

// queryHeap designed for both min-heap and find-by-node
type queryHeap struct {
	minHeap
	nodeToItem map[graph.NodeID]*heapNode
}

func (qh *queryHeap) popMin() *heapNode {
	x := heap.Pop(&qh.minHeap).(*heapNode)
	return x
}
func (qh *queryHeap) find(node graph.NodeID) (*heapNode, bool) {
	item, ok := qh.nodeToItem[node]
	if !ok {
		return nil, false
	}
	if item.index == -1 { //have been popped
		return nil, false
	}
	return item, true
}
func (qh *queryHeap) update(item *heapNode, parent graph.NodeID, key weightedgraph.Weight) {
	//NOTE: this update will only affect nodeItem of nodeToItem,
	// since the nodeToItem and priorityQueue shared the same pointer

	item.parent = parent
	item.key = key
	heap.Fix(&qh.minHeap, item.index)
}
