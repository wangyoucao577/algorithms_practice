package minspanningtree

import (
	"container/heap"

	"github.com/wangyoucao577/algorithms_practice/graph"
	"github.com/wangyoucao577/algorithms_practice/weightedgraph"
)

// Prim algorithm calculate minimum spanning tree on the input undirected graph
func Prim(g weightedgraph.WeightedGraph) (MinSpanningTree, error) {

	mst := MinSpanningTree{[]graph.EdgeID{}, g}

	pqHeap := priorityQueue{}
	g.IterateAllNodes(func(u graph.NodeID) {
		pqHeap = append(pqHeap, &nodeItem{self: u, parent: graph.InvalidNodeID, key: maxKey})
	})
	pqHeap[0].key = 0 // random starting node
	heap.Init(&pqHeap)

	//assign key for each node
	for len(pqHeap) > 0 {

		uItem := pqHeap.popMin() //pop min key node

		if uItem.parent != graph.InvalidNodeID {
			mst.edges = append(mst.edges, graph.EdgeID{From: uItem.parent, To: uItem.self})
		}

		g.IterateAdjacencyNodes(uItem.self, func(v graph.NodeID) {
			item, ok := pqHeap.find(v)
			uvWeight, _ := g.Weight(uItem.self, v)
			if ok && uvWeight < item.key {
				pqHeap.update(item, uItem.self, uvWeight)
			}
		})
	}

	return mst, nil
}

const (
	maxKey weightedgraph.Weight = weightedgraph.Weight((^uint(0)) >> 1)
)

type nodeItem struct {
	self   graph.NodeID
	parent graph.NodeID

	key weightedgraph.Weight

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}
type priorityQueue []*nodeItem

// Below 5 functions implement interfaces of Heap
func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].key < pq[j].key
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*nodeItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

//TODO: should be better way to implement the find function
func (pq priorityQueue) find(v graph.NodeID) (*nodeItem, bool) {
	for _, vItem := range pq {
		if vItem.self == v {
			return vItem, true
		}
	}
	return nil, false
}

func (pq *priorityQueue) popMin() *nodeItem {
	x := heap.Pop(pq).(*nodeItem)
	return x
}

// update modifies the key and value of an Item in the queue.
func (pq *priorityQueue) update(item *nodeItem, parent graph.NodeID, key weightedgraph.Weight) {
	item.parent = parent
	item.key = key
	heap.Fix(pq, item.index)
}
