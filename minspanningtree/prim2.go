package minspanningtree

import (
	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Prim2 algorithm calculate minimum spanning tree on the input undirected graph
// Note that MinHeap in this implmenetation is implemented by myself.
func Prim2(g weightedgraph.WeightedGraph) (MinSpanningTree, error) {
	mst := MinSpanningTree{[]graph.EdgeID{}, g}

	// initialize for queryHeap
	queringHeap := &myQueryHeap{myMinBinaryHeap{}, map[graph.NodeID]*heapNode{}}
	g.IterateAllNodes(func(u graph.NodeID) {
		item := &heapNode{self: u, parent: graph.InvalidNodeID, key: maxKey}
		queringHeap.insert(item)
		queringHeap.nodeToItem[u] = item
	})
	queringHeap.decreaseKey(1, 0) // random starting node

	// iterate all nodes by decreasing key
	for queringHeap.Len() > 0 {

		uItem := queringHeap.extractMin() //pop min key node

		if uItem.parent != graph.InvalidNodeID {
			mst.edges = append(mst.edges, graph.EdgeID{From: uItem.parent, To: uItem.self})
		}

		g.IterateAdjacencyNodes(uItem.self, func(v graph.NodeID) {
			item, ok := queringHeap.find(v)
			uvWeight, _ := g.Weight(uItem.self, v)
			if ok && uvWeight < item.key {
				queringHeap.decreaseKey(item.index+1, uvWeight)
			}
		})
	}

	return mst, nil
}

type myMinBinaryHeap []*heapNode // priority queue with minimum key at the top, my implementation

func (mh myMinBinaryHeap) Len() int { return len(mh) }
func (mh myMinBinaryHeap) Less(i, j int) bool {
	return mh[i].key < mh[j].key
}
func (mh myMinBinaryHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].index = i
	mh[j].index = j
}

// `n` start by 1 here, not 0
func (mh myMinBinaryHeap) leftChild(n int) int {
	return n * 2
}
func (mh myMinBinaryHeap) rightChild(n int) int {
	return n*2 + 1
}
func (mh myMinBinaryHeap) parent(n int) int {
	return n / 2
}

func (mh *myMinBinaryHeap) minHeapity(n int) { // i.e. down

	for {
		left := mh.leftChild(n)
		right := mh.rightChild(n)

		smallest := n

		if left <= mh.Len() && mh.Less(left-1, smallest-1) {
			smallest = left
		}
		if right <= mh.Len() && mh.Less(right-1, smallest-1) {
			smallest = right
		}

		if smallest == n {
			break
		}

		mh.Swap(smallest-1, n-1)
		n = smallest
	}
}

func (mh *myMinBinaryHeap) decreaseKey(n int, key weightedgraph.Weight) { // i.e. up
	if n > mh.Len() || n <= 0 {
		return
	}

	if (*mh)[n-1].key <= key {
		return // don't allow increase key for min heap
	}
	(*mh)[n-1].key = key

	for {
		if n == 1 { // already the root
			break
		}

		p := mh.parent(n)

		if !mh.Less(n-1, p-1) {
			break
		}

		mh.Swap(n-1, p-1)
		n = p
	}
}

func (mh *myMinBinaryHeap) insert(node *heapNode) {
	node.index = mh.Len()
	*mh = append(*mh, node)

	mh.decreaseKey(node.index+1, node.key)
}

func (mh *myMinBinaryHeap) extractMin() *heapNode {
	if mh.Len() <= 0 {
		return nil
	}

	mh.Swap(0, mh.Len()-1)
	min := (*mh)[mh.Len()-1]
	*mh = (*mh)[:mh.Len()-1]

	mh.minHeapity(1)

	return min
}

type myQueryHeap struct {
	myMinBinaryHeap
	nodeToItem map[graph.NodeID]*heapNode
}

func (qh *myQueryHeap) find(node graph.NodeID) (*heapNode, bool) {
	item, ok := qh.nodeToItem[node]
	if !ok {
		return nil, false
	}
	if item.index == -1 { //have been popped
		return nil, false
	}
	return item, true
}
