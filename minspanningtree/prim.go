package minspanningtree

import (
	"container/heap"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Prim algorithm calculate minimum spanning tree on the input undirected graph
// Note that MinHeap in this implmenetation is based on "container/heap"
func Prim(g weightedgraph.WeightedGraph) (MinSpanningTree, error) {

	mst := MinSpanningTree{[]graph.EdgeID{}, g}

	// initialize for queryHeap
	queringHeap := &queryHeap{minHeap{}, map[graph.NodeID]*heapNode{}}
	g.IterateAllNodes(func(u graph.NodeID) {
		item := &heapNode{self: u, parent: graph.InvalidNodeID, key: maxKey}
		queringHeap.minHeap = append(queringHeap.minHeap, item)
		queringHeap.nodeToItem[u] = item
	})
	queringHeap.minHeap[0].key = 0 // random starting node
	heap.Init(&queringHeap.minHeap)

	// iterate all nodes by decreasing key
	for queringHeap.Len() > 0 {

		uItem := queringHeap.popMin() //pop min key node

		if uItem.parent != graph.InvalidNodeID {
			mst.edges = append(mst.edges, graph.EdgeID{From: uItem.parent, To: uItem.self})
		}

		g.IterateAdjacencyNodes(uItem.self, func(v graph.NodeID) {
			item, ok := queringHeap.find(v)
			uvWeight, _ := g.Weight(uItem.self, v)
			if ok && uvWeight < item.key {
				queringHeap.update(item, uItem.self, uvWeight)
			}
		})
	}

	return mst, nil
}

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
